package controller

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

// Params struct holds the incoming request parameters
type Params struct {
	TimeLimit   int    `form:"time_limit"`
	Params      string `form:"params"`
	ExcerciseId int    `form:"execercise_id"`
}

type TestCasePayLoad struct {
	TimeLimit   int `form:"time_limit"`
	ExcerciseId int `form:"execercise_id"`
}

type AddTestData struct {
	inputs string
	result string
}

// PythonProcessPool represents a pool of Python processes.
type PythonProcessPool struct {
	mu     sync.Mutex
	pool   []*exec.Cmd
	maxCap int
}

type Result struct {
	Match    bool
	Output   string
	Expected string
}

// NewPythonProcessPool creates a new Python process pool with a specified maximum capacity.
func NewPythonProcessPool(maxCap int) *PythonProcessPool {
	return &PythonProcessPool{
		pool:   make([]*exec.Cmd, 0, maxCap),
		maxCap: maxCap,
	}
}

// Get retrieves a Python process from the pool. If the pool is empty, a new process is created.
func (p *PythonProcessPool) Get() (*exec.Cmd, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if len(p.pool) > 0 {
		process := p.pool[len(p.pool)-1]
		p.pool = p.pool[:len(p.pool)-1]
		return process, nil
	}

	return exec.Command("python"), nil
}

// Put returns a Python process to the pool.
func (p *PythonProcessPool) Put(process *exec.Cmd) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if len(p.pool) < p.maxCap {
		p.pool = append(p.pool, process)
	} else {
		process.Process.Kill()
	}
}

func standardizeLineBreaks(input string) string {
	return strings.ReplaceAll(input, "\r\n", "\n")
}

// runScript executes the Python script and returns the output or an error
func runScript(ctx context.Context, pythonProcessPool *PythonProcessPool, userCode string, params string) ([]byte, error) {
	cmd, err := pythonProcessPool.Get()
	if err != nil {
		return nil, err
	}
	defer pythonProcessPool.Put(cmd)

	// Pass the file path of the Python wrapper script, the user's Python code, and the inputs to the Python process
	cmd = exec.Command("python", "./runner/runner.py", userCode, params)

	return cmd.CombinedOutput()
}

func getUserCode(c *fiber.Ctx) (string, error) {
	// Retrieve the user's Python code from the request
	userCode := c.FormValue("python_code")

	// If python_code form field is empty, then check for a .py file in the form data
	if userCode == "" {
		file, err := c.FormFile("python_file")
		if err != nil {
			return "", err
		}

		// If there's a .py file, read its contents into userCode
		pyFile, err := file.Open()
		if err != nil {
			return "", err
		}
		defer pyFile.Close()

		pyCodeBytes, err := io.ReadAll(pyFile)
		if err != nil {
			return "", err
		}
		userCode = string(pyCodeBytes)
	}
	return userCode, nil
}

type ResponseBody struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func respondWithJSON(c *fiber.Ctx, status int, payload ResponseBody) error {
	return c.Status(status).JSON(payload)
}

func handleError(c *fiber.Ctx, err error, status int, message string) error {
	fmt.Println("err: ", err)
	return respondWithJSON(c, status, ResponseBody{
		Status:  "error",
		Message: message,
		Data:    err.Error(),
	})
}

// ExecuteCode is the Fiber handler for executing Python code
func ExecuteCode(c *fiber.Ctx) error {
	params := new(Params)
	if err := c.BodyParser(params); err != nil {
		return handleError(c, err, fiber.StatusBadRequest, "Failed to parse request body")
	}

	userCode, err := getUserCode(c)
	if err != nil {
		return handleError(c, err, fiber.StatusBadRequest, "No Python code or .py file provided")
	}

	pythonProcessPool := NewPythonProcessPool(20) // Adjust the pool size as per your requirements
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(params.TimeLimit)*time.Second)
	defer cancel()

	output, err := runScript(ctx, pythonProcessPool, userCode, params.Params)
	if ctx.Err() == context.DeadlineExceeded {
		return handleError(c, err, fiber.StatusRequestTimeout, "Python script execution timed out")
	} else if err != nil {
		return handleError(c, err, fiber.StatusInternalServerError, "Failed to execute Python script")
	}

	return c.SendString(string(output))

	// return respondWithJSON(c, fiber.StatusOK, ResponseBody{
	// 	Status:  "success",
	// 	Message: "Code executed successfully",
	// 	Data:    string(output),
	// })
}

func getTestCase(exerciseID int) ([]AddTestData, error) {
	var (
		testData []AddTestData
	)
	var (
		testcaseContent string
		testcaseOutput  string
	)

	host := os.Getenv("DB_HOST")
	portStr := os.Getenv("DB_PORT")
	userdb := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	port, err := strconv.Atoi(portStr) // Convert port string to int
	if err != nil {
		log.Fatal("Error converting port to int:", err)
	}

	var connectString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", userdb, password, host, port, dbName)
	db, err := sql.Open("mysql", connectString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT testcase_content, testcase_output FROM exercise_testcase WHERE exercise_id = %d", exerciseID)

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&testcaseContent, &testcaseOutput); err != nil {
			log.Fatal(err)
		}

		// Append data to the testData slice
		testData = append(testData, AddTestData{
			inputs: testcaseContent,
			result: testcaseOutput,
		})
	}

	return testData, nil
}

func RunCodeWithTestCase(c *fiber.Ctx) error {
	//get testdata from db
	exerciseID := 1301449
	testData, err := getTestCase(exerciseID)
	if err != nil {
		return handleError(c, err, fiber.StatusInternalServerError, "Failed to retrieve test data")
	}

	params := new(TestCasePayLoad)
	if err := c.BodyParser(params); err != nil {
		return handleError(c, err, fiber.StatusBadRequest, "Failed to parse request body")
	}

	userCode, err := getUserCode(c)
	if err != nil {
		return handleError(c, err, fiber.StatusBadRequest, "No Python code or .py file provided")
	}

	pythonProcessPool := NewPythonProcessPool(20) // Adjust the pool size as per your requirements
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(params.TimeLimit)*time.Second)
	defer cancel()

	// testData := []AddTestData{
	// 	{"1 2 3 4 5 ",
	// 		" *** Sum even / Subtract odd ***\r\nEnter numbers : 1 2 3 4 5 \r\nSum is -3\n"},
	// 	{"4 6 3 5 8 ",
	// 		" *** Sum even / Subtract odd ***\r\nEnter numbers : 4 6 3 5 8 \r\nSum is 10\r\n"},
	// 	{"0 0 1 1 2 ",
	// 		" *** Sum even / Subtract odd ***\r\nEnter numbers : 0 0 1 1 2 \r\nSum is 0\r\n"},
	// }

	var resultList []Result // Declare resultList as a local variable

	for _, datum := range testData {
		output, err := runScript(ctx, pythonProcessPool, userCode, datum.inputs) // Declare output as a local variable
		result := Result{
			Match:    standardizeLineBreaks(string(output)) == standardizeLineBreaks(datum.result),
			Output:   standardizeLineBreaks(string(output)),
			Expected: standardizeLineBreaks(datum.result),
		}
		resultList = append(resultList, result)
		if ctx.Err() == context.DeadlineExceeded {
			return handleError(c, err, fiber.StatusRequestTimeout, "Python script execution timed out")
		} else if err != nil {
			return handleError(c, err, fiber.StatusInternalServerError, "Failed to execute Python script")
		}
	}
	return respondWithJSON(c, fiber.StatusOK, ResponseBody{
		Status:  "success",
		Message: "Code executed successfully",
		Data:    resultList,
	})
}

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

var pythonProcessPool = NewPythonProcessPool(100)

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

type Result struct {
	Match    bool
	Output   string
	Expected string
}

type PythonProcessPool struct {
	mu     sync.Mutex
	active int
	maxCap int
}

func NewPythonProcessPool(maxCap int) *PythonProcessPool {
	return &PythonProcessPool{
		maxCap: maxCap,
	}
}

func (p *PythonProcessPool) CanRun() bool {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.active < p.maxCap {
		p.active++
		return true
	}
	return false
}

func (p *PythonProcessPool) Done() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.active--
}

func standardizeLineBreaks(input string) string {
	return strings.ReplaceAll(input, "\r\n", "\n")
}

func runScript(ctx context.Context, pool *PythonProcessPool, userCode, params string) ([]byte, error) {
	if !pool.CanRun() {
		return nil, fmt.Errorf("maximum concurrent executions reached")
	}
	defer pool.Done()

	cmd := exec.CommandContext(ctx, "python", "./runner/runner.py", userCode, params)
	return cmd.CombinedOutput()
}

func getUserCode(c *fiber.Ctx) (string, error) {
	userCode := c.FormValue("python_code")

	if userCode == "" {
		file, err := c.FormFile("python_file")
		if err != nil {
			return "", err
		}

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

func respondWithError(c *fiber.Ctx, err error, status int, message string) error {
	fmt.Println("err: ", err)
	return c.Status(status).JSON(fiber.Map{
		"status":  "error",
		"message": message,
		"data":    err.Error(),
	})
}

func executeWithTimeout(c *fiber.Ctx, timeLimit int, userCode, inputParams string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeLimit)*time.Second)
	defer cancel()

	output, err := runScript(ctx, pythonProcessPool, userCode, inputParams)

	if ctx.Err() == context.DeadlineExceeded {
		return nil, fmt.Errorf("Timeout")
	}

	return output, err
}

func ExecuteCode(c *fiber.Ctx) error {
	params := new(Params)
	if err := c.BodyParser(params); err != nil {
		return respondWithError(c, err, fiber.StatusBadRequest, "Failed to parse request body")
	}

	userCode, err := getUserCode(c)
	if err != nil {
		return respondWithError(c, err, fiber.StatusBadRequest, "No Python code or .py file provided")
	}

	output, err := executeWithTimeout(c, params.TimeLimit, userCode, params.Params)
	if err != nil {
		if err.Error() == "Timeout" {
			return respondWithError(c, err, fiber.StatusRequestTimeout, "Python script execution timed out")
		}
		return respondWithError(c, err, fiber.StatusInternalServerError, "Failed to execute Python script")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Code executed successfully",
		"data":    standardizeLineBreaks(string(output)),
	})
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
		return respondWithError(c, err, fiber.StatusInternalServerError, "Failed to retrieve test data")
	}

	params := new(TestCasePayLoad)
	if err := c.BodyParser(params); err != nil {
		return respondWithError(c, err, fiber.StatusBadRequest, "Failed to parse request body")
	}

	userCode, err := getUserCode(c)
	if err != nil {
		return respondWithError(c, err, fiber.StatusBadRequest, "No Python code or .py file provided")
	}

	var resultList []Result

	for _, datum := range testData {
		output, err := executeWithTimeout(c, params.TimeLimit, userCode, datum.inputs)
		if err != nil {
			if err.Error() == "Timeout" {
				return respondWithError(c, err, fiber.StatusRequestTimeout, "Python script execution timed out")
			}
			return respondWithError(c, err, fiber.StatusInternalServerError, "Failed to execute Python script")
		}

		resultList = append(resultList, Result{
			Match:    standardizeLineBreaks(string(output)) == standardizeLineBreaks(datum.result),
			Output:   standardizeLineBreaks(string(output)),
			Expected: standardizeLineBreaks(datum.result),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Code executed successfully",
		"data":    resultList,
	})
}

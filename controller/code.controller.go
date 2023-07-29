package controller

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Params struct holds the incoming request parameters
type Params struct {
	TimeLimit int    `form:"time_limit"`
	Params    string `form:"params"`
}

// PythonProcessPool represents a pool of Python processes.
type PythonProcessPool struct {
	mu     sync.Mutex
	pool   []*exec.Cmd
	maxCap int
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

// runScript executes the Python script and returns the output or an error
func runScript(ctx context.Context, userCode string, params string) ([]byte, error) {
	pythonProcessPool := NewPythonProcessPool(20) // Adjust the pool size as per your requirements
	cmd, err := pythonProcessPool.Get()
	if err != nil {
		fmt.Println("err: ", err)

		return nil, err
	}
	defer pythonProcessPool.Put(cmd)

	// Pass the file path of the Python wrapper script, the user's Python code, and the inputs to the Python process
	cmd = exec.Command("python", "./runner/runner.py", userCode, params)

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("err: ", err)

		return nil, err
	}

	return out, nil
}

// ExecuteCode is the Fiber handler for executing Python code
func ExecuteCode(c *fiber.Ctx) error {
	params := new(Params)
	if err := c.BodyParser(params); err != nil {
		fmt.Println("err: ", err)

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	// Retrieve the user's Python code from the request
	userCode := c.FormValue("python_code")

	// If python_code form field is empty, then check for a .py file in the form data
	if userCode == "" {
		file, err := c.FormFile("python_file")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "No Python code or .py file provided",
			})
		}

		// If there's a .py file, read its contents into userCode
		pyFile, err := file.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("Failed to open Python file: %s", err.Error()),
			})
		}
		defer pyFile.Close()

		pyCodeBytes, err := io.ReadAll(pyFile)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("Failed to read Python file: %s", err.Error()),
			})
		}
		userCode = string(pyCodeBytes)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(params.TimeLimit)*time.Second)
	defer cancel()

	output, err := runScript(ctx, userCode, params.Params)
	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("err: ", err)

		return c.Status(fiber.StatusRequestTimeout).JSON(fiber.Map{
			"error": "Python script execution timed out",
		})
	} else if err != nil {
		fmt.Println("err: ", err)

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to execute Python script: %s", err.Error()),
		})
	}

	return c.SendString(string(output))
}

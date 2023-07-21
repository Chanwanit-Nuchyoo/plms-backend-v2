package controller

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

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

	// If the pool is not empty, retrieve a process from the pool.
	if len(p.pool) > 0 {
		process := p.pool[len(p.pool)-1]
		p.pool = p.pool[:len(p.pool)-1]
		return process, nil
	}

	// If the pool is empty, create a new process.
	return exec.Command("python"), nil
}

// Put returns a Python process to the pool.
func (p *PythonProcessPool) Put(process *exec.Cmd) {
	p.mu.Lock()
	defer p.mu.Unlock()

	// If the pool is not full, return the process to the pool.
	if len(p.pool) < p.maxCap {
		p.pool = append(p.pool, process)
	} else {
		// If the pool is full, kill the process.
		process.Process.Kill()
	}
}

func ExecuteCode(c *fiber.Ctx) error {
	// Parse the incoming .py file from the request body
	file, err := c.FormFile("python_code")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse .py file from request",
		})
	}

	// Parse the time limit from the request
	timeLimitStr := c.FormValue("time_limit")

	// Validate and convert the time limit
	timeLimit, err := strconv.Atoi(timeLimitStr)
	if err != nil || timeLimit <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid or missing time_limit value",
		})
	}

	// Save the .py file temporarily on the server
	tempFile, err := ioutil.TempFile("", "*.py")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create temporary file",
		})
	}
	defer func() {
		tempFile.Close() // Close the file handle when done
		os.Remove(tempFile.Name())
	}()

	// Read the file contents and save them to the temporary file
	fileBytes, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to read .py file",
		})
	}
	defer fileBytes.Close()

	_, err = io.Copy(tempFile, fileBytes)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to write temporary file",
		})
	}

	// Execute the Python script using the `os/exec` package with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeLimit)*time.Second)
	defer cancel()

	pythonProcessPool := NewPythonProcessPool(20) // Adjust the pool size as per your requirements

	cmd, err := pythonProcessPool.Get()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get Python process from the pool",
		})
	}
	defer pythonProcessPool.Put(cmd)

	// cmd.SysProcAttr = &syscall.SysProcAttr{ // Set the memory limit for the command (if supported by the OS)
	// 	Rlimit: &syscall.Rlimit{
	// 		Cur: int64(100 * 1024 * 1024), // 100 MB memory limit (adjust as needed)
	// 		Max: int64(100 * 1024 * 1024), // 100 MB memory limit (adjust as needed)
	// 	},
	// }

	cmd = exec.CommandContext(ctx, "python", tempFile.Name()) // Use "python" to execute the script

	// Capture the output (stdout and stderr) of the Python script execution
	output, err := cmd.CombinedOutput()
	if ctx.Err() == context.DeadlineExceeded {
		return c.Status(fiber.StatusRequestTimeout).JSON(fiber.Map{
			"error": "Python script execution timed out",
		})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to execute Python script: %s", err.Error()),
		})
	}

	// Send the output back as the response
	return c.SendString(string(output))
}

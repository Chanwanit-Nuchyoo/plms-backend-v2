package controller

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ExecuteCode(c *fiber.Ctx) error {
	// Parse the incoming .py file from the request body
	file, err := c.FormFile("python_code")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse .py file from request",
		})
	}

	// Parse the time limit and memory limit from the request
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

	cmd := exec.CommandContext(ctx, "python", tempFile.Name()) // Use "python" to execute the script

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

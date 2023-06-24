package customErr

import "github.com/gofiber/fiber/v2"

type ErrorCode string

const (
	ErrBodyParsingFailed ErrorCode = "01"
	ErrUserAlreadyExist  ErrorCode = "02"
	ErrUserCreateFailed  ErrorCode = "03"
	// Add more error codes as needed
)

type ErrorMessage struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}

var ErrorMessages = map[ErrorCode]ErrorMessage{
	ErrBodyParsingFailed: {ErrorCode: fiber.StatusBadRequest, Message: "Invalid request data."},
	ErrUserAlreadyExist:  {ErrorCode: fiber.StatusBadRequest, Message: "User already exists."},
	ErrUserCreateFailed:  {ErrorCode: fiber.StatusInternalServerError, Message: "Failed to perform the requested operation."},
	// Add more error messages as needed
}

// GetErrorMessage retrieves the error message for a given error code
func GetErrorMessage(code ErrorCode) ErrorMessage {
	msg, ok := ErrorMessages[code]

	if !ok {
		return ErrorMessage{
			ErrorCode: fiber.StatusInternalServerError,
			Message:   "An unknown error occurred. Please try again later or contact support for assistance.",
		}
	}

	return msg
}

func HandleError(c *fiber.Ctx, code ErrorCode) error {
	msg := GetErrorMessage(code)
	return c.Status(msg.ErrorCode).JSON(msg)
}

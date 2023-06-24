package controller

import (
	"plms-backend/customErr"
	"plms-backend/db"
	"plms-backend/ent"
	"plms-backend/ent/user"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	req := new(ent.User)

	if err := c.BodyParser(req); err != nil {
		return customErr.HandleError(c, customErr.ErrBodyParsingFailed)
	}

	dbClient, _ := db.GetDBClient()

	queriedUser, _ := dbClient.User.Query().Where(user.Or(
		user.Username(req.Username),
		user.Email(req.Email),
	)).Only(c.Context())

	if queriedUser != nil {

		return customErr.HandleError(c, customErr.ErrUserAlreadyExist)
	}

	user, err := dbClient.User.Create().SetUsername(req.Username).SetEmail(req.Email).SetName(req.Name).SetPassword(req.Password).Save(c.Context())

	if err != nil {
		return customErr.HandleError(c, customErr.ErrUserCreateFailed)
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

package handlers

import (
	"github.com/go-interview/database"
	"github.com/go-interview/helper"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
)

type InputRequest struct {
	AccountID string `json:"account_id"`
	Reference string `json:"reference"`
	Amount    uint   `json:"amount"`
}

func (s InputRequest) Validate() error {
	valid := validation.ValidateStruct(&s,
		validation.Field(&s.AccountID, validation.Required),
		validation.Field(&s.Reference, validation.Required),
		validation.Field(&s.Amount, validation.Required),
	)

	return valid
}
func ProcessTransaction(c *fiber.Ctx) error {
	db := database.DB
	var (
		input InputRequest
	)

	if err := c.BodyParser(&input); err != nil {
		return helper.Response(c, err, err.Error(), false, 400)
	}

	if err := input.Validate(); err != nil {
		return helper.Response(c, err, err.Error(), false, 400)
	}
	return c.JSON("")
}

func GetTransaction(c *fiber.Ctx) error {
	return c.JSON("")
}

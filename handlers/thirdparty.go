package handlers

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
)

type InputRequest struct {
	AccountID uint   `json:"account_id"`
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
	var (
		input InputRequest
	)

	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	if err := input.Validate(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}
	output := InputRequest{
		AccountID: input.AccountID,
		Reference: input.Reference,
		Amount:    input.Amount,
	}
	return c.Status(http.StatusOK).JSON(output)
}

func GetTransaction(c *fiber.Ctx) error {

	reference := c.Params("ref")
	output := InputRequest{
		AccountID: 1111111111,
		Reference: reference,
		Amount:    2000000,
	}

	return c.Status(http.StatusAccepted).JSON(output)
}

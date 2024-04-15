package handlers

import (
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
func ProcessTransaction(c *fiber.Ctx) (InputRequest, error) {
	var (
		input InputRequest
	)

	if err := c.BodyParser(&input); err != nil {
		return InputRequest{}, err
	}

	if err := input.Validate(); err != nil {
		return InputRequest{}, err
	}
	output := InputRequest{
		AccountID: input.AccountID,
		Reference: input.Reference,
		Amount:    input.Amount,
	}
	return output, nil
}

func GetTransaction(c *fiber.Ctx) error {
	return c.JSON("")
}

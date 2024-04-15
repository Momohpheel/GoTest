package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/go-interview/database"
	"github.com/go-interview/helper"
	"github.com/go-interview/models"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
)

var (
	baseUrl = "http://127.0.0.1:4500"
)

type requestBody struct {
	AccountID uint   `json:"account_id" validate:"required"`
	Reference string `json:"reference" validate:"required"`
	Amount    uint   `json:"amount" validate:"required"`
}

func (s requestBody) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.AccountID, validation.Required),
		validation.Field(&s.Reference, validation.Required),
		validation.Field(&s.Amount, validation.Required),
	)
}
func CreateTransaction(c *fiber.Ctx) error {
	var input requestBody
	err := c.BodyParser(&input)
	if err != nil {
		return helper.Response(c, err, err.Error(), false, 400)
	}

	if err = input.Validate(); err != nil {
		return helper.Response(c, err, err.Error(), false, 400)
	}

	db := database.DB
	var user models.User

	rows := db.Where("id = ?", input.AccountID).First(&user)
	if rows.RowsAffected == 0 {
		return helper.Response(c, rows.Error, "User not found", false, 401)
	}

	request := models.Transaction{
		AccountID: input.AccountID,
		Reference: input.Reference,
		Amount:    input.Amount,
	}

	rows = db.Create(&request)
	if rows.RowsAffected == 0 {
		return helper.Response(c, rows.Error, "transaction failed", false, 401)
	}

	//third party integration
	response, err := ThirdParty(c, request.ID)
	if err != nil {
		return helper.Response(c, err.Error(), "Couldn't create transaction in third-party provider system", false, 401)
	}

	return helper.Response(c, response, "transaction successfull", true, 201)

}

type AccountCreation struct {
	AccountID uint   `json:"account_id"`
	Reference string `json:"reference"`
	Amount    uint   `json:"amount"`
}

type ThirdPartyResponse struct {
	AccountID uint   `json:"account_id"`
	Reference string `json:"reference"`
	Amount    uint   `json:"amount"`
}

func ThirdParty(c *fiber.Ctx, id uint) (ThirdPartyResponse, error) {

	var (
		data models.Transaction
		user models.User
	)
	db := database.DB

	url := baseUrl + "/v1/payments"

	row := db.Where("id = ?", id).First(&data)

	if row.RowsAffected == 0 {
		return ThirdPartyResponse{}, errors.New("data does not exist in the database")
	}

	db.Where("id = ?", data.AccountID).First(&user)

	request := AccountCreation{
		AccountID: data.AccountID,
		Reference: data.Reference,
		Amount:    data.Amount,
	}

	jsonReq, err := json.Marshal(request)
	if err != nil {
		return ThirdPartyResponse{}, errors.New("Error making a request -" + err.Error())
	}
	client := &http.Client{}
	responseBody := bytes.NewBuffer(jsonReq)

	req, err := http.NewRequest(http.MethodPost, url, responseBody)

	if err != nil {
		return ThirdPartyResponse{}, errors.New("Error making a request 1- " + err.Error())
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return ThirdPartyResponse{}, errors.New("Error making a request 2- " + err.Error())
	}
	defer resp.Body.Close()

	bodyByte, err := io.ReadAll(resp.Body)

	if err != nil {
		return ThirdPartyResponse{}, errors.New("Error making a request 3- " + err.Error())
	}

	//return ThirdPartyResponse{}, errors.New(string(bodyByte))
	var res ThirdPartyResponse
	errs := json.Unmarshal(bodyByte, &res)

	if errs != nil {
		return ThirdPartyResponse{}, errors.New("Error making a request 4- " + errs.Error())
	}

	if len(res.Reference) > 1 {
		amount := user.AccountBalance
		totalamount := amount + data.Amount

		//update amount on user table
		db.Model(&models.User{}).Where("id = ?", user.ID).Updates(models.User{AccountBalance: totalamount})

		return res, nil
	} else {

		return ThirdPartyResponse{}, errors.New("error connecting to thirdparty")
	}

}

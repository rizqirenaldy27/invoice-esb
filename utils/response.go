package utils

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ResponseOKWithDataModel struct {
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type ResponseBasicModel struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func ResponseDuplicate(c *fiber.Ctx, data interface{}) error {
	response := ResponseOKWithDataModel{
		Status:  false,
		Data:    data,
		Message: "Data Duplicate",
	}

	return c.Status(http.StatusBadRequest).JSON(response)
}

func ResponseError(c *fiber.Ctx, message string) error {
	response := ResponseBasicModel{
		Status:  false,
		Message: message,
	}

	return c.Status(http.StatusBadRequest).JSON(response)
}

func ResponseErrWithCode(c *fiber.Ctx, message string, code int) error {
	response := ResponseBasicModel{
		Status:  false,
		Message: message,
	}

	return c.Status(code).JSON(response)
}

func ResponseCreated(c *fiber.Ctx, data interface{}) error {
	response := ResponseOKWithDataModel{
		Status:  true,
		Data:    data,
		Message: "Created",
	}

	return c.Status(http.StatusCreated).JSON(response)
}

func ResponseUpdated(c *fiber.Ctx, data interface{}) error {
	response := ResponseOKWithDataModel{
		Status:  true,
		Data:    data,
		Message: "Updated",
	}

	return c.JSON(response)
}

func ResponseOK(c *fiber.Ctx, message string) error {
	response := ResponseBasicModel{
		Status:  true,
		Message: message,
	}

	return c.JSON(response)
}

func ResponseDetailOK(c *fiber.Ctx, data interface{}) error {
	response := ResponseOKWithDataModel{
		Status:  true,
		Data:    data,
		Message: "OK",
	}

	return c.Status(http.StatusCreated).JSON(response)
}

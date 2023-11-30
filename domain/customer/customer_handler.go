package customer

import (
	"errors"
	"strconv"
	"strings"

	"github.com/rizqirenaldy27/invoice-esb/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CustomerHandler struct {
	db *gorm.DB
}

func NewCustomerHandler(db *gorm.DB) *CustomerHandler {
	var customerHandler = CustomerHandler{}
	customerHandler.db = db
	return &customerHandler
}

func (ch *CustomerHandler) Create(c *fiber.Ctx) error {
	var customer Customer

	if err := c.BodyParser(&customer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	result := ch.db.Create(&customer)
	if result.Error != nil {
		msgError := result.Error.Error()
		if strings.Contains(msgError, "Duplicate entry") {
			utils.ResponseError(c, "Customer Data Duplicate "+customer.CustomerName)
			return nil
		} else {
			utils.ResponseError(c, msgError)
			return nil
		}
	}

	utils.ResponseCreated(c, customer)
	return nil
}

func (ch *CustomerHandler) ReadByID(c *fiber.Ctx) error {
	customerID, err := strconv.Atoi(c.Params("customer_id"))
	if err != nil {
		utils.ResponseError(c, err.Error())
		return nil
	}

	var customer Customer
	err = ch.db.First(&customer, customerID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ResponseError(c, "Customer not found")
		return nil
	}

	utils.ResponseDetailOK(c, customer)
	return nil
}

func (ch *CustomerHandler) Read(c *fiber.Ctx) error {
	var customer []Customer
	err := ch.db.Find(&customer).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ResponseError(c, "Customer not found")
		return nil
	}

	utils.ResponseDetailOK(c, customer)
	return nil
}

func (ch *CustomerHandler) Update(c *fiber.Ctx) error {
	customerID, err := strconv.Atoi(c.Params("customer_id"))
	if err != nil {
		utils.ResponseError(c, err.Error())
		return nil
	}

	var existingCustomer Customer
	err = ch.db.First(&existingCustomer, customerID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ResponseError(c, "Customer not found")
		return nil
	}

	var updatedCustomerData Customer
	if err := c.BodyParser(&updatedCustomerData); err != nil {
		utils.ResponseError(c, err.Error())
		return nil
	}

	errUpdate := ch.db.Model(&existingCustomer).Updates(updatedCustomerData).Error
	if errUpdate != nil {
		msgError := errUpdate.Error()
		if strings.Contains(msgError, "Duplicate entry") {
			utils.ResponseError(c, "Customer Data Duplicate")
			return nil
		} else {
			utils.ResponseError(c, msgError)
			return nil
		}
	}

	utils.ResponseUpdated(c, existingCustomer)
	return nil
}

func (ch *CustomerHandler) Delete(c *fiber.Ctx) error {
	customerID, err := strconv.Atoi(c.Params("customer_id"))
	if err != nil {
		utils.ResponseError(c, err.Error())
		return nil
	}

	var existingCustomer Customer
	err = ch.db.First(&existingCustomer, customerID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ResponseError(c, "Customer not found")
		return nil
	}

	errDelete := ch.db.Delete(&existingCustomer).Error
	if errDelete != nil {
		utils.ResponseError(c, errDelete.Error())
		return nil
	}

	utils.ResponseOK(c, "Customer deleted successfully")
	return nil
}

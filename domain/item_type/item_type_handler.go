package item_type

import (
	"errors"
	"strconv"
	"strings"

	"github.com/rizqirenaldy27/invoice-esb/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ItemTypeHandler struct {
	db *gorm.DB
}

func NewItemTypeHandler(db *gorm.DB) *ItemTypeHandler {
	var itemTypeHandler = ItemTypeHandler{}
	itemTypeHandler.db = db
	return &itemTypeHandler
}

func (ch *ItemTypeHandler) Create(c *fiber.Ctx) error {
	var itemType ItemType

	if err := c.BodyParser(&itemType); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	result := ch.db.Create(&itemType)
	if result.Error != nil {
		msgError := result.Error.Error()
		if strings.Contains(msgError, "Duplicate entry") {
			utils.ResponseError(c, "Item Type Data Duplicate "+itemType.ItemTypeName)
			return nil
		} else {
			utils.ResponseError(c, msgError)
			return nil
		}
	}

	utils.ResponseCreated(c, itemType)
	return nil
}

func (ch *ItemTypeHandler) ReadByID(c *fiber.Ctx) error {
	ItemTypeID, err := strconv.Atoi(c.Params("item_type_id"))
	if err != nil {
		utils.ResponseError(c, err.Error())
		return nil
	}

	var itemType ItemType
	err = ch.db.First(&itemType, ItemTypeID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ResponseError(c, "Item Type not found")
		return nil
	}

	utils.ResponseDetailOK(c, itemType)
	return nil
}

func (ch *ItemTypeHandler) Read(c *fiber.Ctx) error {
	var itemType []ItemType
	err := ch.db.Find(&itemType).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ResponseError(c, "Item Type not found")
		return nil
	}

	utils.ResponseDetailOK(c, itemType)
	return nil
}

func (ch *ItemTypeHandler) Update(c *fiber.Ctx) error {
	ItemTypeID, err := strconv.Atoi(c.Params("item_type_id"))
	if err != nil {
		utils.ResponseError(c, err.Error())
		return nil
	}

	var existingCustomer ItemType
	err = ch.db.First(&existingCustomer, ItemTypeID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ResponseError(c, "Item Type not found")
		return nil
	}

	var updatedCustomerData ItemType
	if err := c.BodyParser(&updatedCustomerData); err != nil {
		utils.ResponseError(c, err.Error())
		return nil
	}

	errUpdate := ch.db.Model(&existingCustomer).Updates(updatedCustomerData).Error
	if errUpdate != nil {
		msgError := errUpdate.Error()
		if strings.Contains(msgError, "Duplicate entry") {
			utils.ResponseError(c, "Item Type Data Duplicate")
			return nil
		} else {
			utils.ResponseError(c, msgError)
			return nil
		}
	}

	utils.ResponseUpdated(c, existingCustomer)
	return nil
}

func (ch *ItemTypeHandler) Delete(c *fiber.Ctx) error {
	ItemTypeID, err := strconv.Atoi(c.Params("item_type_id"))
	if err != nil {
		utils.ResponseError(c, err.Error())
		return nil
	}

	var existingCustomer ItemType
	err = ch.db.First(&existingCustomer, ItemTypeID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ResponseError(c, "Item Type not found")
		return nil
	}

	errDelete := ch.db.Delete(&existingCustomer).Error
	if errDelete != nil {
		utils.ResponseError(c, errDelete.Error())
		return nil
	}

	utils.ResponseOK(c, "Item Type deleted successfully")
	return nil
}

package item

import (
	"errors"
	"strconv"
	"strings"

	"github.com/rizqirenaldy27/invoice-esb/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ItemHandler struct {
	db *gorm.DB
}

func NewItemHandler(db *gorm.DB) *ItemHandler {
	var itemHandler = ItemHandler{}
	itemHandler.db = db
	return &itemHandler
}

func (ch *ItemHandler) Create(c *fiber.Ctx) error {
	var item Item

	if err := c.BodyParser(&item); err != nil {
		utils.ResponseError(c, err.Error())
		return nil
	}

	result := ch.db.Create(&item)
	if result.Error != nil {
		msgError := result.Error.Error()
		if strings.Contains(msgError, "Duplicate entry") {
			utils.ResponseError(c, "Item Data Duplicate "+item.ItemName)
			return nil
		} else if strings.Contains(msgError, "foreign key constraint fails") {
			utils.ResponseError(c, "Item Type ID not found")
			return nil
		} else {
			utils.ResponseError(c, msgError)
			return nil
		}
	}

	utils.ResponseCreated(c, item)
	return nil
}

func (ch *ItemHandler) ReadByID(c *fiber.Ctx) error {
	itemID, err := strconv.Atoi(c.Params("item_id"))
	if err != nil {
		utils.ResponseError(c, err.Error())
		return nil
	}

	var item Item
	err = ch.db.Preload("ItemType").First(&item, itemID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ResponseError(c, "Item not found")
		return nil
	}

	utils.ResponseDetailOK(c, item)
	return nil
}

func (ch *ItemHandler) Read(c *fiber.Ctx) error {
	var item []ItemWithTypeName
	err := ch.db.Table("items").Select("items.*, item_types.item_type_name").
		Joins("INNER JOIN item_types ON items.item_type_id = item_types.item_type_id").
		Order("items.created_at ASC").
		Find(&item).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ResponseError(c, "Item not found")
		return nil
	}

	utils.ResponseDetailOK(c, item)
	return nil
}

func (ch *ItemHandler) Update(c *fiber.Ctx) error {
	itemID, err := strconv.Atoi(c.Params("item_id"))
	if err != nil {
		utils.ResponseError(c, err.Error())
		return nil
	}

	var existingCustomer Item
	err = ch.db.First(&existingCustomer, itemID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ResponseError(c, "Item not found")
		return nil
	}

	var updatedCustomerData Item
	if err := c.BodyParser(&updatedCustomerData); err != nil {
		utils.ResponseError(c, err.Error())
		return nil
	}

	errUpdate := ch.db.Model(&existingCustomer).Updates(updatedCustomerData).Error
	if errUpdate != nil {
		msgError := errUpdate.Error()
		if strings.Contains(msgError, "Duplicate entry") {
			utils.ResponseError(c, "Item Data Duplicate")
			return nil
		} else {
			utils.ResponseError(c, msgError)
			return nil
		}
	}

	utils.ResponseUpdated(c, existingCustomer)
	return nil
}

func (ch *ItemHandler) Delete(c *fiber.Ctx) error {
	itemID, err := strconv.Atoi(c.Params("item_id"))
	if err != nil {
		utils.ResponseError(c, err.Error())
		return nil
	}

	var existingCustomer Item
	err = ch.db.First(&existingCustomer, itemID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ResponseError(c, "Item not found")
		return nil
	}

	errDelete := ch.db.Delete(&existingCustomer).Error
	if errDelete != nil {
		utils.ResponseError(c, errDelete.Error())
		return nil
	}

	utils.ResponseOK(c, "Item deleted successfully")
	return nil
}

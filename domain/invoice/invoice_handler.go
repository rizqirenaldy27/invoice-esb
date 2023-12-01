package invoice

import (
	"fmt"
	"strconv"

	"github.com/rizqirenaldy27/invoice-esb/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type InvoiceHandler struct {
	db *gorm.DB
}

func NewInvoiceHandler(db *gorm.DB) *InvoiceHandler {
	var invoiceHandler = InvoiceHandler{}
	invoiceHandler.db = db
	return &invoiceHandler
}

func (ch *InvoiceHandler) Create(c *fiber.Ctx) error {
	var input InvoiceInput
	if err := c.BodyParser(&input); err != nil {
		utils.ResponseError(c, err.Error())
		return nil
	}

	tx := ch.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	newInvoice := Invoice{
		IssueDate:  utils.ParseDate(input.IssueDate),
		DueDate:    utils.ParseDate(input.DueDate),
		CustomerID: input.CustomerID,
		Subject:    input.Subject,
		Status:     "UNPAID",
		TotalItems: input.TotalItems,
		SubTotal:   input.SubTotal,
		Tax:        input.Tax,
		GrandTotal: input.GrandTotal,
	}

	if err := tx.Create(&newInvoice).Error; err != nil {
		tx.Rollback()
		utils.ResponseError(c, "Failed to create invoice")
		return nil
	}

	for _, item := range input.Items {
		newItem := InvoiceItem{
			InvoiceID: newInvoice.InvoiceID,
			ItemID:    item.ItemID,
			Quantity:  item.Quantity,
			Amount:    item.Amount,
		}
		if err := tx.Create(&newItem).Error; err != nil {
			tx.Rollback()
			utils.ResponseError(c, "Failed to create invoice item")
			return nil
		}
	}

	tx.Commit()

	InvoiceIDString := fmt.Sprintf("%05d", newInvoice.InvoiceID)

	output := InvoiceOutput{
		InvoiceID:  InvoiceIDString,
		IssueDate:  newInvoice.IssueDate.Local().Format("2006-01-02"),
		DueDate:    newInvoice.DueDate.Local().Format("2006-01-02"),
		CustomerID: newInvoice.CustomerID,
		Subject:    newInvoice.Subject,
		Status:     newInvoice.Status,
		TotalItems: newInvoice.TotalItems,
		SubTotal:   newInvoice.SubTotal,
		Tax:        newInvoice.Tax,
		GrandTotal: newInvoice.GrandTotal,
	}

	for _, item := range input.Items {
		outputItem := InvoiceItemOutput{
			ItemID:    item.ItemID,
			Quantity:  item.Quantity,
			UnitPrice: item.UnitPrice,
			Amount:    item.Amount,
		}
		output.Items = append(output.Items, outputItem)
	}

	utils.ResponseCreated(c, output)
	return nil
}

func (ch *InvoiceHandler) ReadByID(c *fiber.Ctx) error {
	invoiceID, err := strconv.Atoi(c.Params("invoice_id"))
	if err != nil {
		utils.ResponseError(c, err.Error())
		return nil
	}

	var invoice Invoice
	if err := ch.db.First(&invoice, invoiceID).Error; err != nil {
		utils.ResponseError(c, "Invoice not found")
		return nil
	}

	var invoiceItems []InvoiceItem
	if err := ch.db.Where("invoice_id = ?", invoiceID).Find(&invoiceItems).Error; err != nil {
		utils.ResponseError(c, "Failed to retrieve invoice items")
		return nil
	}

	InvoiceIDString := fmt.Sprintf("%05d", invoice.InvoiceID)

	output := InvoiceOutput{
		InvoiceID:  InvoiceIDString,
		IssueDate:  invoice.IssueDate.Local().Format("2006-01-02"),
		DueDate:    invoice.DueDate.Local().Format("2006-01-02"),
		CustomerID: invoice.CustomerID,
		Subject:    invoice.Subject,
		Status:     invoice.Status,
		TotalItems: invoice.TotalItems,
		SubTotal:   invoice.SubTotal,
		Tax:        invoice.Tax,
		GrandTotal: invoice.GrandTotal,
		Items:      make([]InvoiceItemOutput, 0),
	}

	for _, item := range invoiceItems {
		unitPrice := item.Amount / float64(item.Quantity)
		outputItem := InvoiceItemOutput{
			InvoiceItemID: item.InvoiceItemID,
			ItemID:        item.ItemID,
			Quantity:      item.Quantity,
			UnitPrice:     unitPrice,
			Amount:        item.Amount,
		}
		output.Items = append(output.Items, outputItem)
	}

	utils.ResponseDetailOK(c, output)
	return nil
}

func (ch *InvoiceHandler) Read(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		utils.ResponseError(c, "Invalid page value")
		return nil
	}

	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil || limit < 1 {
		utils.ResponseError(c, "Invalid limit value")
		return nil
	}

	offset := (page - 1) * limit
	var filter InvoiceFilter
	filter.Subject = c.Query("subject")
	filter.InvoiceID = c.Query("invoice_id")
	filter.CustomerID = c.Query("customer_id")
	filter.TotalItems = c.Query("total_item")
	filter.IssueDate = c.Query("issue_date")
	filter.DueDate = c.Query("due_date")
	filter.Status = c.Query("status")

	// Membuat kondisi tambahan berdasarkan filter
	db := ch.db.Offset(offset).Limit(limit)
	if filter.Subject != "" {
		db = db.Where("subject LIKE ?", "%"+filter.Subject+"%")
	}
	if filter.InvoiceID != "" {
		invoiceIDFilter, _ := strconv.Atoi(filter.InvoiceID)
		db = db.Where("invoice_id = ?", invoiceIDFilter)
	}
	if filter.CustomerID != "" {
		custIDFilter, _ := strconv.Atoi(filter.CustomerID)
		db = db.Where("customer_id = ?", custIDFilter)
	}
	if filter.TotalItems != "" {
		totalItemFilter, _ := strconv.Atoi(filter.TotalItems)
		db = db.Where("total_items = ?", totalItemFilter)
	}
	if filter.IssueDate != "" {
		issueDateFilter := utils.ParseDate(filter.IssueDate)
		db = db.Where("issue_date = ?", issueDateFilter)
	}
	if filter.DueDate != "" {
		dueDateFilter := utils.ParseDate(filter.DueDate)
		db = db.Where("issue_date = ?", dueDateFilter)
	}
	if filter.Status != "" {
		db = db.Where("status = ?", filter.Status)
	}

	var invoices []Invoice
	if err := db.Find(&invoices).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve invoices"})
	}

	var outputList []InvoiceOutputRead
	for _, invoice := range invoices {
		InvoiceIDString := fmt.Sprintf("%05d", invoice.InvoiceID)

		output := InvoiceOutputRead{
			InvoiceID:  InvoiceIDString,
			IssueDate:  invoice.IssueDate.Local().Format("2006-01-02"),
			DueDate:    invoice.DueDate.Local().Format("2006-01-02"),
			CustomerID: invoice.CustomerID,
			Subject:    invoice.Subject,
			Status:     invoice.Status,
			TotalItems: invoice.TotalItems,
			SubTotal:   invoice.SubTotal,
			Tax:        invoice.Tax,
			GrandTotal: invoice.GrandTotal,
		}

		outputList = append(outputList, output)
	}

	utils.ResponseCreated(c, outputList)
	return nil
}

func (ch *InvoiceHandler) Update(c *fiber.Ctx) error {
	invoiceID, err := strconv.Atoi(c.Params("invoice_id"))
	if err != nil {
		utils.ResponseError(c, err.Error())
		return nil
	}

	var existingInvoice Invoice
	if err := ch.db.First(&existingInvoice, invoiceID).Error; err != nil {
		utils.ResponseError(c, "Invoice not found")
		return nil
	}

	// Parse data input dari request
	var input InvoiceInput
	if err := c.BodyParser(&input); err != nil {
		utils.ResponseError(c, "Invalid input format")
		return nil
	}

	// Mulai transaksi
	tx := ch.db.Begin()

	// Defer rollback jika terjadi panic
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	existingInvoice.IssueDate = utils.ParseDate(input.IssueDate)
	existingInvoice.DueDate = utils.ParseDate(input.DueDate)
	existingInvoice.CustomerID = input.CustomerID
	existingInvoice.Subject = input.Subject
	existingInvoice.TotalItems = input.TotalItems
	existingInvoice.SubTotal = input.SubTotal
	existingInvoice.Tax = input.Tax
	existingInvoice.GrandTotal = input.GrandTotal

	if err := tx.Save(&existingInvoice).Error; err != nil {
		tx.Rollback()
		utils.ResponseError(c, "Failed to update invoice")
		return nil
	}

	for _, item := range input.Items {
		var existingItem InvoiceItem

		if item.InvoiceItemID > 0 {
			// Jika invoice_item_id sudah ada, update item yang ada
			if err := tx.First(&existingItem, item.InvoiceItemID).Error; err != nil {
				tx.Rollback()
				utils.ResponseError(c, "Failed to find invoice item")
				return nil
			}

			existingItem.ItemID = item.ItemID
			existingItem.Quantity = item.Quantity
			existingItem.Amount = item.Amount

			if err := tx.Save(&existingItem).Error; err != nil {
				tx.Rollback()
				utils.ResponseError(c, "Failed to update invoice item")
				return nil
			}
		} else {
			// Jika invoice_item_id tidak ada, tambahkan item baru
			newItem := InvoiceItem{
				InvoiceID: existingInvoice.InvoiceID,
				ItemID:    item.ItemID,
				Quantity:  item.Quantity,
				Amount:    item.Amount,
			}

			if err := tx.Create(&newItem).Error; err != nil {
				tx.Rollback()
				utils.ResponseError(c, "Failed to create invoice item")
				return nil
			}
		}
	}

	tx.Commit()

	var invoice Invoice
	if err := ch.db.First(&invoice, invoiceID).Error; err != nil {
		utils.ResponseError(c, "Invoice not found")
		return nil
	}

	var invoiceItems []InvoiceItem
	if err := ch.db.Where("invoice_id = ?", invoiceID).Find(&invoiceItems).Error; err != nil {
		utils.ResponseError(c, "Failed to retrieve invoice items")
		return nil
	}

	InvoiceIDString := fmt.Sprintf("%05d", invoice.InvoiceID)

	output := InvoiceOutput{
		InvoiceID:  InvoiceIDString,
		IssueDate:  invoice.IssueDate.Local().Format("2006-01-02"),
		DueDate:    invoice.DueDate.Local().Format("2006-01-02"),
		CustomerID: invoice.CustomerID,
		Subject:    invoice.Subject,
		Status:     invoice.Status,
		TotalItems: invoice.TotalItems,
		SubTotal:   invoice.SubTotal,
		Tax:        invoice.Tax,
		GrandTotal: invoice.GrandTotal,
		Items:      make([]InvoiceItemOutput, 0),
	}

	for _, item := range invoiceItems {
		unitPrice := item.Amount / float64(item.Quantity)
		outputItem := InvoiceItemOutput{
			InvoiceItemID: item.InvoiceItemID,
			ItemID:        item.ItemID,
			Quantity:      item.Quantity,
			UnitPrice:     unitPrice,
			Amount:        item.Amount,
		}
		output.Items = append(output.Items, outputItem)
	}

	utils.ResponseDetailOK(c, output)
	return nil
}

func (ch *InvoiceHandler) Delete(c *fiber.Ctx) error {
	invoiceID, err := strconv.Atoi(c.Params("invoice_id"))
	if err != nil {
		utils.ResponseError(c, err.Error())
		return nil
	}

	var invoice Invoice
	if err := ch.db.First(&invoice, invoiceID).Error; err != nil {
		utils.ResponseError(c, "Invoice not found")
		return nil
	}

	tx := ch.db.Begin()

	if err := tx.Where("invoice_id = ?", invoiceID).Delete(&InvoiceItem{}).Error; err != nil {
		tx.Rollback()
		utils.ResponseError(c, "Failed to delete invoice items")
		return nil
	}

	// Hapus faktur
	if err := tx.Delete(&invoice).Error; err != nil {
		tx.Rollback()
		utils.ResponseError(c, "Failed to delete invoice")
		return nil
	}

	// Commit transaksi
	tx.Commit()

	utils.ResponseOK(c, "Invoice deleted successfully")
	return nil
}

package http

import (
	"github.com/rizqirenaldy27/invoice-esb/domain/customer"
	"github.com/rizqirenaldy27/invoice-esb/domain/item"
	"github.com/rizqirenaldy27/invoice-esb/domain/item_type"
	"github.com/rizqirenaldy27/invoice-esb/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

func CORSMiddleware() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Origin, Content-Type, Accept",
		ExposeHeaders:    "",
		AllowCredentials: true,
	})
}

func InitRouter(db *gorm.DB) *fiber.App {
	app := fiber.New()

	customerHandler := customer.NewCustomerHandler(db)
	customers := app.Group("/v1/customer")
	customers.Use(CORSMiddleware())
	{
		customers.Post("/", customerHandler.Create)
		customers.Get("/:customer_id", customerHandler.ReadByID)
		customers.Get("/", customerHandler.Read)
		customers.Put("/:customer_id", customerHandler.Update)
		customers.Delete("/:customer_id", customerHandler.Delete)
	}

	itemTypeHandler := item_type.NewItemTypeHandler(db)
	itemTypes := app.Group("/v1/item-type")
	itemTypes.Use(CORSMiddleware())
	{
		itemTypes.Post("/", itemTypeHandler.Create)
		itemTypes.Get("/:item_type_id", itemTypeHandler.ReadByID)
		itemTypes.Get("/", itemTypeHandler.Read)
		itemTypes.Put("/:item_type_id", itemTypeHandler.Update)
		itemTypes.Delete("/:item_type_id", itemTypeHandler.Delete)
	}

	itemHandler := item.NewItemHandler(db)
	items := app.Group("/v1/item")
	items.Use(CORSMiddleware())
	{
		items.Post("/", itemHandler.Create)
		items.Get("/:item_id", itemHandler.ReadByID)
		items.Get("/", itemHandler.Read)
		items.Put("/:item_id", itemHandler.Update)
		items.Delete("/:item_id", itemHandler.Delete)
	}

	app.Use(func(c *fiber.Ctx) error {
		utils.ResponseErrWithCode(c, "Method Not Allowed", fiber.StatusMethodNotAllowed)
		return nil
	})

	return app
}

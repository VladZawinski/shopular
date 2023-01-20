package handlers

import (
	"shopular/middlewares"
	"shopular/services"

	"github.com/gofiber/fiber/v2"
)

func SetUpHandlers(app *fiber.App, sp *services.Provider) {
	api := app.Group("/api")
	categoryRoutes(api, sp.Category)
	subCategoryRoutes(api, sp.SubCategory)
	productRoutes(api, sp.Product)
	uploadRoutes(api)
	staticRoutes(app)
	authRoutes(api, sp.User)
	cartRoutes(api, sp.Cart)
	userRoutes(api, sp.User)
	customerRoutes(api, sp.Customer)
}

func categoryRoutes(router fiber.Router, service services.CategoryService) {
	category := router.Group("/category")
	ch := NewCategoryHandler(&service)
	category.Get("/all", ch.GetAll)
	category.Use(middlewares.JwtGuard())
	category.Use(middlewares.RoleGuard(middlewares.AdminRole))
	category.Post("/create", ch.Create)
}

func productRoutes(router fiber.Router, service services.ProductService) {
	product := router.Group("/product")
	ph := NewProductHandler(&service)
	product.Post("/create", ph.Create)
	product.Get("/all", ph.GetAll)
}

func subCategoryRoutes(router fiber.Router, service services.SubCategoryService) {
	subCategory := router.Group("/sub-category")
	sch := NewSubCategoryHandler(&service)
	subCategory.Get("/all", sch.GetAll)
	subCategory.Post("/create", sch.Create)
}

func uploadRoutes(router fiber.Router) {
	upload := router.Group("file-upload")
	upload.Use(middlewares.JwtGuard())
	upload.Use(middlewares.RoleGuard(middlewares.AdminRole))
	uh := NewUploadHandler()
	upload.Post("/product", uh.UploadProductImage)
}

func authRoutes(router fiber.Router, service services.UserService) {
	auth := router.Group("auth")
	ah := NewAuthHandler(service)
	auth.Post("/signup", ah.Register)
	auth.Post("/login", ah.Login)
}

func cartRoutes(router fiber.Router, service services.CartService) {
	cart := router.Group("/cart")
	cart.Use(middlewares.JwtGuard())
	sch := NewCartHandler(service)
	cart.Get("/", sch.ViewCart)
	cart.Post("/add/:productId", sch.AddToCart)
	cart.Post("/remove/:cartId", sch.RemoveFromCart)
}

func userRoutes(router fiber.Router, service services.UserService) {
	user := router.Group("/user")
	user.Use(middlewares.JwtGuard())
	user.Use(middlewares.RoleGuard(middlewares.AdminRole))
	uh := NewUserHandler(&service)
	user.Get("/all", uh.FindAllUser)
}

func customerRoutes(router fiber.Router, service services.CustomerService) {
	customer := router.Group("/customer")
	customer.Use(middlewares.JwtGuard())
	ch := CustomerHandler{service: &service}
	customer.Get("/all", ch.FindAll)
}

func staticRoutes(app *fiber.App) {
	app.Static("/img", "./images")
}

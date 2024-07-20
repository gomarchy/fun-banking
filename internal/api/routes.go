package api

import (
	"html/template"
	"time"

	"github.com/bytebury/fun-banking/internal/api/handler"
	"github.com/bytebury/fun-banking/internal/api/middleware"
	"github.com/bytebury/fun-banking/internal/utils"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func Start() {
	router.SetFuncMap(template.FuncMap{
		"currency": func(amount float64) string { return utils.FormatCurrency(amount) },
		"sub":      func(a, b int) int { return a - b },
		"add":      func(a, b int) int { return a + b },
		"mul":      func(a, b int) int { return a * b },
		"mulfloat": func(a, b float64) float64 { return a * b },
		"datetime": func(dateTime time.Time) string { return dateTime.Format("July 02, 2006 at 3:04 PM") },
	})
	// Load Templates
	router.LoadHTMLGlob("templates/**/*")
	// Load Static Files
	router.Static("/static", "public/")
	// Middleware
	router.Use(middleware.Audit())
	// Setup Routes
	setupHomepageRoutes()
	setupActionRoutes()
	setupUserRoutes()
	setupSessionRoutes()
	setupBankRoutes()
	setupCustomerRoutes()
	setupAccountRoutes()
	setupTransactionRoutes()
	// Run the application
	router.Run()
}

func setupHomepageRoutes() {
	handler := handler.NewHomePageHandler()

	router.
		GET("", handler.Homepage)
}

func setupSessionRoutes() {
	handler := handler.NewSessionHandler()

	router.
		GET("signin", middleware.NoAuth(), handler.SignIn).
		POST("signin", middleware.NoAuth(), handler.CreateSession).
		DELETE("signout", handler.DestroySession)
}

func setupUserRoutes() {
	handler := handler.NewUserHandler()

	router.GET("signup", middleware.NoAuth(), handler.SignUp)

	router.
		Group("users").
		PUT("", handler.Create)
}

func setupBankRoutes() {
	handler := handler.NewBankHandler()

	router.
		Group("banks").
		GET("", middleware.Auth(), handler.MyBanks).
		PUT("", middleware.Auth(), handler.CreateBank).
		// TODO: PATCH should take in an ID
		PATCH("", middleware.Auth(), handler.UpdateBank).
		GET("modal-create", middleware.Auth(), handler.CreateModal).
		GET(":username/:slug", middleware.Auth(), handler.ViewBank).
		GET("settings", middleware.Auth(), handler.Settings)
}

func setupCustomerRoutes() {
	handler := handler.NewCustomerHandler()

	router.
		Group("customers").
		PUT("", middleware.Auth(), handler.CreateCustomer).
		// TODO: move this to the banks route!
		GET("modal-create", middleware.Auth(), handler.OpenCreateModal).
		GET(":id", middleware.Auth(), handler.GetCustomer).
		PATCH(":id", middleware.Auth(), handler.Update).
		GET(":id/settings", middleware.Auth(), handler.OpenSettings)
}

func setupAccountRoutes() {
	handler := handler.NewAccountHandler()

	router.
		Group("accounts").
		GET(":id", middleware.Auth(), handler.Get).
		PATCH(":id", middleware.Auth(), handler.Update).
		GET(":id/transactions", middleware.Auth(), handler.GetTransactions).
		GET(":id/settings", middleware.Auth(), handler.OpenSettings).
		GET(":id/cash-flow", middleware.Auth(), handler.CashFlow).
		GET(":id/withdraw-or-deposit", middleware.Auth(), handler.OpenWithdrawOrDeposit).
		PUT(":id/withdraw-or-deposit", middleware.Auth(), handler.WithdrawOrDeposit).
		GET(":id/send-money", middleware.Auth(), handler.OpenSendMoneyModal)
}

func setupTransactionRoutes() {
	handler := handler.NewTransactionHandler()

	router.
		Group("transactions").
		PUT("", middleware.Auth(), handler.Create)
}

func setupActionRoutes() {
	handler := handler.NewActionHandler()

	router.Group("actions").
		GET("open-app-drawer", middleware.Audit(), handler.OpenAppDrawer)
}

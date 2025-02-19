package main

import (
	authController "crud_fire/controller"
	productController "crud_fire/controller"
	firebase "crud_fire/database"
	"crud_fire/repository"
	authService "crud_fire/service"
	productService "crud_fire/service"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// inisialisasi firebase
	firestoreClient := firebase.InitFirebase().FirestoreClient
	authClient := firebase.InitFirebase().AuthClient
	apiKeys := os.Getenv("API_KEYS")
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	senderEmail := os.Getenv("SENDER_EMAIL")

	// konversi port ke int
	smtpPortInt, err := strconv.Atoi(smtpPort)
	if err != nil {
		log.Fatal("Error konversi port ke int", err)
	}
	// product
	repo := repository.NewProductRepository(firestoreClient)
	service := productService.NewProductService(repo)
	controller := productController.NewProductController(service)

	// auth
	authService := authService.NewServiceAuth(authClient, apiKeys, smtpUsername, smtpPassword, smtpHost, senderEmail, smtpPortInt)
	authController := authController.NewAuthController(authService)

	// routing
	router := echo.New()

	// method
	router.GET("/product/get", controller.GetAllProducts)
	router.POST("/product/post", controller.InsertProduct)
	router.PUT("/product/put", controller.EditProduct)
	router.DELETE("/product/delete", controller.DeleteProduct)

	router.POST("/auth/login", authController.Login)
	router.POST("/auth/register", authController.Register)
	router.GET("/auth/verify-email", authController.VerifyEmail)

	// run server
	router.Start(":8080")

}

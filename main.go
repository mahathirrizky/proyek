package main

import (
	"log"
	"net/http"
	"proyek/adminproyek"
	"proyek/auth"
	"proyek/handlers"
	"proyek/helper"
	"proyek/material"
	"proyek/pembelian"
	"proyek/proyek"
	"proyek/user"
	"strings"


	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Replace your MySQL connection details here
	dsn := "root:@tcp(localhost:3306)/proyek?charset=utf8mb4&parseTime=True&loc=Local"

	// Initialize GORM
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Auto Migrate the models
	err = db.AutoMigrate(&user.UserTable{}, &adminproyek.AdminProyekTable{}, &material.MaterialTable{}, &material.StokTable{}, &proyek.ProyekTable{}, &pembelian.PembelianTable{})
	if err != nil {
		log.Fatalf("Error auto migrating models: %v", err)
	}

	// Initialize Gin router
	router := gin.Default()

	// Initialize repositories
	userRepo := user.NewUserRepository(db)
	adminProyekRepo := adminproyek.NewAdminProyekRepository(db)
	materialRepo := material.NewMaterialRepository(db)
	proyekRepo := proyek.NewRepository(db)
	pembelianRepo := pembelian.NewPembelianRepository(db)

	// Initialize services with repositories
	userService := user.NewUserService(userRepo)
	adminProyekService := adminproyek.NewAdminProyekService(adminProyekRepo)
	materialService := material.NewMaterialService(materialRepo)
	proyekService := proyek.NewService(proyekRepo)
	pembelianService := pembelian.NewPembelianService(pembelianRepo)
	authService := auth.NewService()

	// Initialize handlers and wire them with services
	userHandler := handlers.NewUserHandler(userService, authService)
	adminProyekHandler := handlers.NewAdminProyekHandler(adminProyekService)
	materialHandler := handlers.NewMaterialHandler(materialService)
	proyekHandler := handlers.NewProyekHandler(proyekService)
	pembelianHandler := handlers.NewPembelianHandler(pembelianService)

	// Routes
	api := router.Group("/api")
	{
		// User routes
		api.POST("/users", userHandler.CreateUser)
		api.PUT("/users/:id", userHandler.UpdateUser)
		api.POST("/login", userHandler.Login)
		api.GET("/users/fetch", authMiddleware(authService, userService), userHandler.FetchUser)
		api.GET("/users/:id", userHandler.GetUserByID)

		// AdminProyek routes
		api.POST("/adminproyek", adminProyekHandler.CreateAdminProyek)
		api.PUT("/adminproyek/:id", adminProyekHandler.UpdateAdminProyek)
		api.DELETE("/adminproyek/:id", adminProyekHandler.DeleteAdminProyek)
		api.GET("/adminproyek/:id", adminProyekHandler.GetAdminProyekByID)
		api.GET("/adminproyek", adminProyekHandler.GetAllAdminProyeks)

		// Material routes
		api.POST("/material", materialHandler.CreateMaterial)
		api.GET("/material", materialHandler.GetAllMaterials)
		api.GET("/material/:id", materialHandler.GetMaterialByID) 
		api.PUT("/material/:id", materialHandler.UpdateMaterial) 


		// Proyek routes
		api.POST("/proyek", proyekHandler.CreateProyek)
		api.PUT("/proyek/:id", proyekHandler.UpdateProyek)
		api.DELETE("/proyek/:id", proyekHandler.DeleteProyek)
		api.GET("/proyek/:id", proyekHandler.GetProyekByID)
		api.GET("/proyek", proyekHandler.GetAllProyeks)

		// Pembelian routes
		api.POST("/pembelian", pembelianHandler.CreatePembelian)
		api.PUT("/pembelian/:id", pembelianHandler.UpdatePembelian)
		api.GET("/pembelian/:id", pembelianHandler.GetPembelianByID)
		api.GET("/pembelian/proyek/:proyek_id", pembelianHandler.GetPembelianByProyekID)
	}

	// Start Gin server
	router.Run(":8080")
}
func authMiddleware(authService auth.Service, userService user.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Retrieve userID from claim
		userIDFloat64, ok := claim["user_id"].(float64)
		if !ok {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userID := int(userIDFloat64)

		print(userID)

		// Retrieve user from userService
		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

			if c.Request.Method == "OPTIONS" {
					c.AbortWithStatus(204)
					return
			}

			c.Next()
	}
}

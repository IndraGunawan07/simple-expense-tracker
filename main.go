package main

import (
	"database/sql"
	"log"
	"os"

	"expense-tracker/controllers"
	"expense-tracker/database"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	// err := godotenv.Load("config/.env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	dbURL := os.Getenv("DATABASE_URL")

	// psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
	// 	os.Getenv("PGHOST"),
	// 	os.Getenv("PGPORT"),
	// 	os.Getenv("PGUSER"),
	// 	os.Getenv("PGPASSWORD"),
	// 	os.Getenv("PGDATABASE"),
	// )

	DB, err = sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatalf("Error opening database: %v\n", err)
	}

	defer DB.Close()
	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	database.DBMigrate(DB)

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	router.SetTrustedProxies(nil)

	// public routes
	router.GET("/", func(c *gin.Context) {
		log.Println("Root route hit")
		c.JSON(200, gin.H{"status": "ok"})
	})
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	// Protected routes
	auth := router.Group("/")
	auth.Use(controllers.AuthMiddleware())
	{
		auth.GET("/user", controllers.GetUser)

		auth.GET("/categories", controllers.GetAllCategory)
		auth.POST("/categories", controllers.InsertCategory)
		auth.PUT("/categories/:id", controllers.UpdateCategory)
		auth.DELETE("/categories/:id", controllers.DeleteCategory)

		auth.GET("/expenses", controllers.GetAllExpense)
		auth.POST("/expenses", controllers.InsertExpense)
		auth.PUT("/expenses/:id", controllers.UpdateExpense)
		auth.DELETE("/expenses/:id", controllers.DeleteExpense)

		auth.GET("/reports", controllers.GetReport)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Environment variables:")
	for _, e := range os.Environ() {
		log.Println(e)
	}

	// Start server with proper error handling
	log.Println("Starting server at port:", port)
	err = router.Run("0.0.0.0:" + port)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
	log.Println("Server has stopped unexpectedly")
}

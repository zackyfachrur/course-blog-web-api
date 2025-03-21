package main

import (
	"fmt"
	"log"
	"myproject/src/db"
	"myproject/src/routes"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()
	database, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}
	defer database.Close()

	godotenv.Load("../.env.local")
	port := os.Getenv("COURSE_VIDEO_BE_PORT")
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Izinkan akses dari frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour, // Cache CORS selama 12 jam
	}))
	fmt.Println("Server is running on Port : ", port)
	routes.SetupRoutes(router, database)
	router.Run(port)
}

// cmd/main.go
package main

import (
	"auth-simple/controllers"
	"auth-simple/models"
	"log"

	"github.com/gin-gonic/gin"
	sqlite "gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "modernc.org/sqlite" // ✅ required to use modern SQLite driver
)

func main() {
	// Init SQLite
	db, err := gorm.Open(sqlite.Dialector{
		DriverName: "sqlite", // 👈 must be set
		DSN:        "auth.db",
	}, &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.User{})

	r := gin.Default()

	r.POST("/register", controllers.Register(db))
	r.POST("/login", controllers.Login(db))
	r.POST("/change-password", controllers.ChangePassword(db))

	r.Run(":8080")
}

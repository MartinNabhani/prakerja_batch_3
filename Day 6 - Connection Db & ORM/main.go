package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
	InitialMigration()
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {
	config := Config{
		DB_Username: "root",
		DB_Password: "12321",
		DB_Port:     "3306",
		DB_Host:     "127.0.0.1",
		DB_Name:     "crud_go",
	}
	connectionString :=
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			config.DB_Username,
			config.DB_Password,
			config.DB_Host,
			config.DB_Port,
			config.DB_Name,
		)
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func InitialMigration() {
	DB.AutoMigrate(&User{})
}
func GetUsersController(c echo.Context) error {
	var users []User
	if err := DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

func getUserByID(c echo.Context) error {

	// ambil data user berdasarkan id
	var user User
	if err := DB.First(&user, c.Param("id")).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	// return data dalam format JSON
	return c.JSON(http.StatusOK, user)
}

// create new user
func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)
	if err := DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

func DeleteUserController(c echo.Context) error {
	id := c.Param("id")
	var user User
	result := DB.First(&user, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}
	DB.Delete(&user)
	return c.JSON(http.StatusOK, "User deleted successfully")
}

func main() {
	// create a new echo instance
	e := echo.New()
	// Route / to handler function
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", getUserByID)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}

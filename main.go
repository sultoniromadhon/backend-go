package main

import (
	"crowdfunding/user"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// dsn := "root:@tcp(127.0.0.1:3306)/crowdfunding_db?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// fmt.Println("Koneksi Berhasil")

	// var users []user.User
	// length := len(users)
	// fmt.Println(length)
	// db.Find(&users)
	// length = len(users)
	// fmt.Println(length)

	// for _, user := range users {
	// 	fmt.Println(user.Email, user.Role, user.ID, user.PasswordHash)
	// }
	// router := gin.Default()
	// router.GET("/handler", handlercontrol)
	// router.Run()

	router := gin.Default()
	router.GET("/", handlercontrol)
	router.Run()
}

func handlercontrol(c *gin.Context) {
	dsn := "root:@tcp(127.0.0.1:3306)/crowdfunding_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	var users []user.User
	db.Find(&users)
	c.JSON(http.StatusOK, users)

}

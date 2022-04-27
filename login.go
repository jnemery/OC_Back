package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

var (
	currUser User
	once     sync.Once
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
	Valid    bool   `json:"valid"`
}

func init() {
	once.Do(initialiseUser)
}

func initialiseUser() {
	currUser = User{}
}

func validateUser(c *gin.Context) {
	hours, minutes, _ := time.Now().Clock()
	currUTCTimeInString := fmt.Sprintf("%d%02d", hours, minutes)
	fmt.Println(currUTCTimeInString)
	validatedUser := User{Username: "c137@onecause.com", Password: "#th@nH@rm#y#r!$100%D0p#", Token: currUTCTimeInString}

	c.Header("Access-Control-Allow-Origin", "http://localhost:4200")

	if err := c.ShouldBindJSON(&currUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currUser.Valid = false

	if currUser.Username == validatedUser.Username &&
		currUser.Password == validatedUser.Password &&
		currUser.Token == validatedUser.Token {
		currUser.Valid = true
		c.JSON(http.StatusOK, gin.H{"valid": currUser.Valid})
		return
	}

	c.JSON(http.StatusOK, gin.H{"valid": currUser.Valid})

}

func main() {
	router := gin.Default()

	router.Use(cors.Default())

	router.POST("/login", validateUser)

	router.Run(":8080")
}

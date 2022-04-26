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
	user          User
	authenticated bool
	mtx           sync.RWMutex
	once          sync.Once
)

func init() {
	once.Do(initialiseUser)
}

func initialiseUser() {
	user = User{}
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
	Valid    bool   `json:"valid"`
}

func Check(user string, password string, token string) bool {

	hours, minutes, _ := time.Now().Clock()
	currUTCTimeInString := fmt.Sprintf("%d%02d", hours, minutes)
	fmt.Println(currUTCTimeInString)
	if user == "c137@onecause.com" && password == "#th@nH@rm#y#r!$100%D0p#" && token == currUTCTimeInString {
		authenticated = true
		fmt.Println("here")
	}
	return authenticated
}

func validateUser(c *gin.Context) {
	var newUser User
	hours, minutes, _ := time.Now().Clock()
	currUTCTimeInString := fmt.Sprintf("%d%02d", hours, minutes)
	fmt.Println(currUTCTimeInString)
	validatedUser := User{Username: "c137@onecause.com", Password: "#th@nH@rm#y#r!$100%D0p#", Token: currUTCTimeInString}

	c.Header("Access-Control-Allow-Origin", "http://localhost:4200")

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser.Valid = false

	if newUser.Username == validatedUser.Username &&
		newUser.Password == validatedUser.Password &&
		newUser.Token == validatedUser.Token {
		newUser.Valid = true
		c.JSON(http.StatusOK, gin.H{"valid": newUser.Valid})
		return
	}

	c.JSON(http.StatusOK, gin.H{"valid": newUser.Valid})

}

func main() {
	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Hello": "World"})
	})

	router.POST("/login", validateUser)

	router.Run(":8080")

	/*hours, minutes, _ := time.Now().Clock()
	currUTCTimeInString := fmt.Sprintf("%d%02d", hours, minutes)
	fmt.Println(currUTCTimeInString)
	♂Check("c137@onecause.com", "#th@nH@rm#y#r!$100%D0p#", currUTCTimeInString)
	*/
}

package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	user          Login
	authenticated bool
	mtx           sync.RWMutex
	once          sync.Once
)

func init() {
	once.Do(initialiseUser)
}

func initialiseUser() {
	user = Login{}
}

type Login struct {
	username string `json:"username"`
	password string `json:"password"`
	token    string `json:"token"`
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

func main() {
	hours, minutes, _ := time.Now().Clock()
	currUTCTimeInString := fmt.Sprintf("%d%02d", hours, minutes)
	fmt.Println(currUTCTimeInString)
	Check("c137@onecause.com", "#th@nH@rm#y#r!$100%D0p#", currUTCTimeInString)
}

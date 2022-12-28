package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	PersonModel "webserverwithGo/Models"
	Packages "webserverwithGo/packages"
)

func SignUp(c *gin.Context) {
	db := Packages.GetMyDb()
	defer db.Close()
	var newPerson PersonModel.Person
	if err := c.BindJSON(&newPerson); err != nil {
		c.Status(400)
		panic(err)
	}
	_, err := db.InsertOne(newPerson)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusCreated, newPerson)
}

func GetAllUsers(c *gin.Context) {
	db := Packages.GetMyDb()
	defer db.Close()
	var users []PersonModel.Person = db.FindAll()
	c.IndentedJSON(http.StatusOK, users)
}

func UpdateUser(c *gin.Context) {
	db := Packages.GetMyDb()
	defer db.Close()
	var newPerson PersonModel.Person

	if err := c.BindJSON(&newPerson); err != nil {
		var errorOutput []Packages.ErrorMsg = Packages.CustomizeErrors(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errorOutput})
		return
	}
	newPerson = db.FindOne(newPerson.FirstName)

	fmt.Println(newPerson)
	c.IndentedJSON(http.StatusOK, gin.H{"msg": newPerson})
}

func main() {
	r := gin.Default()
	r.POST("/sign-up", SignUp)
	r.GET("/users", GetAllUsers)
	r.PATCH("/users", UpdateUser)

	r.Run()
}

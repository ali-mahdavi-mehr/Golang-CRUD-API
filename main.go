package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	PersonModel "webserverwithGo/Models"
	mydatabase "webserverwithGo/packages"
)

func SignUp(c *gin.Context) {
	db := mydatabase.GetMyDb()
	defer db.Close()
	var newPerson PersonModel.Person
	if err := c.BindJSON(&newPerson); err != nil {
		c.Status(400)
		panic(err)
	}
	c.IndentedJSON(http.StatusCreated, newPerson)
	db.InsertOne(newPerson)
}

func main() {
	r := gin.Default()
	r.POST("/sign-up", SignUp)
	r.Run()
}

package main

import (
	"fmt"
	"github.com/pluralsight/webservice/models"
)

func main() {
	u := models.User{
		ID: 2,
		Firstname: "Arthur",
		LastName: "Dane",
	}
	fmt.Println(u)
}
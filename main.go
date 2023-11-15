package main

import (
	"fmt"
	"net/http"

	"anshulgithub.com/anshul/usermangement/controller"
)

func main() {
	fmt.Println("welcome to go lang")
	router := controller.GetContrller()
	http.ListenAndServe(":4000", &router)
}

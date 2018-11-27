package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	//router := httprouter.New()
	http.ListenAndServe(":8000", RegisterHandlers())
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name",Login)
	return router
}

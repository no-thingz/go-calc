package main

import (
	"log"
	"net/http"

	"github.com/no-thingz/go-calc/calculatorserver/controllers"
)

func main() {
	log.Printf("Calculator Server Starting, Listening on port%v", controllers.GetListenAddress())
	controllers.RegisterControllers()
	if err := http.ListenAndServe(controllers.GetListenAddress(), nil); err != nil {
		panic(err)
	}
}

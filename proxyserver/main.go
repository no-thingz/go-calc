package main

import (
	"log"
	"net/http"

	"github.com/no-thingz/go-calc/proxyserver/controllers"
)

func main() {
	log.Printf("Proxy Server Starting, Listening on port%v", controllers.GetListenAddress())
	controllers.RegisterControllers()
	if err := http.ListenAndServe(controllers.GetListenAddress(), nil); err != nil {
		panic(err)
	}
}

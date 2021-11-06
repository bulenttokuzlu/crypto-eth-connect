package main

import (
	"log"

	//	_ "expvar"         //profiler
	//	"net/http"         //profiler
	//	_ "net/http/pprof" //profiler
	"github.com/bulenttokuzlu/crypto-eth-connect/controller"
	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/eth/v1/transaction/:id", controller.GetTxnInfo)
}

func main() {
	// profiler - OPTION 1: Add this to your application
	//go func() {
	//	http.ListenAndServe("0.0.0.0:6060", nil)
	//}()
	//	app.Initialize()

	app := fiber.New()
	setupRoutes(app)
	err := app.Listen(":8080")
	if err != nil {
		log.Panic(err)
	}

}

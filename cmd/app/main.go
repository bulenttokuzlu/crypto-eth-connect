package main

import (
	"log"

	//	_ "expvar"         //profiler
	//	"net/http"         //profiler
	//	_ "net/http/pprof" //profiler
	"crypto-eth.com/connect/controller"
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

/*type GetTxnResponse struct {
	Status string `json:"status"` //Running, Completed, Error, Success
	Hash   string `json:"hash"`
}

func GetTxnInfo(c *fiber.Ctx) error {
	conn, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatalf("Hadi len")
	}
	ctx := context.Background()
	tx, pending, _ := conn.TransactionByHash(ctx, common.HexToHash(c.Params("id")))
	if !pending {
		c.JSON(tx)
	} else {
		resp := GetTxnResponse{Status: "Pending"}
		resp.Hash = c.Params("id")
		c.JSON(resp)
	}
	return nil
}
*/

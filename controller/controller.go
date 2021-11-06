package controller

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gofiber/fiber"
)

type GetTxnResponse struct {
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

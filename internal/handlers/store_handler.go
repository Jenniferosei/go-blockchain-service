package handlers

import (
    "context"
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/Jenniferosei/go-blockchain-service/internal/db"
    "github.com/Jenniferosei/go-blockchain-service/internal/services"
)

type StoreHandler struct {
    eth *services.EthService
}

func NewStoreHandler(eth *services.EthService) *StoreHandler {
    return &StoreHandler{eth: eth}
}

func (h *StoreHandler) StoreLatestBlock(c *gin.Context) {
    ctx := context.Background()

    blockNumber, err := h.eth.GetLatestBlock()
    if err != nil {
        log.Println("Error getting latest block:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get latest block number"})
        return
    }

    // Fetch block HEADER only (no transactions)
    header, err := h.eth.Client.HeaderByNumber(ctx, nil)
    if err != nil {
        log.Println("Error fetching block header:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch block header"})
        return
    }

    err = db.StoreBlockHeader(ctx, header)
    if err != nil {
        log.Println("Error storing block:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to store block"})
        return
    }

    log.Printf("Stored block number %d with hash %s\n", blockNumber, header.Hash().Hex())
    c.JSON(http.StatusOK, gin.H{"message": "block stored successfully", "blockNumber": blockNumber})
}

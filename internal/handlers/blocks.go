package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/Jenniferosei/go-blockchain-service/internal/services"
)

type BlockHandler struct {
    Eth *services.EthService
}

func NewBlockHandler(eth *services.EthService) *BlockHandler {
    return &BlockHandler{Eth: eth}
}

func (h *BlockHandler) LatestBlock(c *gin.Context) {
    block, err := h.Eth.GetLatestBlock()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "latest_block": block,
    })
}

package main

import (
    "github.com/gin-gonic/gin"
    "github.com/prometheus/client_golang/prometheus/promhttp"

    "github.com/Jenniferosei/go-blockchain-service/internal/db"
    "github.com/Jenniferosei/go-blockchain-service/internal/handlers"
    "github.com/Jenniferosei/go-blockchain-service/internal/metrics"
    "github.com/Jenniferosei/go-blockchain-service/internal/services"
    "os"
)

func main() {
    // Env vars
    rpcURL := os.Getenv("ETH_RPC_URL")
    postgresURL := os.Getenv("POSTGRES_URL")

    // Services
    pg, err := db.NewPostgres(postgresURL)
    if err != nil {
        panic(err)
    }
    defer pg.Pool.Close()

    eth, err := services.NewEthService(rpcURL)
    if err != nil {
        panic(err)
    }

    // Init metrics
    metrics.InitMetrics()

    // Router
    r := gin.Default()

    // API routes
    blockHandler := handlers.NewBlockHandler(eth)

    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

    r.GET("/latest-block", func(c *gin.Context) {
        metrics.RequestsTotal.WithLabelValues("/latest-block").Inc()
        blockHandler.LatestBlock(c)
    })

    // Prometheus endpoint
    r.GET("/metrics", gin.WrapH(promhttp.Handler()))

    // Start server
    r.Run(":8080")
}

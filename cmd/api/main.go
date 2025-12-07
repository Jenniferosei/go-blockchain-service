// package main

// import (
//     "log"
//     "os"

//     "github.com/gin-gonic/gin"
//     "github.com/joho/godotenv"
//     "github.com/prometheus/client_golang/prometheus/promhttp"

//     "github.com/Jenniferosei/go-blockchain-service/internal/db"
//     "github.com/Jenniferosei/go-blockchain-service/internal/handlers"
//     "github.com/Jenniferosei/go-blockchain-service/internal/metrics"
//     "github.com/Jenniferosei/go-blockchain-service/internal/services"
// )

// func main() {

//     // Load environment variables from .env
//     if err := godotenv.Load(); err != nil {
//         log.Println("Warning: .env file not found, using system environment variables")
//     }

//     // ENV variables
//     rpcURL := os.Getenv("ETH_RPC_URL")       // Infura / Alchemy RPC
//     postgresURL := os.Getenv("DATABASE_URL") // Postgres DB URL

//     // Init Postgres
//     pg, err := db.NewPostgres(postgresURL)
//     if err != nil {
//         panic(err)
//     }
//     defer pg.Close()

//     // Init Ethereum RPC client
//     eth, err := services.NewEthService(rpcURL)
//     if err != nil {
//         panic(err)
//     }

//     // Init Prometheus metrics
//     metrics.InitMetrics()

//     // Router
//     r := gin.Default()

//     // Handlers
//     blockHandler := handlers.NewBlockHandler(eth)
//     storeHandler := handlers.NewStoreHandler(eth)

//     // Health check
//     r.GET("/health", func(c *gin.Context) {
//         c.JSON(200, gin.H{"status": "ok"})
//     })

//     // GET latest block (does NOT store in DB)
//     r.GET("/latest-block", func(c *gin.Context) {
//         metrics.RequestsTotal.WithLabelValues("/latest-block").Inc()
//         blockHandler.LatestBlock(c)
//     })

//     // POST: fetch + store latest block
//     r.POST("/store-latest-block", func(c *gin.Context) {
//         metrics.RequestsTotal.WithLabelValues("/store-latest-block").Inc()
//         storeHandler.StoreLatestBlock(c)
//     })

//     // Prometheus metrics
//     r.GET("/metrics", gin.WrapH(promhttp.Handler()))

//     // Start API server
//     port := os.Getenv("PORT")
//     if port == "" {
//         port = "8080"
//     }

//     log.Println("API running on port:", port)
//     r.Run(":" + port)
// }


package main

import (
    "context"
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/Jenniferosei/go-blockchain-service/internal/db"
    "github.com/Jenniferosei/go-blockchain-service/internal/handlers"
    "github.com/Jenniferosei/go-blockchain-service/internal/metrics"
    "github.com/Jenniferosei/go-blockchain-service/internal/services"
)


func main() {
    if err := godotenv.Load(); err != nil {
        log.Println("Warning: .env file not found, using system environment variables")
    }

    rpcURL := os.Getenv("ETH_RPC_URL")
    postgresURL := os.Getenv("DATABASE_URL")

    pgPool, err := pgxpool.New(context.Background(), postgresURL)
    if err != nil {
        panic(err)
    }
    defer pgPool.Close()

    db.InitDB(pgPool)

    ethService, err := services.NewEthService(rpcURL)
    if err != nil {
        panic(err)
    }

    metrics.InitMetrics()

    r := gin.Default()

    storeHandler := handlers.NewStoreHandler(ethService)

    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

    r.POST("/store-latest-block", storeHandler.StoreLatestBlock)

    r.GET("/metrics", gin.WrapH(promhttp.Handler()))

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Println("API running on port:", port)
    r.Run(":" + port)
}

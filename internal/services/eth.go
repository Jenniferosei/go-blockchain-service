package services

import (
    "context"
    "fmt"
    "log"
    "math/big"

    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/Jenniferosei/go-blockchain-service/internal/db"
)

type EthService struct {
    Client *ethclient.Client
}

func NewEthService(rpcURL string) (*EthService, error) {
    client, err := ethclient.Dial(rpcURL)
    if err != nil {
        return nil, err
    }
    return &EthService{Client: client}, nil
}

func (e *EthService) GetLatestBlock() (uint64, error) {
    header, err := e.Client.HeaderByNumber(context.Background(), nil)
    if err != nil {
        return 0, err
    }
    return header.Number.Uint64(), nil
}

func (e *EthService) GetClientVersion() string {
    v, err := e.Client.NetworkID(context.Background())
    if err != nil {
        log.Println("error getting network ID:", err)
        return "unknown"
    }
    return v.String()
}

// FetchAndStoreBlock fetches a block by number from the Ethereum node and stores it in Postgres.
func (e *EthService) FetchAndStoreBlock(ctx context.Context, blockNumber int64) (*types.Block, error) {
    block, err := e.Client.BlockByNumber(ctx, big.NewInt(blockNumber))
    if err != nil {
        return nil, fmt.Errorf("failed to fetch block: %w", err)
    }

    err = db.StoreBlock(ctx, block)
    if err != nil {
        return nil, fmt.Errorf("failed to store block in DB: %w", err)
    }

    return block, nil
}

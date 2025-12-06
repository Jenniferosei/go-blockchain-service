package services

import (
    "context"
    "github.com/ethereum/go-ethereum/ethclient"
    "log"
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

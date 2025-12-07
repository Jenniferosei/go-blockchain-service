package db

import (
    "context"
    "log"
    "time"

    "github.com/ethereum/go-ethereum/core/types"
)

func StoreBlock(ctx context.Context, block *types.Block) error {
    log.Printf("Storing block #%d\n", block.Number().Int64())

    _, err := DB.Exec(ctx,
        `INSERT INTO blocks (number, hash, parent_hash, timestamp) 
         VALUES ($1, $2, $3, $4)
         ON CONFLICT (number) DO NOTHING`,
        block.Number().Int64(),
        block.Hash().Hex(),
        block.ParentHash().Hex(),
        time.Unix(int64(block.Time()), 0),
    )
    if err != nil {
        log.Printf("Error inserting block: %v\n", err)
        return err
    }

    log.Printf("Block %d stored successfully\n", block.Number().Int64())
    return nil
}

// New function for headers only
func StoreBlockHeader(ctx context.Context, header *types.Header) error {
    log.Printf("Storing block header #%d\n", header.Number.Int64())

    _, err := DB.Exec(ctx,
        `INSERT INTO blocks (number, hash, parent_hash, timestamp) 
         VALUES ($1, $2, $3, $4)
         ON CONFLICT (number) DO NOTHING`,
        header.Number.Int64(),
        header.Hash().Hex(),
        header.ParentHash.Hex(),
        time.Unix(int64(header.Time), 0),
    )
    if err != nil {
        log.Printf("Error inserting block: %v\n", err)
        return err
    }

    log.Printf("Block %d stored successfully\n", header.Number.Int64())
    return nil
}
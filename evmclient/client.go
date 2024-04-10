package evmclient

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var ErrNoConfiguredPRC = errors.New("no configured rpc")

type client struct {
	client *ethclient.Client
}

func NewClient(c *ethclient.Client) *client {
	return &client{client: c}
}

type Header struct {
	Hash       common.Hash `json:"hash"`
	ParentHash common.Hash `json:"parentHash"`
	Number     *big.Int    `json:"number"`
	Time       uint64      `json:"timestamp"`
}

func (c *client) BlockNumber(ctx context.Context) (uint64, error) {
	return c.client.BlockNumber(ctx)
}

func (c *client) HeaderByNumber(ctx context.Context, number *big.Int) (Header, error) {
	header, err := c.client.HeaderByNumber(ctx, number)
	return convertHeader(header), err
}

func (c *client) HeaderByHash(ctx context.Context, hash common.Hash) (Header, error) {
	header, err := c.client.HeaderByHash(ctx, hash)
	return convertHeader(header), err
}

func (c *client) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	return c.client.FilterLogs(ctx, q)
}

func convertHeader(header *types.Header) Header {
	return Header{
		Hash:       header.Hash(),
		ParentHash: header.ParentHash,
		Number:     header.Number,
		Time:       header.Time,
	}
}

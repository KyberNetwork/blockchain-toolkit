package avalanche

import (
	"context"
	"errors"
	"math/big"
	"net/http"
	"time"

	avaxtypes "github.com/ava-labs/coreth/core/types"
	avaxclient "github.com/ava-labs/coreth/ethclient"
	"github.com/ava-labs/coreth/interfaces"
	"github.com/ava-labs/coreth/rpc"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/KyberNetwork/blockchain-toolkit/evmclient"
)

var ErrNoConfiguredPRC = errors.New("no configured rpc")

type client struct {
	avaxClient avaxclient.Client
}

func NewClientWithTimeout(rpcs []string, timeout time.Duration) (*client, error) {
	hc := &http.Client{Timeout: timeout}
	for _, url := range rpcs {
		rc, err := rpc.DialOptions(context.Background(), url, rpc.WithHTTPClient(hc))
		if err != nil {
			continue
		}

		ec := avaxclient.NewClient(rc)

		return &client{avaxClient: ec}, nil
	}

	return nil, ErrNoConfiguredPRC
}

func (c *client) BlockNumber(ctx context.Context) (uint64, error) {
	return c.avaxClient.BlockNumber(ctx)
}

func (c *client) HeaderByNumber(ctx context.Context, number *big.Int) (evmclient.Header, error) {
	header, err := c.avaxClient.HeaderByNumber(ctx, number)
	return convertAvaxHeader(header), err
}

func (c *client) HeaderByHash(ctx context.Context, hash common.Hash) (evmclient.Header, error) {
	header, err := c.avaxClient.HeaderByHash(ctx, hash)
	return convertAvaxHeader(header), err
}

func (c *client) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	ecq := interfaces.FilterQuery(q)

	logs, err := c.avaxClient.FilterLogs(ctx, ecq)

	if err != nil {
		return nil, err
	}

	evmLogs := []types.Log{}
	for _, log := range logs {
		evmLogs = append(evmLogs, types.Log(log))
	}

	return evmLogs, nil
}

func convertAvaxHeader(header *avaxtypes.Header) evmclient.Header {
	// Only convert fields that we use.
	return evmclient.Header{
		Hash:       header.Hash(),
		ParentHash: header.ParentHash,
		Number:     header.Number,
		Time:       header.Time,
	}
}

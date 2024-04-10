package zksync

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/KyberNetwork/blockchain-toolkit/evmclient"
)

var ErrNoConfiguredPRC = errors.New("no configured rpc")

type client struct {
	ethClient *ethclient.Client
	rpcClient *rpc.Client
}

func NewClientWithTimeout(rpcs []string, timeout time.Duration) (*client, error) {
	hc := &http.Client{Timeout: timeout}
	for _, url := range rpcs {
		rc, err := rpc.DialOptions(context.Background(), url, rpc.WithHTTPClient(hc))
		if err != nil {
			continue
		}

		ec := ethclient.NewClient(rc)

		return &client{ethClient: ec, rpcClient: rc}, nil
	}
	return nil, ErrNoConfiguredPRC
}

type Header struct {
	Hash common.Hash `json:"hash"`
	types.Header
}

func (h *Header) UnmarshalJSON(data []byte) error {
	type Header struct {
		Hash            *common.Hash      `json:"hash"`
		ParentHash      *common.Hash      `json:"parentHash"`
		UncleHash       *common.Hash      `json:"sha3Uncles"`
		Coinbase        *common.Address   `json:"miner"`
		Root            *common.Hash      `json:"stateRoot"`
		TxHash          *common.Hash      `json:"transactionsRoot"`
		ReceiptHash     *common.Hash      `json:"receiptsRoot"`
		Bloom           *types.Bloom      `json:"logsBloom"`
		Difficulty      *hexutil.Big      `json:"difficulty"`
		Number          *hexutil.Big      `json:"number"`
		GasLimit        *hexutil.Uint64   `json:"gasLimit"`
		GasUsed         *hexutil.Uint64   `json:"gasUsed"`
		Time            *hexutil.Uint64   `json:"timestamp"`
		Extra           *hexutil.Bytes    `json:"extraData"`
		MixDigest       *common.Hash      `json:"mixHash"`
		Nonce           *types.BlockNonce `json:"nonce"`
		BaseFee         *hexutil.Big      `json:"baseFeePerGas"`
		WithdrawalsHash *common.Hash      `json:"withdrawalsRoot"`
	}

	var dec Header
	if err := json.Unmarshal(data, &dec); err != nil {
		return err
	}

	if dec.Hash == nil {
		return errors.New("missing required field 'hash' for Header")
	}
	h.Hash = *dec.Hash

	if dec.ParentHash == nil {
		return errors.New("missing required field 'parentHash' for Header")
	}
	h.ParentHash = *dec.ParentHash

	if dec.UncleHash == nil {
		return errors.New("missing required field 'sha3Uncles' for Header")
	}
	h.UncleHash = *dec.UncleHash

	if dec.Coinbase != nil {
		h.Coinbase = *dec.Coinbase
	}
	if dec.Root == nil {
		return errors.New("missing required field 'stateRoot' for Header")
	}
	h.Root = *dec.Root

	if dec.TxHash == nil {
		return errors.New("missing required field 'transactionsRoot' for Header")
	}
	h.TxHash = *dec.TxHash

	if dec.ReceiptHash == nil {
		return errors.New("missing required field 'receiptsRoot' for Header")
	}
	h.ReceiptHash = *dec.ReceiptHash

	if dec.Bloom == nil {
		return errors.New("missing required field 'logsBloom' for Header")
	}
	h.Bloom = *dec.Bloom

	if dec.Difficulty == nil {
		return errors.New("missing required field 'difficulty' for Header")
	}
	h.Difficulty = (*big.Int)(dec.Difficulty)

	if dec.Number == nil {
		return errors.New("missing required field 'number' for Header")
	}
	h.Number = (*big.Int)(dec.Number)

	if dec.GasLimit == nil {
		return errors.New("missing required field 'gasLimit' for Header")
	}
	h.GasLimit = uint64(*dec.GasLimit)

	if dec.GasUsed == nil {
		return errors.New("missing required field 'gasUsed' for Header")
	}
	h.GasUsed = uint64(*dec.GasUsed)

	if dec.Time == nil {
		return errors.New("missing required field 'timestamp' for Header")
	}
	h.Time = uint64(*dec.Time)

	if dec.Extra == nil {
		return errors.New("missing required field 'extraData' for Header")
	}
	h.Extra = *dec.Extra

	if dec.MixDigest != nil {
		h.MixDigest = *dec.MixDigest
	}
	if dec.Nonce != nil {
		h.Nonce = *dec.Nonce
	}
	if dec.BaseFee != nil {
		h.BaseFee = (*big.Int)(dec.BaseFee)
	}
	if dec.WithdrawalsHash != nil {
		h.WithdrawalsHash = dec.WithdrawalsHash
	}

	return nil
}

func (c *client) BlockNumber(ctx context.Context) (uint64, error) {
	return c.ethClient.BlockNumber(ctx)
}

func (c *client) HeaderByNumber(ctx context.Context, number *big.Int) (evmclient.Header, error) {
	var head *Header
	err := c.rpcClient.CallContext(ctx, &head, "eth_getBlockByNumber", toBlockNumArg(number), false)
	if err == nil && head == nil {
		err = ethereum.NotFound
	}

	return evmclient.Header{
		Hash:       head.Hash,
		ParentHash: head.ParentHash,
		Number:     head.Number,
		Time:       head.Time,
	}, err
}

func (c *client) HeaderByHash(ctx context.Context, hash common.Hash) (evmclient.Header, error) {
	var head *Header
	err := c.rpcClient.CallContext(ctx, &head, "eth_getBlockByHash", hash, false)
	if err == nil && head == nil {
		err = ethereum.NotFound
	}

	return evmclient.Header{
		Hash:       head.Hash,
		ParentHash: head.ParentHash,
		Number:     head.Number,
		Time:       head.Time,
	}, err
}

func (c *client) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	return c.ethClient.FilterLogs(ctx, q)
}

func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	if number.Sign() >= 0 {
		return hexutil.EncodeBig(number)
	}
	// It's negative.
	if number.IsInt64() {
		return rpc.BlockNumber(number.Int64()).String()
	}
	// It's negative and large, which is invalid.
	return fmt.Sprintf("<invalid %d>", number)
}

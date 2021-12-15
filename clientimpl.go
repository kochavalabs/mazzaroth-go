package mazzaroth

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/pkg/errors"
)

const version = "v1"

var _ Client = &ClientImpl{}

// ClientImpl is the actual client implementation.
type ClientImpl struct {
	httpClient *http.Client
	address    string
}

// NewMazzarothClient creates a production object.
func NewMazzarothClient(options ...Options) (*ClientImpl, error) {
	clientOptions := defaultOption()

	// set all options if supplied
	for _, opt := range options {
		opt.apply(clientOptions)
	}

	return &ClientImpl{
		httpClient: clientOptions.httpClient,
		address:    clientOptions.address,
	}, nil
}

// AccountLookup calls the endpoint: /v1/channels/{channel_id}/accounts/{account_id}.
func (c *ClientImpl) AccountLookup(ctx context.Context, channelID string, accountID string) (*xdr.Account, error) {
	url := fmt.Sprintf("%s/%s/channels/%s/accounts/%s", c.address, version, channelID, accountID)

	xdrResp, err := c.do(ctx, url, http.MethodGet, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to handle http response")
	}

	if xdrResp.Account != nil {
		return xdrResp.Account, nil
	}

	return nil, errors.New("Missing account")
}

// AuthorizationLookup calls the endpoint: /v1/channels/{channel_id}/accounts/{account_id}/authorized
func (c *ClientImpl) AuthorizationLookup(ctx context.Context, channelID string, accountID string) (*xdr.Authorized, error) {
	url := fmt.Sprintf("%s/%s/channels/%s/accounts/%s/authorized", c.address, version, channelID, accountID)

	xdrResp, err := c.do(ctx, url, http.MethodGet, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to handle http response")
	}

	if xdrResp.Authorized != nil {
		return xdrResp.Authorized, nil
	}
	return nil, errors.New("missing authorized accounts")
}

// BlockHeight calls the endpoint: /v1/channels/{channel_id}/blocks/height.
func (c *ClientImpl) BlockHeight(ctx context.Context, channelID string) (*xdr.BlockHeight, error) {
	url := fmt.Sprintf("%s/%s/channels/%s/blocks/height", c.address, version, channelID)

	xdrResp, err := c.do(ctx, url, http.MethodGet, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to handle http response")
	}

	if xdrResp.Height != nil {
		return xdrResp.Height, nil
	}

	return nil, errors.New("Missing block height")
}

// BlockLookup calls the endpoint: /v1/channels/{channel_id}/blocks/{id}.
func (c *ClientImpl) BlockLookup(ctx context.Context, channelID, blockID string) (*xdr.Block, error) {
	url := fmt.Sprintf("%s/%s/channels/%s/blocks/%s", c.address, version, channelID, blockID)

	xdrResp, err := c.do(ctx, url, http.MethodGet, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to handle http response")
	}

	if xdrResp.Block != nil {
		return xdrResp.Block, nil
	}

	return nil, errors.New("Missing block")
}

// BlockList calls the endpoint: /v1/channels/{channel_id}/blocks?{number,height}.
func (c *ClientImpl) BlockList(ctx context.Context, channelID string, blockHeight int, number int) ([]*xdr.Block, error) {
	url := fmt.Sprintf("%s/%s/channels/%s/blocks?height=%d&number=%d", c.address, version, channelID, blockHeight, number)

	xdrResp, err := c.do(ctx, url, http.MethodGet, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to handle http response")
	}

	if xdrResp.Blocks != nil {
		return xdrResp.Blocks, nil
	}

	return nil, errors.New("Missing blocks")
}

// BlockHeaderLookup calls the endpoint: /v1/channels/{channel_id}/blockheaders/{id}.
func (c *ClientImpl) BlockHeaderLookup(ctx context.Context, channelID, blockID string) (*xdr.BlockHeader, error) {
	url := fmt.Sprintf("%s/%s/channels/%s/blockheaders/%s", c.address, version, channelID, blockID)

	xdrResp, err := c.do(ctx, url, http.MethodGet, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to handle http response")
	}

	if xdrResp.BlockHeader != nil {
		return xdrResp.BlockHeader, nil
	}

	return nil, errors.New("Missing blockheader")
}

// BlockHeaderList calls the endpoint: /v1/channels/{channel_id}/blockheaders?{blockHeight,number}.
func (c *ClientImpl) BlockHeaderList(ctx context.Context, channelID string, blockHeight int, number int) ([]*xdr.BlockHeader, error) {
	url := fmt.Sprintf("%s/%s/channels/%s/blockheaders?height=%d&number=%d", c.address, version, channelID, blockHeight, number)

	xdrResp, err := c.do(ctx, url, http.MethodGet, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to handle http response")
	}

	if xdrResp.BlockHeaders != nil {
		return xdrResp.BlockHeaders, nil
	}

	return nil, errors.New("Missing blockHeaders")
}

// ChannelLookup calls the endpoint: /v1/channels/{channel_id}/config.
func (c *ClientImpl) ChannelConfig(ctx context.Context, channelID string) (*xdr.Config, error) {
	url := fmt.Sprintf("%s/%s/channels/%s/config", c.address, version, channelID)

	xdrResp, err := c.do(ctx, url, http.MethodGet, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to handle http response")
	}

	if xdrResp.Config != nil {
		return xdrResp.Config, nil
	}

	return nil, errors.New("Missing channel config")
}

// ChannelAbi calls the endpoint: /v1/channels/{channel_id}/abi.
func (c *ClientImpl) ChannelAbi(ctx context.Context, channelID string) (*xdr.Abi, error) {
	url := fmt.Sprintf("%s/%s/channels/%s/abi", c.address, version, channelID)

	xdrResp, err := c.do(ctx, url, http.MethodGet, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to handle http response")
	}

	if xdrResp.Abi != nil {
		return xdrResp.Abi, nil
	}

	return nil, errors.New("Missing channel abi")
}

// ReceiptLookup calls the endpoint: /v1/channels/{channel_id}/receipts/{id}.
func (c *ClientImpl) ReceiptLookup(ctx context.Context, channelID, transactionID string) (*xdr.Receipt, error) {
	url := fmt.Sprintf("%s/%s/channels/%s/receipts/%s", c.address, version, channelID, transactionID)

	xdrResp, err := c.do(ctx, url, http.MethodGet, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to handle http response")
	}

	if xdrResp.Receipt != nil {
		return xdrResp.Receipt, nil
	}

	return nil, errors.New("Missing receipt")
}

// TransactionSubmit calls the endpoint: /v1/channels/{channel_id}/transactions.
func (c *ClientImpl) TransactionSubmit(ctx context.Context, transaction *xdr.Transaction) (*xdr.ID, *xdr.Receipt, error) {
	b, err := json.Marshal(transaction)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to marshal to json")
	}

	channelID := hex.EncodeToString(transaction.Data.ChannelID[:])
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to unmarshal the channelID")
	}

	url := fmt.Sprintf("%s/%s/channels/%s/transactions", c.address, version, channelID)

	xdrResp, err := c.do(ctx, url, http.MethodPost, b)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to make a request to transaction submit endpoint")
	}

	if xdrResp.TransactionID != nil {
		return xdrResp.TransactionID, nil, nil
	} else if xdrResp.Receipt != nil {
		return &xdrResp.Receipt.TransactionID, xdrResp.Receipt, nil
	}
	return nil, nil, errors.New("unable to process transaction")
}

// TransactionLookup calls the endpoint: /v1/channels/{channel_id}/transactions/{id}.
func (c *ClientImpl) TransactionLookup(ctx context.Context, channelID string, transactionID string) (*xdr.Transaction, error) {
	url := fmt.Sprintf("%s/%s/channels/%s/transactions/%s", c.address, version, channelID, transactionID)

	xdrResp, err := c.do(ctx, url, http.MethodGet, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to handle http response")
	}

	if xdrResp.Transaction != nil {
		return xdrResp.Transaction, nil
	}

	return nil, errors.New("missing transaction")
}

func (c *ClientImpl) do(ctx context.Context, url string, method string, body []byte) (*xdr.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(body))
	if err != nil {
		return nil, errors.Wrap(err, "unable to create a new request")
	}

	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "unable to make http request")
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status %d", response.StatusCode)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "could not read the body")
	}

	xdrResp := xdr.Response{}
	err = xdrResp.UnmarshalJSON(responseBody)
	if err != nil {
		return nil, errors.Wrap(err, "could not unmarshal the body")
	}

	return &xdrResp, nil
}

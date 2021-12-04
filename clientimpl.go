package mazzaroth

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
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
func NewMazzarothClient(servers []string, options ...Options) (*ClientImpl, error) {
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

// TransactionSubmit calls the endpoint: /v1/channels/{channel_id}/transactions.
func (c *ClientImpl) TransactionSubmit(transaction *xdr.Transaction) (*xdr.ID, *xdr.Receipt, error) {
	b, err := json.Marshal(transaction)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to marshal to json")
	}

	channelID := hex.EncodeToString(transaction.Data.ChannelID[:])
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to unmarshal the channelID")
	}

	url := fmt.Sprintf("%s/%s/channels/%s/transactions", c.address, version, channelID)

	response, err := makeRequest(c.httpClient, http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to make a request to transaction submit endpoint")
	}

	return nil, response.Receipt, nil
}

// TransactionLookup calls the endpoint: /v1/channels/{channel_id}/transactions/{id}.
func (c *ClientImpl) TransactionLookup(channelID string, transactionID string) (*xdr.Transaction, error) {
	url := fmt.Sprintf("%s/%s/channels/%s/transactions/%s", c.address, version, channelID, transactionID)

	response, err := makeRequest(c.httpClient, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to make a request to transaction lookup endpoint")
	}

	return response.Transaction, nil
}

// ReceiptLookup calls the endpoint: /v1/channels/{channel_id}/receipts/{id}.
func (c *ClientImpl) ReceiptLookup(channelID, transactionID string) (*xdr.Receipt, error) {
	url := fmt.Sprintf("%s/%s/channels/%s/receipts/%s", c.address, version, channelID, transactionID)

	response, err := makeRequest(c.httpClient, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to make a request to receipts lookup endpoint")
	}

	return response.Receipt, nil
}

// BlockLookup calls the endpoint: /v1/channels/{channel_id}/blocks/{id}.
func (c *ClientImpl) BlockLookup(channelID, blockID string) (*xdr.Block, error) {
	url := fmt.Sprintf("%s/%s/channels/%s/blocks/%s", c.address, version, channelID, blockID)

	response, err := makeRequest(c.httpClient, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to make a request to block lookup endpoint")
	}

	return response.Block, nil
}

// BlockList calls the endpoint: /v1/channels/{channel_id}/blocks?{number,height}.
func (c *ClientImpl) BlockList(channelID string, blockHeight int, number int) (*[]xdr.Block, error) {
	url := fmt.Sprintf("%s/%s/channels/%s/blocks?height=%d&number=%d", c.address, version, channelID, blockHeight, number)

	response, err := makeRequest(c.httpClient, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to make a request to block by height lookup endpoint")
	}

	return response.Blocks, nil
}

// BlockHeaderLookup calls the endpoint: /v1/channels/{channel_id}/blockheaders/{id}.
func (c *ClientImpl) BlockHeaderLookup(channelID, blockID string) (*xdr.BlockHeader, error) {
	url := fmt.Sprintf("%s/%s/channels/%s/blockheaders/%s", c.address, version, channelID, blockID)

	response, err := makeRequest(c.httpClient, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to make a request to blockheaders lookup endpoint")
	}

	return response.BlockHeader, nil
}

// BlockHeaderList calls the endpoint: /v1/channels/{channel_id}/blockheaders?{blockHeight,number}.
func (c *ClientImpl) BlockHeaderList(channelID string, blockHeight int, number int) (*[]xdr.BlockHeader, error) {
	url := fmt.Sprintf("%s/%s/channels/%s/blockheaders?height=%d&number=%d", c.address, version, channelID, blockHeight, number)

	response, err := makeRequest(c.httpClient, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to make a request to blockheaders by blockheight lookup endpoint")
	}

	return response.BlockHeaders, nil
}

// ChannelLookup calls the endpoint: /v1/channels/{channel_id}.
func (c *ClientImpl) ChannelLookup(channelID string) (*xdr.Config, error) {
	url := fmt.Sprintf("%s/%s/channels/%s", c.address, version, channelID)

	response, err := makeRequest(c.httpClient, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to make a request to channel lookup endpoint")
	}

	return response.Config, nil
}

// ChannelAbi calls the endpoint: /v1/channels/{channel_id}/abi.
func (c *ClientImpl) ChannelAbi(channelID string) (*xdr.Abi, error) {
	url := fmt.Sprintf("%s/%s/channels/%s/abi", c.address, version, channelID)

	response, err := makeRequest(c.httpClient, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to make a request to channel lookup endpoint")
	}

	return response.Abi, nil
}

// BlockHeight calls the endpoint: /v1/channels/{channel_id}/blocks/height.
func (c *ClientImpl) BlockHeight(channelID string) (*xdr.BlockHeight, error) {
	url := fmt.Sprintf("%s/%s/channels/%s/blocks/height", c.address, version, channelID)

	response, err := makeRequest(c.httpClient, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to make a request to channel lookup endpoint")
	}

	return response.Height, nil
}

// AccountLookup calls the endpoint: /v1/channels/{channel_id}/accounts/{accout_id}.
func (c *ClientImpl) AccountLookup(channelID string, seed string) (*xdr.Account, error) {
	seedBin, err := hex.DecodeString(seed)
	if err != nil {
		return nil, errors.Wrap(err, "unable to decode seed")
	}
	privateKey := ed25519.NewKeyFromSeed(seedBin)
	address, err := xdr.IDFromPublicKey(privateKey.Public())
	if err != nil {
		return nil, errors.Wrap(err, "unable to create private key")
	}

	url := fmt.Sprintf("%s/%s/channels/%s/accounts/%s", c.address, version, channelID, hex.EncodeToString(address[:]))

	response, err := makeRequest(c.httpClient, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to make a request to channel lookup endpoint")
	}

	return response.Account, nil
}

func (c *ClientImpl) AuthorizationLookup(channelID string, accountID string) (*xdr.Authorization, error) {
	return nil, nil
}

func makeRequest(httpClient *http.Client, method, url string, body io.Reader) (*xdr.Response, error) {

	req, err := http.NewRequestWithContext(context.Background(), method, url, body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create a new request")
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "unable to make http request")
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusNotFound:
		return nil, ErrNotFound
	case http.StatusInternalServerError:
		return nil, ErrInternalServer
	default:
		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "could not read the body")
		}

		return nil, fmt.Errorf("http status %d - %s", resp.StatusCode, string(responseBody))
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "could not read the body")
	}

	responseXDR := xdr.Response{}
	err = responseXDR.UnmarshalJSON(responseBody)
	if err != nil {
		return nil, errors.Wrap(err, "could not unmarshal the body")
	}

	return &responseXDR, nil
}

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

// ErrNotFound is raised when the searched entity is not found.
var ErrNotFound = errors.New("entity not found")

// ErrInternalServer is raised after a 500 status code.
var ErrInternalServer = errors.New("internal server error")

var _ Client = &ClientImpl{}

// ClientImpl is the actual client implementation.
type ClientImpl struct {
	serverSelector ServerSelector
	httpClient     *http.Client
}

// NewMazzarothClient creates a production object.
func NewMazzarothClient(servers []string, options ...Options) (*ClientImpl, error) {
	clientOptions := defaultOption()

	// set all options if supplied
	for _, opt := range options {
		opt.apply(clientOptions)
	}

	serverSelector, err := NewRoundRobinServerSelector(servers...)
	if err != nil {
		return nil, errors.Wrap(err, "could not create round robin server selector")
	}
	return &ClientImpl{
		httpClient:     clientOptions.HttpClient,
		serverSelector: serverSelector,
	}, nil
}

// TransactionSubmit calls the endpoint: /v1/channels/{channel_id}/transactions.
func (c *ClientImpl) TransactionSubmit(transaction *xdr.Transaction) (*xdr.Response, error) {
	b, err := json.Marshal(transaction)
	if err != nil {
		return nil, errors.Wrap(err, "unable to marshal to json")
	}

	channelID := hex.EncodeToString(transaction.Action.ChannelID[:])
	if err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal the channelID")
	}

	url := fmt.Sprintf("%s/%s/channels/%s/transactions", c.serverSelector.Pick(), version, channelID)

	response, err := makeRequest(c.httpClient, http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		return nil, errors.Wrap(err, "unable to make a request to transaction submit endpoint")
	}

	return response, nil
}

// TransactionLookup calls the endpoint: /v1/channels/{channel_id}/transactions/{id}.
func (c *ClientImpl) TransactionLookup(channelID string, transactionID string) (*xdr.Response, error) {
	url := fmt.Sprintf("%s/%s/channels/%s/transactions/%s", c.serverSelector.Pick(), version, channelID, transactionID)

	response, err := makeRequest(c.httpClient, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to make a request to transaction lookup endpoint")
	}

	return response, nil
}

// TransactionLookupByBlockHeight calls the endpoint: /v1/channels/{channel_id}/transactions?{blockHeight}.
func (c *ClientImpl) TransactionLookupByBlockHeight(channelID string, blockHeight int) (*xdr.Response, error) {
	url := fmt.Sprintf("%s/%s/channels/%s/transactions?blockheight=%d", c.serverSelector.Pick(), version, channelID, blockHeight)

	response, err := makeRequest(c.httpClient, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to make a request to transaction lookup by height endpoint")
	}

	return response, nil
}

// TransactionLookupByBlockID calls the endpoint: /v1/channels/{channel_id}/transactions?{blockID}.
func (c *ClientImpl) TransactionLookupByBlockID(channelID string, blockID string) (*xdr.Response, error) {
	url := fmt.Sprintf("%s/%s/channels/%s/transactions?blockid=%s", c.serverSelector.Pick(), version, channelID, blockID)

	response, err := makeRequest(c.httpClient, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to make a request to transaction lookup by blockid endpoint")
	}

	return response, nil
}

// ReceiptLookup calls the endpoint: /v1/channels/{channel_id}/receipts/{id}.
func (c *ClientImpl) ReceiptLookup(channelID, transactionID string) (*xdr.Response, error) {
	url := fmt.Sprintf("%s/%s/channels/%s/receipts/%s", c.serverSelector.Pick(), version, channelID, transactionID)

	response, err := makeRequest(c.httpClient, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to make a request to receipts lookup endpoint")
	}

	return response, nil
}

// BlockLookup calls the endpoint: /v1/channels/{channel_id}/blocks/{id}.
func (c *ClientImpl) BlockLookup(channelID, blockID string) (*xdr.Response, error) {
	url := fmt.Sprintf("%s/%s/channels/%s/blocks/%s", c.serverSelector.Pick(), version, channelID, blockID)

	response, err := makeRequest(c.httpClient, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to make a request to block lookup endpoint")
	}

	return response, nil
}

// BlockList calls the endpoint: /v1/channels/{channel_id}/blocks?{number,height}.
func (c *ClientImpl) BlockList(channelID string, blockHeight int, number int) (*xdr.Response, error) {
	url := fmt.Sprintf("%s/%s/channels/%s/blocks?height=%d&number=%d", c.serverSelector.Pick(), version, channelID, blockHeight, number)

	response, err := makeRequest(c.httpClient, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to make a request to block by height lookup endpoint")
	}

	return response, nil
}

// BlockHeaderLookup calls the endpoint: /v1/channels/{channel_id}/blockheaders/{id}.
func (c *ClientImpl) BlockHeaderLookup(channelID, blockID string) (*xdr.Response, error) {
	url := fmt.Sprintf("%s/%s/channels/%s/blockheaders/%s", c.serverSelector.Pick(), version, channelID, blockID)

	response, err := makeRequest(c.httpClient, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to make a request to blockheaders lookup endpoint")
	}

	return response, nil
}

// BlockHeaderList calls the endpoint: /v1/channels/{channel_id}/blockheaders?{blockHeight,number}.
func (c *ClientImpl) BlockHeaderList(channelID string, blockHeight int, number int) (*xdr.Response, error) {
	url := fmt.Sprintf("%s/%s/channels/%s/blockheaders?height=%d&number=%d", c.serverSelector.Pick(), version, channelID, blockHeight, number)

	response, err := makeRequest(c.httpClient, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to make a request to blockheaders by blockheight lookup endpoint")
	}

	return response, nil
}

// ChannelLookup calls the endpoint: /v1/channels/{channel_id}.
func (c *ClientImpl) ChannelLookup(channelID string) (*xdr.Response, error) {
	url := fmt.Sprintf("%s/%s/channels/%s", c.serverSelector.Pick(), version, channelID)

	response, err := makeRequest(c.httpClient, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to make a request to channel lookup endpoint")
	}

	return response, nil
}

// ChannelAbi calls the endpoint: /v1/channels/{channel_id}/abi.
func (c *ClientImpl) ChannelAbi(channelID string) (*xdr.Response, error) {
	url := fmt.Sprintf("%s/%s/channels/%s/abi", c.serverSelector.Pick(), version, channelID)

	response, err := makeRequest(c.httpClient, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to make a request to channel lookup endpoint")
	}

	return response, nil
}

// BlockHeight calls the endpoint: /v1/channels/{channel_id}/blocks/height.
func (c *ClientImpl) BlockHeight(channelID string) (*xdr.Response, error) {
	url := fmt.Sprintf("%s/%s/channels/%s/blocks/height", c.serverSelector.Pick(), version, channelID)

	response, err := makeRequest(c.httpClient, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to make a request to channel lookup endpoint")
	}

	return response, nil
}

// AccountLookup calls the endpoint: /v1/channels/{channel_id}/accounts/{accout_id}.
func (c *ClientImpl) AccountLookup(channelID string, seed string) (*xdr.Response, error) {
	seedBin, err := hex.DecodeString(seed)
	if err != nil {
		return nil, errors.Wrap(err, "unable to decode seed")
	}
	privateKey := ed25519.NewKeyFromSeed(seedBin)
	address, err := xdr.IDFromPublicKey(privateKey.Public())
	if err != nil {
		return nil, errors.Wrap(err, "unable to create private key")
	}

	url := fmt.Sprintf("%s/%s/channels/%s/accounts/%s", c.serverSelector.Pick(), version, channelID, hex.EncodeToString(address[:]))

	response, err := makeRequest(c.httpClient, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to make a request to channel lookup endpoint")
	}

	return response, nil
}

func makeRequest(httpClient *http.Client, method, url string, body io.Reader) (*xdr.Response, error) {
	fmt.Println("--------------------------------------------------------------------------------------------------------")
	fmt.Println(url)

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

	fmt.Println(string(responseBody))
	fmt.Println("--------------------------------------------------------------------------------------------------------")

	responseXDR := xdr.Response{}
	err = responseXDR.UnmarshalJSON(responseBody)
	if err != nil {
		return nil, errors.Wrap(err, "could not unmarshal the body")
	}

	return &responseXDR, nil
}

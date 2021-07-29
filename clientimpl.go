package mazzaroth

import (
	"context"
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/pkg/errors"
)

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

// TransactionSubmit calls the endpoint: /transaction/submit.
func (c *ClientImpl) TransactionSubmit(transaction xdr.Transaction) (*xdr.TransactionSubmitResponse, error) {
	transactionRequest := xdr.TransactionSubmitRequest{
		Transaction: transaction,
	}

	xdrRequest, err := transactionRequest.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "unable to marshal to xdr binary")
	}

	binaryResp, err := makeRequest(c.httpClient, c.serverSelector.Pick()+"/transaction/submit", xdrRequest)
	if err != nil {
		return nil, errors.Wrap(err, "unable to make a request to transaction submit endpoint")
	}

	response := xdr.TransactionSubmitResponse{}
	if err := response.UnmarshalBinary(binaryResp); err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal xdr response")
	}
	return &response, nil
}

// TransactionLookup calls the endpoint: /transaction/lookup.
func (c *ClientImpl) TransactionLookup(transactionID xdr.ID) (*xdr.TransactionLookupResponse, error) {
	request := xdr.TransactionLookupRequest{
		TransactionID: transactionID,
	}

	xdrRequest, err := request.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "unable to marshal to xdr binary")
	}

	binaryResp, err := makeRequest(c.httpClient, c.serverSelector.Pick()+"/transaction/lookup", xdrRequest)
	if err != nil {
		return nil, errors.Wrap(err, "unable to call transaction lookup endpoint")
	}

	response := xdr.TransactionLookupResponse{}
	if err := response.UnmarshalBinary(binaryResp); err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal xdr response")
	}
	return &response, nil
}

// ReceiptLookup calls the endpoint: /receipt/lookup.
func (c *ClientImpl) ReceiptLookup(transactionID xdr.ID) (*xdr.ReceiptLookupResponse, error) {
	request := xdr.ReceiptLookupRequest{
		TransactionID: transactionID,
	}

	xdrRequest, err := request.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "unable to marshal to xdr binary")
	}

	binaryResp, err := makeRequest(c.httpClient, c.serverSelector.Pick()+"/receipt/lookup", xdrRequest)
	if err != nil {
		return nil, errors.Wrap(err, "unable to call the receipt lookup endpoint")
	}

	response := xdr.ReceiptLookupResponse{}
	if err := response.UnmarshalBinary(binaryResp); err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal xdr response")
	}
	return &response, nil
}

// BlockLookup calls the endpoint: /block/lookup.
func (c *ClientImpl) BlockLookup(blockID xdr.Identifier) (*xdr.BlockLookupResponse, error) {
	request := xdr.BlockLookupRequest{
		Identifier: blockID,
	}

	xdrRequest, err := request.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "unable to marshal to xdr binary")
	}

	binaryResp, err := makeRequest(c.httpClient, c.serverSelector.Pick()+"/block/lookup", xdrRequest)
	if err != nil {
		return nil, errors.Wrap(err, "unable to call block lookup endpoint")
	}

	response := xdr.BlockLookupResponse{}
	if err := response.UnmarshalBinary(binaryResp); err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal xdr response")
	}
	return &response, nil
}

// BlockHeaderLookup calls the endpoint: /block/header/lookup.
func (c *ClientImpl) BlockHeaderLookup(blockID xdr.Identifier) (*xdr.BlockHeaderLookupResponse, error) {
	request := xdr.BlockHeaderLookupRequest{
		Identifier: blockID,
	}

	xdrRequest, err := request.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "unable to marshal to xdr binary")
	}

	binaryResp, err := makeRequest(c.httpClient, c.serverSelector.Pick()+"/block/header/lookup", xdrRequest)
	if err != nil {
		return nil, errors.Wrap(err, "unable to call block header lookup endpoint")
	}

	response := xdr.BlockHeaderLookupResponse{}
	if err := response.UnmarshalBinary(binaryResp); err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal xdr response")
	}
	return &response, nil
}

// AccountInfoLookup calls the endpoint: /account/info/lookup.
func (c *ClientImpl) AccountInfoLookup(accountID xdr.ID) (*xdr.AccountInfoLookupResponse, error) {
	request := xdr.AccountInfoLookupRequest{
		Account: accountID,
	}

	xdrRequest, err := request.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "unable to marshal to xdr binary")
	}

	binaryResp, err := makeRequest(c.httpClient, c.serverSelector.Pick()+"/account/info/lookup", xdrRequest)
	if err != nil {
		return nil, errors.Wrap(err, "unable to call account info lookup endpoint")
	}

	response := xdr.AccountInfoLookupResponse{}
	if err := response.UnmarshalBinary(binaryResp); err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal xdr response")
	}
	return &response, nil
}

// ChannelInfoLookup calls the endpoint: /channel/info/lookup.
func (c *ClientImpl) ChannelInfoLookup(channelInfoType xdr.ChannelInfoType) (*xdr.ChannelInfoLookupResponse, error) {
	request := xdr.ChannelInfoLookupRequest{
		InfoType: channelInfoType,
	}

	xdrRequest, err := request.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "unable to marshal to xdr binary")
	}

	binaryResp, err := makeRequest(c.httpClient, c.serverSelector.Pick()+"/channel/info/lookup", xdrRequest)
	if err != nil {
		return nil, errors.Wrap(err, "unable to call channel info lookup endpoint")
	}

	response := xdr.ChannelInfoLookupResponse{}
	if err := response.UnmarshalBinary(binaryResp); err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal xdr response")
	}
	return &response, nil
}

// BlockHeightLookup retrieves the current block height
func (c *ClientImpl) BlockHeightLookup() (uint64, error) {
	url := c.serverSelector.Pick() + "/block/height"
	binaryResp, err := makeRequest(c.httpClient, url, nil)
	if err != nil {
		return 0, errors.Wrap(err, "unable to call block height lookup endpoint")
	}
	blockHeight, err := strconv.ParseUint(string(binaryResp), 10, 64)
	return blockHeight, errors.Wrap(err, "unable to convert block height response to uint64")
}

func makeRequest(httpClient *http.Client, url string, xdrRequest []byte) ([]byte, error) {
	b64request := base64.StdEncoding.WithPadding(base64.StdPadding).EncodeToString(xdrRequest)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, url, strings.NewReader(b64request))
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

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Wrap(err, "status code is not OK")
	}

	b64Resp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "could not read the body")
	}

	binaryResp, err := base64.StdEncoding.DecodeString(string(b64Resp))
	if err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal binary response")
	}
	return binaryResp, nil
}

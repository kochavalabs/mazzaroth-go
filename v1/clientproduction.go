package mazzaroth

import (
	"context"
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/pkg/errors"
)

var _ Mazzaroth = &ProductionClient{}

// ProductionClient is the actual client implementation.
type ProductionClient struct {
	server     string
	httpClient http.Client
}

// NewProductionClient creates a production object.
func NewProductionClient(httpClient http.Client) *ProductionClient {
	return &ProductionClient{server: "http://localhost:8081"}
}

// NewProductionClientWithDefaultHTTPClient creates a production object.
func NewProductionClientWithDefaultHTTPClient() *ProductionClient {
	client := http.Client{
		Timeout: 500 * time.Millisecond,
	}
	return NewProductionClient(client)
}

func (pc *ProductionClient) TransactionSubmit(transaction xdr.Transaction) (*xdr.TransactionSubmitResponse, error) {
	transactionRequest := xdr.TransactionSubmitRequest{
		Transaction: transaction,
	}

	xdrRequest, err := transactionRequest.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "could not marshall to xdr binary")
	}

	binaryResp, err := makeRequest(pc.httpClient, pc.server+"/transaction/submit", xdrRequest)
	if err != nil {
		return nil, errors.Wrap(err, "could not call the endpoint")
	}

	transactionResponse := xdr.TransactionSubmitResponse{}

	err = transactionResponse.UnmarshalBinary(binaryResp)
	if err != nil {
		return nil, errors.Wrap(err, "could not unmarshal the response")
	}

	return &transactionResponse, nil
}

func (pc *ProductionClient) ReadOnly(function string, parameters ...xdr.Parameter) (*xdr.ReadonlyResponse, error) {
	request := xdr.ReadonlyRequest{
		Call: xdr.Call{
			Function:   function,
			Parameters: parameters,
		},
	}

	xdrRequest, err := request.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "could not marshall to xdr binary")
	}

	binaryResp, err := makeRequest(pc.httpClient, pc.server+"/readonly", xdrRequest)
	if err != nil {
		return nil, errors.Wrap(err, "could not call the endpoint")
	}

	response := xdr.ReadonlyResponse{}

	err = response.UnmarshalBinary(binaryResp)
	if err != nil {
		return nil, errors.Wrap(err, "could not unmarshal the response")
	}

	return &response, nil
}
func (pc *ProductionClient) TransactionLookup(transactionID xdr.ID) (*xdr.TransactionLookupResponse, error) {
	request := xdr.TransactionLookupRequest{
		TransactionID: transactionID,
	}

	xdrRequest, err := request.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "could not marshall to xdr binary")
	}

	binaryResp, err := makeRequest(pc.httpClient, pc.server+"/transaction/lookup", xdrRequest)
	if err != nil {
		return nil, errors.Wrap(err, "could not call the endpoint")
	}

	response := xdr.TransactionLookupResponse{}

	err = response.UnmarshalBinary(binaryResp)
	if err != nil {
		return nil, errors.Wrap(err, "could not unmarshal the response")
	}

	return &response, nil
}

func (pc *ProductionClient) ReceiptLookup(transactionID xdr.ID) (*xdr.ReceiptLookupResponse, error) {
	request := xdr.ReceiptLookupRequest{
		TransactionID: transactionID,
	}

	xdrRequest, err := request.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "could not marshall to xdr binary")
	}

	binaryResp, err := makeRequest(pc.httpClient, pc.server+"/receipt/lookup", xdrRequest)
	if err != nil {
		return nil, errors.Wrap(err, "could not call the endpoint")
	}

	response := xdr.ReceiptLookupResponse{}

	err = response.UnmarshalBinary(binaryResp)
	if err != nil {
		return nil, errors.Wrap(err, "could not unmarshal the response")
	}

	return &response, nil
}

func (pc *ProductionClient) BlockLookup(blockID xdr.Identifier) (*xdr.BlockLookupResponse, error) {
	request := xdr.BlockLookupRequest{
		ID: blockID,
	}

	xdrRequest, err := request.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "could not marshall to xdr binary")
	}

	binaryResp, err := makeRequest(pc.httpClient, pc.server+"/block/lookup", xdrRequest)
	if err != nil {
		return nil, errors.Wrap(err, "could not call the endpoint")
	}

	response := xdr.BlockLookupResponse{}

	err = response.UnmarshalBinary(binaryResp)
	if err != nil {
		return nil, errors.Wrap(err, "could not unmarshal the response")
	}

	return &response, nil
}
func (pc *ProductionClient) BlockHeaderLookup(blockID xdr.Identifier) (*xdr.BlockHeaderLookupResponse, error) {
	request := xdr.BlockHeaderLookupRequest{
		ID: blockID,
	}

	xdrRequest, err := request.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "could not marshall to xdr binary")
	}

	binaryResp, err := makeRequest(pc.httpClient, pc.server+"/block/header/lookup", xdrRequest)
	if err != nil {
		return nil, errors.Wrap(err, "could not call the endpoint")
	}

	response := xdr.BlockHeaderLookupResponse{}

	err = response.UnmarshalBinary(binaryResp)
	if err != nil {
		return nil, errors.Wrap(err, "could not unmarshal the response")
	}

	return &response, nil
}
func (pc *ProductionClient) AccountInfoLookup(accountID xdr.ID) (*xdr.AccountInfoLookupResponse, error) {
	return nil, nil
}
func (pc *ProductionClient) NonceLookup(accountID xdr.ID) (*xdr.AccountNonceLookupResponse, error) {
	return nil, nil
}

func (pc *ProductionClient) ChannelInfoLookup(channelInfoType xdr.ChannelInfoType) (*xdr.ChannelInfoLookupResponse, error) {
	return nil, nil
}

func makeRequest(httpClient http.Client, url string, xdrRequest []byte) ([]byte, error) {
	b64request := base64.StdEncoding.WithPadding(base64.StdPadding).EncodeToString(xdrRequest)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, url, strings.NewReader(b64request))
	if err != nil {
		return nil, errors.Wrap(err, "could not create a new request")
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "could not call the endpoint")
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
		return nil, errors.Wrap(err, "could not decode base64 body")
	}

	return binaryResp, nil
}

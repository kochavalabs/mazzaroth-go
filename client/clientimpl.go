

var _ MazzarothClient = &MazzarothClient{}

// MazzarothClient is the actual client implementation.
type MazzarothClient struct {
	serverSelector ServerSelector
	httpClient     *http.Client
}

// NewMazzarothClient creates a production object.
func NewMazzarothClient(opts *Options) (*MazzarothClient, error) {
	serverSelector, err := NewRoundRobinServerSelector(opts.servers)
	if err != nil {
		return nil, errors.Wrap(err, "could not create round robin server selector")
	}

	return &MazzarothClient{
		httpClient:     opts.httpClient,
		serverSelector: serverSelector,
	}, nil
}

// TransactionSubmit calls the endpoint: /transaction/submit.
func (pc *MazzarothClient) TransactionSubmit(transaction xdr.Transaction) (*xdr.TransactionSubmitResponse, error) {
	transactionRequest := xdr.TransactionSubmitRequest{
		Transaction: transaction,
	}

	xdrRequest, err := transactionRequest.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "could not marshal to xdr binary")
	}

	binaryResp, err := makeRequest(pc.httpClient, pc.serverSelector.Pick()+"/transaction/submit", xdrRequest)
	if err != nil {
		return nil, errors.Wrap(err, "could not call the endpoint")
	}

	transactionResponse := xdr.TransactionSubmitResponse{}

	err = transactionResponse.UnmarshalBinary(binaryResp)

	return &transactionResponse, errors.Wrap(err, "could not unmarshal the response")
}

// ReadOnly calls the endpoint: /readonly.
func (pc *MazzarothClient) ReadOnly(function string, parameters ...xdr.Parameter) (*xdr.ReadonlyResponse, error) {
	request := xdr.ReadonlyRequest{
		Call: xdr.Call{
			Function:   function,
			Parameters: parameters,
		},
	}

	xdrRequest, err := request.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "could not marshal to xdr binary")
	}

	binaryResp, err := makeRequest(pc.httpClient, pc.serverSelector.Pick()+"/readonly", xdrRequest)
	if err != nil {
		return nil, errors.Wrap(err, "could not call the endpoint")
	}

	response := xdr.ReadonlyResponse{}

	err = response.UnmarshalBinary(binaryResp)

	return &response, errors.Wrap(err, "could not unmarshal the response")
}

// TransactionLookup calls the endpoint: /transaction/lookup.
func (pc *MazzarothClient) TransactionLookup(transactionID xdr.ID) (*xdr.TransactionLookupResponse, error) {
	request := xdr.TransactionLookupRequest{
		TransactionID: transactionID,
	}

	xdrRequest, err := request.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "could not marshal to xdr binary")
	}

	binaryResp, err := makeRequest(pc.httpClient, pc.serverSelector.Pick()+"/transaction/lookup", xdrRequest)
	if err != nil {
		return nil, errors.Wrap(err, "could not call the endpoint")
	}

	response := xdr.TransactionLookupResponse{}

	err = response.UnmarshalBinary(binaryResp)

	return &response, errors.Wrap(err, "could not unmarshal the response")
}

// ReceiptLookup calls the endpoint: /receipt/lookup.
func (pc *MazzarothClient) ReceiptLookup(transactionID xdr.ID) (*xdr.ReceiptLookupResponse, error) {
	request := xdr.ReceiptLookupRequest{
		TransactionID: transactionID,
	}

	xdrRequest, err := request.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "could not marshal to xdr binary")
	}

	binaryResp, err := makeRequest(pc.httpClient, pc.serverSelector.Pick()+"/receipt/lookup", xdrRequest)
	if err != nil {
		return nil, errors.Wrap(err, "could not call the endpoint")
	}

	response := xdr.ReceiptLookupResponse{}

	err = response.UnmarshalBinary(binaryResp)

	return &response, errors.Wrap(err, "could not unmarshal the response")
}

// BlockLookup calls the endpoint: /block/lookup.
func (pc *MazzarothClient) BlockLookup(blockID xdr.Identifier) (*xdr.BlockLookupResponse, error) {
	request := xdr.BlockLookupRequest{
		ID: blockID,
	}

	xdrRequest, err := request.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "could not marshal to xdr binary")
	}

	binaryResp, err := makeRequest(pc.httpClient, pc.serverSelector.Pick()+"/block/lookup", xdrRequest)
	if err != nil {
		return nil, errors.Wrap(err, "could not call the endpoint")
	}

	response := xdr.BlockLookupResponse{}

	err = response.UnmarshalBinary(binaryResp)

	return &response, errors.Wrap(err, "could not unmarshal the response")
}

// BlockHeaderLookup calls the endpoint: /block/header/lookup.
func (pc *MazzarothClient) BlockHeaderLookup(blockID xdr.Identifier) (*xdr.BlockHeaderLookupResponse, error) {
	request := xdr.BlockHeaderLookupRequest{
		ID: blockID,
	}

	xdrRequest, err := request.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "could not marshal to xdr binary")
	}

	binaryResp, err := makeRequest(pc.httpClient, pc.serverSelector.Pick()+"/block/header/lookup", xdrRequest)
	if err != nil {
		return nil, errors.Wrap(err, "could not call the endpoint")
	}

	response := xdr.BlockHeaderLookupResponse{}

	err = response.UnmarshalBinary(binaryResp)

	return &response, errors.Wrap(err, "could not unmarshal the response")
}

// AccountInfoLookup calls the endpoint: /account/info/lookup.
func (pc *MazzarothClient) AccountInfoLookup(accountID xdr.ID) (*xdr.AccountInfoLookupResponse, error) {
	request := xdr.AccountInfoLookupRequest{
		Account: accountID,
	}

	xdrRequest, err := request.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "could not marshal to xdr binary")
	}

	binaryResp, err := makeRequest(pc.httpClient, pc.serverSelector.Pick()+"/account/info/lookup", xdrRequest)
	if err != nil {
		return nil, errors.Wrap(err, "could not call the endpoint")
	}

	response := xdr.AccountInfoLookupResponse{}

	err = response.UnmarshalBinary(binaryResp)

	return &response, errors.Wrap(err, "could not unmarshal the response")
}

// NonceLookup calls the endpoint: /account/nonce/lookup.
func (pc *MazzarothClient) NonceLookup(accountID xdr.ID) (*xdr.AccountNonceLookupResponse, error) {
	request := xdr.AccountNonceLookupRequest{
		Account: accountID,
	}

	xdrRequest, err := request.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "could not marshal to xdr binary")
	}

	binaryResp, err := makeRequest(pc.httpClient, pc.serverSelector.Pick()+"/account/nonce/lookup", xdrRequest)
	if err != nil {
		return nil, errors.Wrap(err, "could not call the endpoint")
	}

	response := xdr.AccountNonceLookupResponse{}

	err = response.UnmarshalBinary(binaryResp)

	return &response, errors.Wrap(err, "could not unmarshal the response")
}

// ChannelInfoLookup calls the endpoint: /channel/info/lookup.
func (pc *MazzarothClient) ChannelInfoLookup(channelInfoType xdr.ChannelInfoType) (*xdr.ChannelInfoLookupResponse, error) {
	request := xdr.ChannelInfoLookupRequest{
		InfoType: channelInfoType,
	}

	xdrRequest, err := request.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "could not marshal to xdr binary")
	}

	binaryResp, err := makeRequest(pc.httpClient, pc.serverSelector.Pick()+"/channel/info/lookup", xdrRequest)
	if err != nil {
		return nil, errors.Wrap(err, "could not call the endpoint")
	}

	response := xdr.ChannelInfoLookupResponse{}

	err = response.UnmarshalBinary(binaryResp)

	return &response, errors.Wrap(err, "could not unmarshal the response")
}

func makeRequest(httpClient *http.Client, url string, xdrRequest []byte) ([]byte, error) {
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

	return binaryResp, errors.Wrap(err, "could not unmarshal the response")
}

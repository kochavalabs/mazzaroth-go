package mazzaroth

// func TestUpdatePermissionBuilder(t *testing.T) {
// 	testAddress, _ := xdr.IDFromSlice([]byte("00000000000000000000000000000000"))
// 	testChannel, _ := xdr.IDFromSlice([]byte("00000000000000000000000000000000"))
// 	publicKey := "0000000000000000000000000000000000000000000000000000000000000000"
// 	seed, _ := hex.DecodeString(publicKey)
// 	privateKey := ed25519.NewKeyFromSeed(seed)
// 	action := xdr.Action{
// 		Address:               testAddress,
// 		ChannelID:             testChannel,
// 		Nonce:                 0,
// 		BlockExpirationNumber: 1,
// 		Category: xdr.ActionCategory{
// 			Type: xdr.ActionCategoryTypeUPDATE,
// 			Update: &xdr.Update{
// 				Type: xdr.UpdateTypePERMISSION,
// 				Permission: &xdr.Permission{
// 					Action: 1,
// 					Key:    testAddress,
// 				},
// 			},
// 		},
// 	}
// 	actionStream, err := action.MarshalBinary()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	signatureSlice := ed25519.Sign(privateKey, actionStream)
// 	signature, err := xdr.SignatureFromSlice(signatureSlice)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	wantTx := &xdr.Transaction{
// 		Signature: signature,
// 		Action:    action,
// 	}
// 	ub := new(UpdatePermissionBuilder)
// 	tx, err := ub.UpdatePermission(&testAddress, &testChannel, 0, 1).
// 		Action(1).
// 		Address(testAddress).Sign(privateKey)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if !reflect.DeepEqual(wantTx, tx) {
// 		t.Fatalf("expected: %v, got: %v", wantTx, tx)
// 	}
// }

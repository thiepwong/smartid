package wallet

import (
	hex "github.com/ethereum/go-ethereum/common/hexutil"

	crypto "github.com/ethereum/go-ethereum/crypto"
)

//Wallet...
type Wallet struct {
	Address    string
	PublicKey  string
	PrivateKey string
}

//CreateWallet
func CreateWallet() (wallet Wallet, err error) {

	key, err := crypto.GenerateKey()
	//	wallet.publicKey = key.PublicKey
	wallet.Address = crypto.PubkeyToAddress(key.PublicKey).Hex()
	wallet.PublicKey = hex.Encode(crypto.FromECDSAPub(&key.PublicKey))
	wallet.PrivateKey = hex.Encode(key.D.Bytes())
	return wallet, err
}

func WalletValidate(w Wallet) bool {
	return true
}

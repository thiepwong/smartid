package wallet

import (
	"github.com/ethereum/go-ethereum/common/hexutil"

	crypto "github.com/ethereum/go-ethereum/crypto"
)

type Wallet struct {
	Address    string
	PublicKey  string
	PrivateKey string
}

func CreateWallet() (wallet Wallet, err error) {

	key, err := crypto.GenerateKey()
	//	wallet.publicKey = key.PublicKey
	wallet.Address = crypto.PubkeyToAddress(key.PublicKey).Hex()
	wallet.PublicKey = hexutil.Encode(crypto.FromECDSAPub(&key.PublicKey))
	wallet.PrivateKey = hexutil.Encode(key.D.Bytes())
	return wallet, err
}

func WalletValidate(w Wallet) bool {
	return true
}

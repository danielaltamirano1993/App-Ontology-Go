package bench

import (
	"fmt"
	"strconv"

	"github.com/ontio/ontology-crypto/keypair"
	ont "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/cmd/utils"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/common/log"
	"github.com/ontio/ontology/core/types"
)

func MultiSigTransfer() {
	m := 5
	var wallets []string
	wallets = []string{
		"./bench/wallet1.dat",
		"./bench/wallet2.dat",
		"./bench/wallet3.dat",
		"./bench/wallet4.dat",
		"./bench/wallet5.dat",
		"./bench/wallet6.dat",
		"./bench/wallet7.dat",
	}
	toAddr := "AYMnqA65pJFKAbbpD8hi5gdNDBmeFBy5hS"

	accs := make([]*ont.Account, 0)
	pks := make([]keypair.PublicKey, 0, len(wallets))
	for _, w := range wallets {
		a, err := loadWallet(w, "pwd")
		if err != nil {
			log.Errorf("load wallet error: %s", err)
			return
		}
		accs = append(accs, a)
		pks = append(pks, a.PublicKey)
		log.Infof("addr:%s", a.Address.ToBase58())
	}
	payer, err := types.AddressFromMultiPubKeys(pks, int(m))
	if err != nil {
		log.Errorf("AddressFromMultiPubKeyserr:%s", err)
		return
	}
	balance, err := utils.GetBalance(payer.ToBase58())
	log.Infof("payer:%s ont:%s", payer.ToBase58(), balance.Ont)
	payerOnt, err := strconv.Atoi(balance.Ont)

	toBalance, err := utils.GetBalance(toAddr)
	if err != nil {
		log.Errorf("addr from base 58 err:%s", err)
	}
	log.Infof("to:%s ont:%s, payerOnt:%d", toAddr, toBalance.Ont, payerOnt)

	txhash, err := sdkMultiTransfer(0, 30000, accs, uint16(m), "ont", payer, toAddr, 100000000)
	if err != nil {
		log.Errorf("err:%s", err)
		return
	}
	log.Infof("hash: %x", txhash)

}

func sdkMultiTransfer(gasPrice, gasLimit uint64, signers []*ont.Account, m uint16, asset string, from common.Address, to string, amount uint64) (string, error) {
	pks := make([]keypair.PublicKey, 0, len(signers))
	for _, signer := range signers {
		pks = append(pks, signer.PublicKey)
	}

	sdk := ont.NewOntologySdk()

	sdk.NewRpcClient().SetAddress("http://localhost:20336")

	mutableTx, err := utils.TransferTx(0, gasLimit, asset, from.ToBase58(), to, amount)
	if err != nil {
		return "", err
	}

	for _, signer := range signers {
		sdk.MultiSignToTransaction(mutableTx, m, pks, signer)
		fmt.Printf("sign\n")
	}
	hash, err := sdk.SendTransaction(mutableTx)
	if err != nil {
		return "", err
	}
	return hash.ToHexString(), nil
}

func loadWallet(wallet, pwd string) (*ont.Account, error) {
	clientImpl, err := ont.OpenWallet(wallet)
	if clientImpl == nil {
		log.Errorf("clientImpl is nil")
		return nil, err
	}
	a, err := clientImpl.GetDefaultAccount([]byte(pwd))
	if a == nil {
		log.Errorf("acc is nil")
		return nil, err
	}
	return a, nil
}

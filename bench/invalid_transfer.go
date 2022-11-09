package bench

// import (
// 	"encoding/hex"
// 	"time"

// 	"github.com/ontio/ontology/core/payload"

// 	"github.com/ontio/ontology-crypto/keypair"
// 	"github.com/ontio/ontology/account"
// 	"github.com/ontio/ontology/cmd/utils"
// 	"github.com/ontio/ontology/common"
// 	"github.com/ontio/ontology/common/log"
// 	"github.com/ontio/ontology/core/types"
// )

// type InvalidTxType int

// const (
// 	InvalidTxUseOtherAcc InvalidTxType = 0
// 	InvalidTxDunplicated               = 1
// 	InvalidTxResend                    = 2
// 	InvalidDoubleSpend                 = 3
// )

// func (this *TestTransfer) InvokeInvalidTransaction(txType InvalidTxType) {
// 	switch txType {
// 	case InvalidTxUseOtherAcc:
// 		this.UseOtherAccTransaction()
// 	case InvalidTxDunplicated:
// 		this.DuplicateTransaction()
// 	case InvalidTxResend:
// 		this.ResendUsedTransaction()
// 	case InvalidDoubleSpend:
// 		this.DoubleSpendTransaction()
// 	}
// }

// func (this *TestTransfer) UseOtherAccTransaction() {
// 	accs := make([]*account.Account, 0)
// 	pks := make([]keypair.PublicKey, 0, 1)
// 	a, err := loadWallet("./bench/invalidtx_wallets/node1_wallet.dat", "pwd")
// 	if err != nil {
// 		log.Errorf("load wallet error: %s", err)
// 		return
// 	}
// 	accs = append(accs, a)
// 	pks = append(pks, a.PublicKey)
// 	log.Infof("addr:%s", a.Address.ToBase58())
// 	otherAddr := "AZJ9Wkpn99gEYZiyG339ueK9s6Ms768hDX"
// 	balance, err := utils.GetBalance(a.Address.ToBase58())
// 	otherUserBal, err := utils.GetBalance(otherAddr)
// 	log.Infof("self balance:%s, other user balance:%s", balance.Ont, otherUserBal.Ont)
// 	log.Infof("otherAddr:%s", otherAddr)
// 	txHash, err := this.transfer(0, 30000, accs, 1, "ont", accs[0].Address.ToBase58(), a.Address.ToBase58(), 10000)
// 	if err != nil {
// 		log.Errorf("tx error:%s", err)
// 	}
// 	log.Infof("tx hash: %s", txHash)
// }

// func (this *TestTransfer) DuplicateTransaction() {
// 	accs := make([]*account.Account, 0)
// 	a, err := loadWallet("./bench/invalidtx_wallets/test_wallet.dat", "pwd")
// 	if err != nil {
// 		log.Errorf("load wallet error: %s", err)
// 		return
// 	}
// 	accs = append(accs, a)
// 	toAddr := "AbB2GVCfQbXVV2bR6wLsTXCAK8ThDEmzBE"
// 	balance, err := utils.GetBalance(a.Address.ToBase58())

// 	log.Infof("addr:%x", a.Address)
// 	toBal, err := utils.GetBalance(toAddr)
// 	log.Infof("self balance:%s, other user balance:%s", balance.Ont, toBal.Ont)
// 	log.Infof("toAddr:%s", toAddr)
// 	for i := 0; i < 10; i++ {
// 		txHash, err := this.transfer(0, 30000, accs, 1, "ont", accs[0].Address.ToBase58(), toAddr, 1)
// 		if err != nil {
// 			log.Errorf("tx error:%s", err)
// 		}
// 		log.Infof("tx hash: %s", txHash)
// 	}

// 	transferTx, err := utils.TransferTx(0, 30000, "ont", a.Address.ToBase58(), toAddr, 1)
// 	if err != nil {
// 		log.Errorf("tx error:%s", err)
// 		return
// 	}

// 	err = utils.SignTransaction(a, transferTx)
// 	if err != nil {
// 		log.Errorf("tx error:%s", err)
// 		return
// 	}

// 	transferTx.Nonce = 1528875844
// 	txHash, err := utils.SendRawTransaction(transferTx)
// 	if err != nil {
// 		log.Errorf("tx error:%s", err)
// 		return
// 	}
// 	log.Infof("tx hash: %s", txHash)
// }

// func (this *TestTransfer) ResendUsedTransaction() {
// 	rawTx := &types.Transaction{}
// 	rawTx.Version = byte(0)
// 	rawTx.Nonce = uint32(time.Now().Unix())
// 	rawTx.GasPrice = 0
// 	rawTx.GasLimit = 30000
// 	rawTx.TxType = 0xd1

// 	payer, err := common.AddressFromBase58("AdqXGzzNqt9cMkUiU5Miq5pupVFF92XfKu")
// 	rawTx.Payer = payer
// 	log.Infof("payer:%x", payer)

// 	payload := &payload.InvokeCode{}
// 	b, err := hex.DecodeString("00c66b14f2036d30f1206549627d00c50484a5d5a64c0cac6a7cc814d4cb272d9e4a343ed878be19f6a5519f309b8b2e6a7cc8516a7cc86c51c1087472616e736665721400000000000000000000000000000000000000010068164f6e746f6c6f67792e4e61746976652e496e766f6b65")
// 	payload.Code = b
// 	log.Infof("payload:%x", payload.Code)

// 	rawTx.Payload = payload

// 	s, err := hex.DecodeString("01f3e1fdeb3db9d82a77e6b3db5ac2fc746917685a3e3da99e4f00787b035275352baa1019d5698b96db4e019c1b0adca67740020a6bd9835512c15f2febbd0bba")
// 	p, err := hex.DecodeString("120203db99606cf5906d57d4c18dba772832e36837328b84a3efe804b38b85d9cb5f8b")
// 	pk, err := keypair.DeserializePublicKey(p)
// 	sig := &types.Sig{}
// 	sig.M = 1
// 	sig.SigData = [][]byte{s}
// 	sig.PubKeys = []keypair.PublicKey{pk}
// 	log.Infof("pk:%x", p)
// 	log.Infof("SigData:%x", s)

// 	rawTx.Sigs = []*types.Sig{sig}
// 	log.Infof("nonce:%d", rawTx.Nonce)

// 	if err != nil {
// 		log.Errorf("make raw tx err:%s", err)
// 	}

// 	txHash, err := utils.SendRawTransaction(rawTx)
// 	if err != nil {
// 		log.Errorf("tx error:%s", err)
// 		return
// 	}
// 	log.Infof("tx hash: %s", txHash)
// }

// func (this *TestTransfer) DoubleSpendTransaction() {
// 	accs := make([]*account.Account, 0)
// 	a, err := loadWallet("./bench/invalidtx_wallets/test_wallet.dat", "pwd")
// 	if err != nil {
// 		log.Errorf("load wallet error: %s", err)
// 		return
// 	}
// 	accs = append(accs, a)
// 	toAddr := "AbB2GVCfQbXVV2bR6wLsTXCAK8ThDEmzBE"
// 	balance, err := utils.GetBalance(a.Address.ToBase58())

// 	log.Infof("addr:%x", a.Address)
// 	toBal, err := utils.GetBalance(toAddr)
// 	log.Infof("self balance:%s, other user balance:%s", balance.Ont, toBal.Ont)
// 	log.Infof("toAddr:%s", toAddr)

// 	transferTx, err := utils.TransferTx(0, 30000, "ont", a.Address.ToBase58(), toAddr, 1)
// 	if err != nil {
// 		log.Errorf("tx error:%s", err)
// 		return
// 	}

// 	err = utils.SignTransaction(a, transferTx)
// 	if err != nil {
// 		log.Errorf("tx error:%s", err)
// 		return
// 	}

// 	//same signature, different nonce
// 	now := time.Now().Unix()
// 	for i := 0; i < 2; i++ {
// 		transferTx.Nonce = uint32(now)
// 		now++
// 		log.Infof("transferTx.Nonce:%d, sig:%x", transferTx.Nonce, transferTx.Sigs[0].SigData[0])
// 		txHash, err := utils.SendRawTransaction(transferTx)
// 		if err != nil {
// 			log.Errorf("tx error:%s", err)
// 			return
// 		}
// 		log.Infof("tx hash: %s", txHash)
// 	}

// 	//different signature, same nonce
// 	for i := 0; i < 2; i++ {
// 		transferTx, err := utils.TransferTx(0, 30000+uint64(i), "ont", a.Address.ToBase58(), toAddr, 9980)
// 		if err != nil {
// 			log.Errorf("tx error:%s", err)
// 			return
// 		}

// 		err = utils.SignTransaction(a, transferTx)
// 		if err != nil {
// 			log.Errorf("tx error:%s", err)
// 			return
// 		}
// 		log.Infof("transferTx.Nonce:%d, sig:%x", transferTx.Nonce, transferTx.Sigs[0].SigData[0])
// 		txHash, err := utils.SendRawTransaction(transferTx)
// 		if err != nil {
// 			log.Errorf("tx error:%s", err)
// 			return
// 		}
// 		log.Infof("tx hash: %s", txHash)
// 	}
// }

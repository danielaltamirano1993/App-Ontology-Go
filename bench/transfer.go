package bench

// import (
// 	"io/ioutil"
// 	"strings"
// 	"sync"
// 	"time"

// 	ontSdk "github.com/ontio/ontology-go-sdk"
// 	ontSdkCom "github.com/ontio/ontology-go-sdk/common"

// 	"github.com/ontio/ontology/account"
// 	"github.com/ontio/ontology/common"
// 	"github.com/ontio/ontology/common/log"
// )

// const (
// 	DEF_WALLET_PWD    = "pwd" //default wallet password
// 	NODES_ADDRS_FILE  = "./addrs"
// 	PAYER_WALLET_FILE = "./bench/wallet.dat"
// )

// const (
// 	TRANSFER_ONT_DURATION     = 1   // transfer ont duration in second
// 	NO_ENOUGH_FOUND_MAX_CHECK = 600 // max check 0 balance times, if reach, stop the timer
// )

type TestTransfer struct {
	// 	defAccount     *account.Account
	// 	sdk            *ontSdk.OntologySdk
	// 	toAddrs        []common.Address
	// 	repeat         int // repeat times for transfer to a address
	// 	noEnoughFound  int
	// 	stopTimerCh    chan bool
	// 	lock           *sync.Mutex
	// 	accountBalance *ontSdkCom.Balance
	// 	tps            int // transaction per second
	// 	amount         uint64
	// 	rpcAddr        string
}

// func NewTestTransfer() *TestTransfer {
// 	return &TestTransfer{}
// }

// func (this *TestTransfer) Start() {
// 	log.Infof("ont_tps:%d, amount:%d, rpc address:%s\n", this.tps, this.amount, this.rpcAddr)
// 	ret := this.initVars()
// 	if !ret {
// 		log.Error("init instance variable failed")
// 		return
// 	}
// 	timer := time.NewTicker(time.Duration(TRANSFER_ONT_DURATION * time.Second))
// 	for {
// 		select {
// 		case <-timer.C:
// 			this.transferOnt()
// 		case <-this.stopTimerCh:
// 			log.Info("stop timer because no enough found")
// 			timer.Stop()
// 			goto FINISHED
// 		}
// 	}
// FINISHED:
// 	log.Info("finished")
// }

// func (this *TestTransfer) SetTps(tps int) {
// 	this.tps = tps
// }

// func (this *TestTransfer) SetAmount(amount uint64) {
// 	this.amount = amount
// }

// func (this *TestTransfer) SetRpc(rpc string) {
// 	this.rpcAddr = rpc
// }

// func (this *TestTransfer) InitSdk() {
// 	this.sdk = ontSdk.NewOntologySdk()
// 	if len(this.rpcAddr) > 0 {
// 		this.sdk.Rpc.SetAddress(this.rpcAddr)
// 	} else {
// 		this.sdk.Rpc.SetAddress("http://localhost:20336")
// 	}
// }

// func (this *TestTransfer) initVars() bool {
// 	this.lock = &sync.Mutex{}
// 	this.toAddrs = getToAddrs()
// 	if this.toAddrs == nil || len(this.toAddrs) == 0 {
// 		log.Warnf("no transfer to address")
// 		return false
// 	}
// 	this.repeat = (int)(this.tps / len(this.toAddrs))
// 	log.Infof("Transfer address count:%d, each address repeat %d", len(this.toAddrs), this.repeat)
// 	this.noEnoughFound = 0
// 	this.stopTimerCh = make(chan bool, 1)

// 	this.InitSdk()

// 	clientImpl, err := account.NewClientImpl(PAYER_WALLET_FILE)
// 	if err != nil {
// 		log.Errorf("import wallet failed")
// 		return false
// 	}

// 	this.defAccount, err = clientImpl.GetDefaultAccount([]byte(DEF_WALLET_PWD))
// 	if err != nil {
// 		log.Errorf("client get default account failed")
// 		return false
// 	}
// 	this.accountBalance, err = this.sdk.Rpc.GetBalance(this.defAccount.Address)
// 	if this.defAccount == nil {
// 		log.Warnf("defaccount is nil")
// 		return false
// 	}
// 	if this.accountBalance == nil {
// 		log.Warn("balance is nil")
// 	}
// 	// log.Infof("default account address:%v, balance:%d", this.defAccount.Address.ToBase58(), this.accountBalance.Ont)
// 	if err != nil {
// 		log.Errorf("get balance failed, error:%s", err)
// 		return false
// 	}

// 	return true
// }

// func (this *TestTransfer) transferOnt() {
// 	if !this.isBalanceEnough() {
// 		return
// 	}
// 	for i := 0; i < this.repeat; i++ {
// 		for _, toAddr := range this.toAddrs {
// 			if !this.isBalanceEnough() {
// 				return
// 			}
// 			gasLimit := 30000 + i
// 			amount := uint64(i)
// 			txHash, err := this.sdk.Rpc.Transfer(0, uint64(gasLimit), "ONT", this.defAccount, toAddr, amount)
// 			if err != nil {
// 				log.Errorf("transfer error:%s, txHash:%x", err, txHash)
// 				continue
// 			}
// 			log.Infof("transfer:%d from:%s to:%s", amount, this.defAccount.Address.ToBase58(), toAddr.ToBase58())
// 		}
// 	}
// }

// func (this *TestTransfer) isBalanceEnough() bool {
// 	this.lock.Lock()
// 	defer this.lock.Unlock()
// 	return true
// 	// if this.accountBalance.Ont == 0 || this.accountBalance.Ont < this.amount {
// 	// 	log.Warnf("no enough ont, balance:%d", this.accountBalance.Ont)
// 	// 	this.noEnoughFound++
// 	// 	if this.noEnoughFound > NO_ENOUGH_FOUND_MAX_CHECK {
// 	// 		this.stopTimerCh <- true
// 	// 	}
// 	// 	return false
// 	// } else {
// 	// 	return true
// 	// }
// }

// // read address from file
// func getToAddrs() []common.Address {
// 	file, err := ioutil.ReadFile(NODES_ADDRS_FILE)
// 	if err != nil {
// 		return nil
// 	}
// 	addrs := strings.Split(string(file), "\n")
// 	var addresses []common.Address
// 	for _, addr := range addrs {
// 		if len(addr) > 0 {
// 			toAddr, err := common.AddressFromBase58(addr)
// 			if err != nil {
// 				continue
// 			}
// 			addresses = append(addresses, toAddr)
// 		}
// 	}
// 	return addresses
// }

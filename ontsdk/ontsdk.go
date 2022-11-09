package ontsdk

import (
	"encoding/hex"
	"fmt"

	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/smartcontract/service/native/dns"
	"github.com/ontio/ontology/smartcontract/service/native/utils"
)

func DnsInit() {
	oSDK := sdk.NewOntologySdk()
	oSDK.NewRpcClient().SetAddress("http://localhost:20336")

	var err error
	testWallet, err := oSDK.OpenWallet("../testdata/wallet1.dat")
	if err != nil {
		fmt.Printf("account.Open error:%s\n", err)
		return
	}
	testDefAcc, err := testWallet.GetDefaultAccount([]byte("pwd"))
	if err != nil {
		fmt.Printf("GetDefaultAccount error:%s\n", err)
		return
	}
	ret, err := oSDK.Native.InvokeNativeContract(0, 30000, testDefAcc, byte(0), utils.OntDNSAddress, dns.INIT_NAME, nil)
	if err != nil {
		fmt.Printf("InvokeNativeContract error:%s\n", err)
		return
	}
	fmt.Printf("tx:%s\n", hex.EncodeToString(common.ToArrayReverse(ret[:])))
}

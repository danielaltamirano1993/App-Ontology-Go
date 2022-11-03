package main

// import (
// 	"github.com/ontio/ontology-test/bench"
// )

// func main() {
// 	test := bench.TestTransfer{}
// 	test.MultiSigTransfer()
// 	// log.InitLog(0, log.PATH, log.Stdout)
// 	// runApp()
// }

// // func testSameReq() {
// // 	conn, err := net.Dial("tcp", "127.0.0.1:20338")
// // 	if err != nil {
// // 		fmt.Printf("dial err:%s\n", err)
// // 	}
// // 	hash, err := common.Uint256FromHexString("1")
// // 	if err != nil {
// // 		fmt.Printf("parse err:%s\n", err)
// // 	}
// // 	msg := msgpack.NewBlkDataReq(hash)
// // 	buf := bytes.NewBuffer(nil)
// // 	err = types.WriteMessage(buf, msg)
// // 	if err != nil {
// // 		fmt.Printf("types write msg err:%s\n", err)
// // 	}

// // 	payload := buf.Bytes()
// // 	_, err = conn.Write(payload)
// // 	if err != nil {
// // 		fmt.Printf("conn write msg err:%s\n", err)
// // 	}
// // 	for {

// // 	}
// // }

// // func CheckHash(file string) {
// // 	con, err := ioutil.ReadFile(fmt.Sprintf("./Log/%s", file))
// // 	ret := strings.Split(string(con), "\n")
// // 	if err != nil {

// // 	}
// // 	var hashes []string
// // 	for _, line := range ret {
// // 		if strings.Index(line, "hash") != -1 {
// // 			hash := strings.Split(line, "hash: ")
// // 			hashes = append(hashes, hash[1])
// // 		}
// // 	}
// // 	for i, hash := range hashes {
// // 		for j, h := range hashes {
// // 			if hash == h && i != j {
// // 				fmt.Println(hash)
// // 			}
// // 		}
// // 	}
// // 	fmt.Printf("done:%d\n", len(hashes))
// // }

// // func runApp() {
// // 	c := cmd.NewCmd()
// // 	c.Run()
// // 	switch c.GetAction() {
// // 	case cmd.CmdActionBatchTransfer:
// // 		t := bench.NewTestTransfer()
// // 		t.SetTps(c.GetOntTPS())
// // 		t.SetAmount(c.GetAmount())
// // 		t.SetRpc(c.GetRpc())
// // 		t.Start()
// // 		// case cmd.CmdActionMutilTransfer:
// // 		// 	t := bench.NewTestTransfer()
// // 		// 	t.SetTps(c.GetOntTPS())
// // 		// 	t.SetAmount(c.GetAmount())
// // 		// 	t.SetRpc(c.GetRpc())
// // 		// 	t.InitSdk()
// // 		// 	t.MultiSigTransfer()
// // 		// case cmd.CmdActionBatchAnalysis:
// // 		// 	txn := analysis.SumUpTxs(c.GetAnalysisPath())
// // 		// 	log.Infof("tx cnt:		%d", txn)
// // 		// case cmd.CmdActionInvalidTransfer:
// // 		// 	ty := c.GetInvalidTxType()
// // 		// 	t := bench.NewTestTransfer()
// // 		// 	t.SetRpc(c.GetRpc())
// // 		// 	t.InvokeInvalidTransaction(bench.InvalidTxType(ty))
// // 		// case cmd.CmdActionSignatureService:
// // 		// 	t := bench.NewTestTransfer()
// // 		// 	t.SetRpc(c.GetRpc())
// // 		// 	t.SignatureService()
// // 	}
// // }

// // package main

// // import (
// // 	"fmt"
// // 	"net"

// // 	"github.com/ontio/ontology/common/log"
// // 	"github.com/ontio/ontology/p2pserver/common"
// // 	"github.com/ontio/ontology/p2pserver/message/types"
// // )

// // func main() {
// // 	log.Init(log.Stdout)

// // 	//conn, err := net.Dial("tcp", "127.0.0.1:21338")
// // 	conn, err := net.Dial("tcp", "127.0.0.1:41338")
// // 	if err != nil {
// // 		fmt.Println("cont err:%s", err)
// // 		return
// // 	}

// // 	for i := 0; i < 12000; i++ {
// // 		msg, err := types.MakeEmptyMessage(common.GET_HEADERS_TYPE)
// // 		if err != nil {
// // 			fmt.Println("cont err:%s", err)
// // 			return
// // 		}
// // 		err = types.WriteMessage(conn, msg)
// // 		if err != nil {
// // 			fmt.Println("cont err:%s", err)
// // 			return
// // 		}
// // 	}
// // 	for i := 0; i < 12000; i++ {
// // 		msg, err := types.MakeEmptyMessage(common.HEADERS_TYPE)
// // 		if err != nil {
// // 			fmt.Println("cont err:%s", err)
// // 			return
// // 		}
// // 		err = types.WriteMessage(conn, msg)
// // 		if err != nil {
// // 			fmt.Println("cont err:%s", err)
// // 			return
// // 		}
// // 	}

// // }

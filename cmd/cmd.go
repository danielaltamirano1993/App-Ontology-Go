package cmd

import (
	"os"

	"log"

	"github.com/urfave/cli"
)

type CmdAction int

type Cmd struct {
	ctx    *cli.Context
	action CmdAction
}

const (
	CmdActionBatchTransfer    CmdAction = 0
	CmdActionBatchAnalysis              = 1
	CmdActionMutilTransfer              = 2
	CmdActionInvalidTransfer            = 3
	CmdActionSignatureService           = 4
)

func NewCmd() *Cmd {
	return &Cmd{}
}

func (this *Cmd) Run() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "tps",
			Value: 3000,
			Usage: "transaction per second",
		},
		cli.IntFlag{
			Name:  "amount",
			Value: 1,
			Usage: "transfer amount",
		},
		cli.StringFlag{
			Name:  "rpc",
			Value: "http://localhost:20336",
			Usage: "local rpc server port",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "transaction",
			Aliases: []string{"tx"},
			Usage:   "calculate tx total count",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "path",
					Value: "",
					Usage: "log file to calc",
				},
			},
			Action: func(c *cli.Context) error {
				this.ctx = c
				this.action = CmdActionBatchAnalysis
				return nil
			},
		},
		{
			Name:    "multisigtx",
			Aliases: []string{"mtx"},
			Usage:   "multi signature transfer",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "rpc",
					Value: "http://localhost:20336",
					Usage: "local rpc server port",
				},
			},
			Action: func(c *cli.Context) error {
				this.ctx = c
				this.action = CmdActionMutilTransfer
				return nil
			},
		},
		{
			Name:    "invalidtx",
			Aliases: []string{"intx"},
			Usage:   "invalid transfer",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "rpc",
					Value: "http://localhost:20336",
					Usage: "local rpc server port",
				},
				cli.IntFlag{
					Name:  "type",
					Value: 0,
					Usage: "invalid tx type",
				},
			},
			Action: func(c *cli.Context) error {
				this.ctx = c
				this.action = CmdActionInvalidTransfer
				return nil
			},
		},
		{
			Name:    "signature",
			Aliases: []string{"sign"},
			Usage:   "sign data",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "rpc",
					Value: "http://localhost:20336",
					Usage: "local rpc server port",
				},
			},
			Action: func(c *cli.Context) error {
				this.ctx = c
				this.action = CmdActionSignatureService
				return nil
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		this.ctx = c
		this.action = CmdActionBatchTransfer
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
func (this *Cmd) GetAction() CmdAction {
	return this.action
}
func (this *Cmd) GetOntTPS() int {
	if this.ctx == nil {
		return 1
	}
	tps := this.ctx.Int("tps")
	if tps == 0 {
		return 1
	}
	return tps
}

func (this *Cmd) GetAmount() uint64 {
	if this.ctx == nil {
		return 1
	}
	amount := this.ctx.Uint64("amount")
	if amount == 0 {
		return 1
	}
	return amount
}

func (this *Cmd) GetRpc() string {
	if this.ctx == nil {
		return ""
	}
	rpc := this.ctx.String("rpc")
	return rpc
}

func (this *Cmd) GetAnalysisPath() string {
	if this.ctx == nil {
		return ""
	}
	path := this.ctx.String("path")
	return path
}

func (this *Cmd) GetInvalidTxType() int {
	if this.ctx == nil {
		return 0
	}
	txType := this.ctx.Int("type")
	return txType
}

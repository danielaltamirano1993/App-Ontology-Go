package analysis

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func SumUpTxs(path string) int {
	if len(path) == 0 {
		return 0
	}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return 0
	}
	var f os.FileInfo
	txCnt := 0
	for i := 0; i < len(files); i++ {
		f = files[i]
		if !f.IsDir() && f.Size() > 0 {
			fName := fmt.Sprintf("%s/%s", path, f.Name())
			// fmt.Println(fName)
			fBuf, _ := ioutil.ReadFile(fName)
			c := string(fBuf)
			lines := strings.Split(c, "\n")
			for _, l := range lines {
				txIndex := strings.Index(l, "numtx=")
				if txIndex != -1 {
					cnt, _ := strconv.Atoi(l[txIndex+6:])
					txCnt += cnt
				}
			}
			fmt.Printf("%s :%d\n", fName, txCnt)
		}
	}

	return txCnt
}

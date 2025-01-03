package main

import (
	"fmt"
	"os"

	"github.com/qiuzhanghua/common/hf"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(-1)
	}
	dsName := os.Args[1]

	answer, _ := hf.HfDatasetsPath(dsName)

	fmt.Print(answer)
}

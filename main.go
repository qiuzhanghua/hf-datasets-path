package main

import (
	"fmt"
	"os"

	"github.com/labstack/gommon/log"
	"github.com/qiuzhanghua/common/hf"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(-1)
	}
	dsName := os.Args[1]

	answer, err := hf.HfDatasetsPath(dsName)
	if err != nil {
		log.Errorf("Error: %v", err)
	}

	fmt.Print(answer)
}

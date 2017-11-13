package main

import (
	"fmt"
	"github.com/cihub/seelog"
)

func main() {
	log, err := seelog.LoggerFromConfigAsFile("prod.xml")
	fmt.Println(err)
	defer log.Flush()

	log.Trace("i am Trace.")
	log.Debug("i am Debug.")
	log.Info("i am Info.")
	log.Warn("i am Warn.")
	log.Error("i am Error.")
	log.Critical("i am Critical.")
	log.Info("i am Info2.")

	fmt.Println("< main.")
}

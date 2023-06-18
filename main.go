package main

import (
	"fmt"
	"os"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
)

func main(){
	if len(os.Args) < 3  {
		fmt.Printf("Usage:%s <time> <message>\n", os.Args[0])
		os.Exit(1)
	}

	now := time.Now()
	when := when.New(nil)

	w.Add(en.All...)
	w.Add(common.All...)
}
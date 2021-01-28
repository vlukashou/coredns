package main

//go:generate go run directives_generate.go
//go:generate go run owners_generate.go

import (
	// "time"
	"fmt"

	_ "github.com/coredns/coredns/core/plugin" // Plug in CoreDNS.
	"github.com/coredns/coredns/coremain"
)

func main() {
	// coreDns := coremain.NewCoreDns("corefile")
	coreDns := coremain.NewCoreDns()
	// go func(){
	// 	for {
	// 		fmt.Printf("%s\n",coreDns.GetLog())
	// 		if coreDns.GetLog() == "СoreDns started" { break }
	// 		time.Sleep(1*time.Second)	
	// 	}
	// }()
	coreDns.Init()
	status := coreDns.GetLog() 
	fmt.Printf("%s\n",status)
	coreDns.Run()
}

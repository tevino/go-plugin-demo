package main

import (
	"fmt"
	"log"
	"os"
	"plugin"

	"github.com/tevino/go-plugin-demo/pinger"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [plugin binary]\n", os.Args[0])
		os.Exit(1)
	}
	pluginPath := os.Args[1]

	// load plugin
	p, err := plugin.Open(pluginPath)
	if err != nil {
		log.Fatalf("Error opening plugin[%s]: %s", pluginPath, err)
	}

	// load Name
	nameSym, err := p.Lookup("Name")
	if err != nil {
		log.Fatalf("Name not defined in plugin(%s)", pluginPath)
	}
	fmt.Printf("Plugin.Name: %s\n", *(nameSym).(*string))

	// load P
	pSym, err := p.Lookup("P")
	if err != nil {
		log.Fatalf("P not defined in plugin(%s)", pluginPath)
	}
	var pp pinger.Pinger = *(pSym).(*pinger.Pinger)

	err = pp.Ping()
	if err != nil {
		log.Fatal("Error from Ping: ", err)
	}
}

package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/Jesusssx/rThreatAgentless/pkg/wmiexec"
)

func main() {

	var command, target, username, password, hash, domain, clientHost, binding string
	var resolveOnly bool
	flag.StringVar(&target, "target", "", "Target - Objetivo")
	flag.StringVar(&username, "username", "", "Username to auth as - Usuario a ser autenticado")
	flag.StringVar(&password, "password", "", "password")
	flag.StringVar(&hash, "hash", "", "hash")
	flag.StringVar(&command, "command", "", "command - comando")
	flag.StringVar(&clientHost, "clientname", "", "Value to send the victim indicating client host - Valor para enviar a la víctima indicando el host del cliente")
	flag.StringVar(&binding, "binding", "", "Value to use in network binding (see output of resolve mode for potential values) - Valor para hacer el binding de red")
	flag.BoolVar(&resolveOnly, "resolveonly", false, "Only resolve network bindings (does not require auth) - Solo verifica si el host es alcanzable")
	flag.Parse()

	if clientHost == "" {
		var err error
		clientHost, err = os.Hostname()
		if err != nil {
			panic(err)
		}
	}

	if target == "" || (password == "" && hash == "" && !resolveOnly) {
		flag.Usage()
		os.Exit(1)
	}

	if !strings.Contains(target, ":") {
		//Puerto por cambiar operacion pendiente
		//fmt.Printf("Bad target specified, requires port (usually 135). expected: 127.0.0.1:135, got %s", target)
		//os.Exit(1)
		target = target + ":135"
	}

	//don't do auth, just get network adaptors
	if resolveOnly {
		values, err := wmiexec.GetNetworkBindings(target)
		if err != nil {
			panic(err)
		}
		log.Println("Resolved names:")
		for _, name := range values {
			log.Println("\t", name)
		}
		return
	}

	err := wmiexec.WMIExec(target, username, password, hash, domain, command, clientHost, binding, nil)
	if err != nil {
		panic(err)
	}

}

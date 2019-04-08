package main

import (
	"log"
	"os"
	"plugin"
	"github.com/zimnx/central/plugin"
	"github.com/zimnx/central/ext"
)

func main() {
	plug, err := plugin.Open(os.Args[1])
	if err != nil {
		log.Fatalf("cant open plugin: %s", err)
	}
	sym, err := plug.Lookup("Init")
	if err != nil {
		log.Fatalf("cant open plugin: %s", err)
	}
	f := sym.(func(ext.Map))
	m := myplugin.NewMyMap()
	f(m)
}
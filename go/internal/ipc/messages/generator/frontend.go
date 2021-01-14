package main

import (
	"github.com/diamondburned/cchat-ipc/go/internal/ipc/messages/generator/genutil"
	"github.com/diamondburned/cchat/repository"
)

func (g *Generator) genFrontendServices() {
	for _, iface := range genutil.FrontendInterfaces(g.pkg) {
		g.genFrontendService(iface)
	}
}

func (g *Generator) genFrontendService(iface repository.Interface) {

}

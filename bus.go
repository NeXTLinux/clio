package clio

import (
	partybus "github.com/nextlinux/gopartybus"
)

type BusConstructor func(Config) *partybus.Bus

var _ BusConstructor = newBus

func newBus(_ Config) *partybus.Bus {
	return partybus.NewBus()
}

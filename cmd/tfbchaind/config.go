package main

import (
	"github.com/threefoldtech/rivine/pkg/daemon"
)

// ExtendedDaemonConfig contains all configurable variables for the deamon.
type ExtendedDaemonConfig struct {
	daemon.Config
}

// DefaultConfig returns the default daemon configuration
func DefaultConfig() daemon.Config {
	cfg := daemon.DefaultConfig()
	cfg.APIaddr = "21110"
	cfg.RPCaddr = ":21112"
	return cfg
}

package options

import (
	"flag"
)

// Options has all the params needed to run a mcp server.
type Options struct {
	Port int

	// Config file needed to load, configs/config-dev.yaml for default.
	Config string
}

// NewOptions returns default mcp server options.
func NewOptions() *Options {
	var opts Options

	flag.IntVar(
		&opts.Port,
		"port",
		8080,
		"The port the server binds to.",
	)

	flag.StringVar(
		&opts.Config,
		"config",
		"configs/config.yaml",
		"The config file need to be loaded.",
	)

	return &opts
}

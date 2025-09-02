package main

import (
	"WildIP-Resolver-DNS/pkg"
	"WildIP-Resolver-DNS/pkg/config"
	"os"
)

// Initialize the configuration
func init() {
	debug := os.Getenv("DEBUG")
	config.LoadConfig(debug == "1" || debug == "true")
}

func main() {
	// Start the DNS server
	pkg.DNS()
}

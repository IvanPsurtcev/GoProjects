package main

import (
	"fmt"

	"metamask-auth/config"
)

func main() {
	cfg := config.CreateConfig()

	dsn := cfg.GetDsn()

	fmt.Printf("%s", dsn)
}

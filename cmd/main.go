package main

import (
	"github.com/gokul656/golbe/config"
	"github.com/gokul656/golbe/internal"
)

func init() {
	config.LoadEnvConfig()
	config.LoadLBConfig()

	internal.SetupLb()
}

func main() {
	internal.ListenAndServeV2()
}

package main

import (
  "os"

  env "github.com/caitlinelfring/go-env-default"

  "github.com/alionapermes/web-cv/config"
  "github.com/alionapermes/web-cv/ui"
)

const defaultConfigPath = "config.yaml"

func run() int {
  configPath := env.GetDefault("CONFIG_PATH", defaultConfigPath)

  config := config.New(configPath)
  ui := ui.New(config)

  if err := ui.Start(); err != nil {
    return 1
  }

  return 0
}

func main() {
  os.Exit(run())
}


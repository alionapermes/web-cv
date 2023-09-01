package config

import (
  "os"

	yaml "gopkg.in/yaml.v3"
)

type Config struct {
  Contacts struct {
    Telegram string `yaml:"telegram"`
    Email    string `yaml:"email"`
    GitHub   string `yaml:"github"`
  } `yaml:"contacts"`
  Texts    struct {
    Experience string `yaml:"experience"`
    About      string `yaml:"about"`
    Education  string `yaml:"education"`
    Status     string `yaml:"status"`
  } `yaml:"texts"`
  PhotoPath string `yaml:"photo-path"`
}

func New(path string) *Config {
  var cfg Config

  data, err := os.ReadFile(path)
  if err != nil {
    panic(err)
  }

  if err = yaml.Unmarshal(data, &cfg); err != nil {
    panic(err)
  }

  return &cfg
}


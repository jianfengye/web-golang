package main

import (
    "flag"
    "utility/configs"
)

var Config = newConfig()

func newConfig() *configs.Config {
    var configFile string
    flag.StringVar(&configFile, "f", "", "config file")
    flag.Parse()
    if configFile == "" {
        panic("usage: ./monitor -f etc/monitor.conf")
    }
    Config := configs.NewConfig()
    if err := Config.Load(configFile); err != nil {
        panic(err.Error())
    }
    return Config
}

package main

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/pumphouse-p/peek-go/pkg/app"
	"github.com/pumphouse-p/peek-go/pkg/version"
)

func main() {
	app := app.NewApp()

	v := viper.GetViper()

	app.BindConfig(v, pflag.CommandLine)

	pflag.Parse()

	log.Printf("Starting peek version: %v", version.VERSION)
	log.Println(strings.Repeat("=", 80))

	dumpConfig(v)
	log.Println(strings.Repeat("=", 80))

	app.LoadConfig(v)
	app.Run()
}

func dumpConfig(v *viper.Viper) {
	settings, err := json.MarshalIndent(v.AllSettings(), "", " ")
	if err != nil {
		log.Printf("Could not dump config: %v", err)
	}
	log.Printf("Config: \n%v\n", string(settings))
}

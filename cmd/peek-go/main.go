package main

import (
	"github.com/spf13/viper"

	"github.com/pumphouse-p/peek-go/pkg/app"
)

func main() {
	app := app.NewApp()

	v := viper.GetViper()
}

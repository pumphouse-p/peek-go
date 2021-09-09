package app

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	Debug    bool
	ListenOn string `mapstructure:listenon`
}

func (k *App) BindConfig(v *viper.Viper, fs *pflag.FlagSet) {
	k.kg.BindConfig(v, fs)
}

func (k *App) LoadConfig(v *viper.Viper) {
	err := v.UnmarshalExact(&k.c)
	if err != nil {
		panic(err)
	}

}

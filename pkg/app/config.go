package app

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	Debug    bool
	ListenOn string `mapstructure:"listen-on"`
}

func (p *App) BindConfig(v *viper.Viper, fs *pflag.FlagSet) {
	fs.Bool("debug", false, "Controls debug mode")
	v.BindPFlag("debug", fs.Lookup("debug"))
	fs.String("listen-on", ":8080", "The address to serve on")
	v.BindPFlag("listen-on", fs.Lookup("listen-on"))
}

func (p *App) LoadConfig(v *viper.Viper) {
	err := v.UnmarshalExact(&p.config)
	if err != nil {
		panic(err)
	}

}

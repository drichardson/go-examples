package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type S1 struct {
	Int int
	Str string
}

var boolFlag1 = pflag.Bool("bool1", false, "Bool flag 1")
var boolFlag2 = pflag.BoolP("bool2", "b", false, "Bool flag 2")

var goBoolFlag1 = flag.Bool("gobool1", false, "Go Bool flag 1")

func main() {
	viper.SetEnvPrefix("viper")

	viper.SetDefault("intopt", 10)
	viper.SetDefault("stropt", "ten")
	viper.SetDefault("durationopt", 10*time.Millisecond)
	viper.SetDefault("mapopt", map[string]string{"k1": "v1", "k2": "v2"})
	viper.SetDefault("structopt", &S1{Int: 20, Str: "twenty"})
	viper.SetDefault("watch", false)

	viper.BindEnv("envopt") // VIPER_ENVOPT

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s\n", err))
	}

	printValues()

	if viper.GetBool("watch") {
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("CONFIG FILE CHANGED:", e.Name, e.Op)
			printValues()
		})

		// block forever
		select {}
		panic("Should never get here")
	}
}

func printValues() {
	fmt.Println("intopt:", viper.GetInt("intopt"))
	fmt.Println("stropt:", viper.GetString("stropt"))
	fmt.Println("durationopt:", viper.GetDuration("durationopt"))
	fmt.Println("mapopt:", viper.GetStringMapString("mapopt"))
	fmt.Println("structopt:", viper.Get("structopt"))
	fmt.Println("envopt: (as string):", viper.GetString("envopt"), ", (as int)", viper.GetInt("envopt"))
	fmt.Println("bool1:", *boolFlag1, ", from viper:", viper.GetBool("bool1"))
	fmt.Println("bool2:", *boolFlag2, ", from viper:", viper.GetBool("bool2"))
	fmt.Println("goBool1:", *goBoolFlag1, ", from viper:", viper.GetBool("goBool1"))
}

package config

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type TestChart struct {
	BaseChart `mapstructure:",squash"`
	Test	string
}
type DeployChart struct {
		Enabled bool
		BaseChart `mapstructure:",squash"`
		ChartVar string
		Env string
		User string
}

type BaseChart struct {
	Token string
	Scaling string
	Creds string
}


func ViperLoadChartConfig() DeployChart {
	c:= DeployChart{}
	viper := viper.New()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		fmt.Println("Error binding")
	}
	pflag.Parse()
	err = viper.Unmarshal(&c)
	if err != nil {
		fmt.Println("Error unmarshaling")
	}
	return c
}

func ViperLoadTestConfig() TestChart {
	t := TestChart{}

	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		fmt.Println("Error binding")
		}
	pflag.Parse()
	err = viper.Unmarshal(&t)
	if err != nil {
		fmt.Println("Error unmarshaling")
	}
	return t
}

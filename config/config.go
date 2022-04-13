package config

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type TestChart struct {
	BaseChart `mapstructure:",squash"`
	Test      string `yaml:"Test"`
}
type DeployChart struct {
	Enabled   bool
	BaseChart `mapstructure:",squash"`
	ChartVar  string
	Env       string
}

type BaseChart struct {
	Token   string `yaml:"Token"`
	Scaling string `yaml:"Scaling"`
	Creds   string `yaml:"Creds"`
	User    string
}

func ViperLoadChartConfig() DeployChart {
	c := DeployChart{}
	v := viper.New()

	err := v.BindPFlags(pflag.CommandLine)
	if err != nil {
		fmt.Println("error")
	}
	pflag.Parse()
	err = v.Unmarshal(&c)
	if err != nil {
		fmt.Println("unmarshal error")
	}
	return c
}

func ViperLoadTestConfig() TestChart {
	t := TestChart{}
	v := viper.New()

	// err := v.BindPFlags(pflag.CommandLine)
	// if err != nil {
	// 	fmt.Println("error")
	// }
	pflag.Parse()
	err := v.Unmarshal(&t)
	if err != nil {
		fmt.Println("unmarshal error")
	}
	return t
}

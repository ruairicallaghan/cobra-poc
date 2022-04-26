package config

import (
	"fmt"

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

	err := viper.Unmarshal(&c)
	if err != nil {
		fmt.Println("unmarshal error")
	}
	return c
}

func ViperLoadTestConfig() TestChart {
	t := TestChart{}

	err := viper.Unmarshal(&t)
	if err != nil {
		fmt.Println("unmarshal error")
	}
	return t
}

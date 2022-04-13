package config

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type TestChart struct {
	BaseChart `mapstructure:",squash"`
	Test	string
	User	string
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
	c := DeployChart{}
	v := viper.New()
	pflag.String("Token", "defaultToken", "token usage")
	pflag.String("Scaling", "defaultScaling", "scaling usage")
	pflag.String("Creds", "defaultCreds", "creds usage")

	pflag.Bool("Enabled", true, "enabled or not")
	pflag.String("ChartVar", "defaultChart", "chart variable help")
	pflag.String("Env", "uat", "env variable")
	pflag.String("User", "defaultUser", "user variable")

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
	pflag.String("Token", "defaultToken", "token usage")
	pflag.String("Scaling", "defaultScaling", "scaling usage")
	pflag.String("Creds", "defaultCreds", "creds usage")

	pflag.String("Test", "defaultTest", "test usage")
	pflag.String("User", "defaultUser", "user usage")

	err := v.BindPFlags(pflag.CommandLine)
	if err != nil {
		fmt.Println("error")
	}
	pflag.Parse()
	err = v.Unmarshal(&t)
	if err != nil {
		fmt.Println("unmarshal error")
	}
	return t
}

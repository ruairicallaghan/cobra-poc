/*
Copyright © 2022 Ruairi Callaghan <ruairiscallaghan@gmail.com>

*/
package main

import (
	"github.com/callrua/cobra-poc/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Token   string
	Scaling string
	Creds   string
	User    string
)

func main() {
	var rootCmd = &cobra.Command{Use: "resource-generator"}
	rootCmd.AddCommand(cmd.ChartCmd, cmd.TestChartCmd)

	rootCmd.PersistentFlags().String("Token", "defaultToken", "Token usage")
	rootCmd.PersistentFlags().String("Scaling", "defaultScaling", "Scaling usage")
	rootCmd.PersistentFlags().String("Creds", "defaultCreds", "Creds usage")
	rootCmd.PersistentFlags().String("User", "defaultUser", "User usage")
	viper.BindPFlags(rootCmd.PersistentFlags())

	err := rootCmd.Execute()
	if err != nil {
		return
	}
}

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

	rootCmd.PersistentFlags().StringVar(&Token, "Token", "defaultToken", "Token usage")
	viper.BindPFlag("Token", rootCmd.PersistentFlags().Lookup("Token"))
	rootCmd.PersistentFlags().StringVar(&Scaling, "Scaling", "defaultScaling", "Scaling usage")
	viper.BindPFlag("Scaling", rootCmd.PersistentFlags().Lookup("Scaling"))
	rootCmd.PersistentFlags().StringVar(&Creds, "Creds", "defaultCreds", "Creds usage")
	viper.BindPFlag("Creds", rootCmd.PersistentFlags().Lookup("Creds"))
	rootCmd.PersistentFlags().StringVar(&User, "User", "defaultUser", "User usage")
	viper.BindPFlag("User", rootCmd.PersistentFlags().Lookup("User"))

	err := rootCmd.Execute()
	if err != nil {
		return
	}
}

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	
	// Global flags
	// User		string
	// Env			string
	Token 	string
	Scaling string
	Creds 	string

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "cobra-poc",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
	examples and usage of using your application. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmdFlags(rootCmd)
}

func RootCmdFlags(cmd *cobra.Command) {
	pflag.String("Token", "defaultToken", "token for app")
	pflag.String("Scaling", "defaultScaling", "scaling for app")
	pflag.String("Creds", "defaultCreds", "creds for app")
	viper.BindPFlags(pflag.CommandLine)
}

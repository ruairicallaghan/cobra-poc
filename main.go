/*
Copyright © 2022 Ruairi Callaghan <ruairiscallaghan@gmail.com>

*/
package main

import (
	"github.com/callrua/cobra-poc/cmd"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "resource-generator"}
	rootCmd.AddCommand(cmd.ChartCmd, cmd.TestChartCmd)
	err := rootCmd.Execute()
	if err != nil {
		return
	}
}

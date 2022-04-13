package cmd

import (
	"fmt"

	"github.com/callrua/cobra-poc/config"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var ( 
	
	// Local flags
	Enabled 	bool
	ChartVar 	string
	Env				string
	User			string
	ChartCmd = &cobra.Command{
		Use: "chart",
		Short: "short",
		Long: "long",
		Run: RunChart,
	}
)

func init() {
	ChartCmdFlags(ChartCmd)
	rootCmd.AddCommand(ChartCmd)
}

func ChartCmdFlags(cmd *cobra.Command) {
	pflag.Bool("Enabled", true, "enabled or not")
	pflag.String("ChartVar", "defaultChart", "chart variable help")
	pflag.String("Env", "uat", "env variable")
	pflag.String("User", "defaultUser", "user variable")
}

func RunChart(_ *cobra.Command, _ []string) {

	chartConfig := config.ViperLoadChartConfig()
	fmt.Println(chartConfig)

}

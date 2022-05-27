package cmd

import (
	"fmt"

	"github.com/callrua/cobra-poc/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewChartCommand() *cobra.Command {

	var chartCmd = &cobra.Command{
		Use:   "chart",
		Short: "short",
		Long:  "long",
		Run:   runChart,
	}

	chartCmd.Flags().Bool("Enabled", true, "Enabled usage")
	chartCmd.Flags().String("ChartVar", "defaultChartVar", "ChartVar usage")
	chartCmd.Flags().String("Env", "defaultEnv", "Env usage")
	chartCmd.Flags().String("Prometheus", "yes", "Env usage")
	return chartCmd

}

func runChart(cmd *cobra.Command, _ []string) {
	viper.BindPFlags(cmd.Flags())
	viperConfig := config.ViperLoadChartConfig()
	fmt.Println(viperConfig)
}

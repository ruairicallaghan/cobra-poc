package cmd

import (
	"fmt"

	"github.com/callrua/cobra-poc/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ChartCmd = &cobra.Command{
	Use:   "chart",
	Short: "short",
	Long:  "long",
	Run:   runChart,
}

func init() {
	ChartCmd.Flags().Bool("Enabled", true, "Enabled usage")
	ChartCmd.Flags().String("ChartVar", "defaultChartVar", "ChartVar usage")
	ChartCmd.Flags().String("Env", "defaultEnv", "Env usage")

}

func runChart(cmd *cobra.Command, _ []string) {
	viper.BindPFlags(cmd.Flags())

	viperConfig := config.ViperLoadChartConfig()
	fmt.Println(viperConfig)

}

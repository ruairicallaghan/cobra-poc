package cmd

import (
	"fmt"

	"github.com/callrua/cobra-poc/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var TestChartCmd = &cobra.Command{
	Use:   "test-chart",
	Short: "short",
	Long:  "long",
	Run:   runTestChart,
}

func init() {
	TestChartCmd.Flags().String("Test", "defaultTest", "Test usage")
}

func runTestChart(cmd *cobra.Command, _ []string) {
	viper.BindPFlags(cmd.Flags())

	viperConfig := config.ViperLoadTestConfig()
	fmt.Println(viperConfig)
}

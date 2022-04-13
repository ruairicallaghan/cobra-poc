package cmd

import (
	"fmt"

	"github.com/callrua/cobra-poc/config"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var ( 

	Test string
	TestChartCmd = &cobra.Command{
		Use: "test-chart",
		Short: "short",
		Long: "long",
		Run: runTestChart,
	}
)

func init() {
	TestCmdFlags(TestChartCmd)
	rootCmd.AddCommand(TestChartCmd)
}

func TestCmdFlags(cmd *cobra.Command) {
	pflag.String("Test", "regressionTest", "type of test")
}

func runTestChart(_ *cobra.Command, _ []string) {
	testConfig := config.ViperLoadTestConfig()
	fmt.Println(testConfig)
}

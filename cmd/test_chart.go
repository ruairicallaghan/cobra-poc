package cmd

import (
	"fmt"

	"github.com/callrua/cobra-poc/config"
	"github.com/spf13/cobra"
)

var ( 
	TestChartCmd = &cobra.Command{
		Use: "test-chart",
		Short: "short",
		Long: "long",
		Run: runTestChart,
	}
)

func runTestChart(_ *cobra.Command, _ []string) {
	testConfig := config.ViperLoadTestConfig()
	fmt.Println(testConfig)
}

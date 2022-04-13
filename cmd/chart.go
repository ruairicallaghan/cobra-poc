package cmd

import (
	"fmt"

	"github.com/callrua/cobra-poc/config"
	"github.com/spf13/cobra"
)

var ( 
	ChartCmd = &cobra.Command{
		Use: "chart",
		Short: "short",
		Long: "long",
		Run: runChart,
	}
)

func runChart(_ *cobra.Command, _ []string) {
	chartConfig := config.ViperLoadChartConfig()
	fmt.Println(chartConfig)
}

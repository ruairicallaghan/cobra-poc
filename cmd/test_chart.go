package cmd

import (
	"fmt"

	"github.com/callrua/cobra-poc/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	T            config.TestChart
	Test         string
	TestChartCmd = &cobra.Command{
		Use:   "test-chart",
		Short: "short",
		Long:  "long",
		Run:   runTestChart,
	}
)

func init() {
	TestCmdFlags(TestChartCmd)
}

func TestCmdFlags(cmd *cobra.Command) {

	cmd.Flags().StringVar(&Test, "Test", "defaultTest", "Test usage")
	viper.BindPFlag("Test", cmd.Flags().Lookup("Test"))
}

func runTestChart(cmd *cobra.Command, _ []string) {
	err := viper.Unmarshal(&T)
	if err != nil {
		fmt.Println("error unmarshalling")
	}
	fmt.Println(T)
}

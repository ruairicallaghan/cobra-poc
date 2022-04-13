package cmd

import (
	"fmt"

	"github.com/callrua/cobra-poc/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	C        config.DeployChart
	Enabled  bool
	ChartVar string
	Env      string
	ChartCmd = &cobra.Command{
		Use:   "chart",
		Short: "short",
		Long:  "long",
		Run:   runChart,
	}
)

func init() {
	ChartCmdFlags(ChartCmd)
}

func ChartCmdFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVar(&Enabled, "Enabled", true, "Enabled usage")
	viper.BindPFlag("Enabled", cmd.Flags().Lookup("Enabled"))
	cmd.Flags().StringVar(&ChartVar, "ChartVar", "defaultChartVar", "ChartVar usage")
	viper.BindPFlag("ChartVar", cmd.Flags().Lookup("ChartVar"))
	cmd.Flags().StringVar(&Env, "Env", "defaultEnv", "Env usage")
	viper.BindPFlag("Env", cmd.Flags().Lookup("Env"))

}

func runChart(cmd *cobra.Command, _ []string) {
	err := viper.Unmarshal(&C)
	if err != nil {
		fmt.Println("error unmarshalling")
	}
	fmt.Println(C)
}

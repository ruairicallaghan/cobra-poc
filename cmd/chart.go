package cmd

import (
	"fmt"

	"github.com/callrua/cobra-poc/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Test     string
	User     string
	Token    string
	Scaling  string
	Creds    string
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
	// pflag.String("Token", "defaultToken", "token usage")
	// pflag.String("Scaling", "defaultScaling", "scaling usage")
	// pflag.String("Creds", "defaultCreds", "creds usage")

	// pflag.Bool("Enabled", true, "enabled or not")
	// pflag.String("ChartVar", "defaultChart", "chart variable help")
	// pflag.String("Env", "uat", "env variable")
	// pflag.String("User", "defaultUser", "user variable")

	cmd.Flags().StringVar(&Token, "Token", "defaultToken", "Token usage")
	viper.BindPFlag("Token", cmd.PersistentFlags().Lookup("Token"))
	cmd.Flags().StringVar(&Scaling, "Scaling", "defaultScaling", "Scaling usage")
	viper.BindPFlag("Scaling", cmd.PersistentFlags().Lookup("Scaling"))
	cmd.Flags().StringVar(&Creds, "Creds", "defaultCreds", "Creds usage")
	viper.BindPFlag("Creds", cmd.PersistentFlags().Lookup("Creds"))
	cmd.Flags().BoolVar(&Enabled, "Enabled", true, "Enabled usage")
	viper.BindPFlag("Enabled", cmd.PersistentFlags().Lookup("Enabled"))
	cmd.Flags().StringVar(&ChartVar, "ChartVar", "defaultChartVar", "ChartVar usage")
	viper.BindPFlag("ChartVar", cmd.PersistentFlags().Lookup("ChartVar"))
	cmd.Flags().StringVar(&Env, "Env", "defaultEnv", "Env usage")
	viper.BindPFlag("Env", cmd.PersistentFlags().Lookup("Env"))
	// cmd.Flags().StringVar(&User, "User", "defaultUser", "User usage")
	// viper.BindPFlag("User", cmd.PersistentFlags().Lookup("User"))

}

func runChart(_ *cobra.Command, _ []string) {
	chartConfig := config.ViperLoadChartConfig()
	fmt.Println(chartConfig)
}

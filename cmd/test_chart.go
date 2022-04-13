package cmd

import (
	"fmt"

	"github.com/callrua/cobra-poc/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	T            config.TestChart
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
	cmd.Flags().StringVar(&Token, "Token", "defaultToken", "Token usage")
	viper.BindPFlag("Token", cmd.PersistentFlags().Lookup("Token"))
	cmd.Flags().StringVar(&Scaling, "Scaling", "defaultScaling", "Scaling usage")
	viper.BindPFlag("Scaling", cmd.PersistentFlags().Lookup("Scaling"))
	cmd.Flags().StringVar(&Creds, "Creds", "defaultCreds", "Creds usage")
	viper.BindPFlag("Creds", cmd.PersistentFlags().Lookup("Creds"))
	cmd.Flags().StringVar(&Test, "Test", "defaultTest", "Test usage")
	viper.BindPFlag("Test", cmd.PersistentFlags().Lookup("Test"))
	cmd.Flags().StringVar(&User, "User", "defaultUser", "User usage")
	viper.BindPFlag("User", cmd.PersistentFlags().Lookup("User"))
}

func runTestChart(_ *cobra.Command, _ []string) {
	// testConfig := config.ViperLoadTestConfig()
	// t := config.TestChart{}
	err := viper.Unmarshal(&T)
	if err != nil {
		fmt.Println("error unmarshalling")
	}
	err = viper.UnmarshalKey("User", &User)
	if err != nil {
		fmt.Println("error UnmarshalKey")
	}
	T.User = User
	fmt.Println(T)
}

package cmd

import (
	"testing"
	"fmt"

	"github.com/spf13/cobra"
	// "github.com/stretchr/testify/assert"
)

func TestBasicChartCmd(t *testing.T) {
	testCmd := &cobra.Command{Use: "chart", Run: RunChart,}
	ChartCmdFlags(testCmd)
}

func TestChartConfigDefaults (t *testing.T) {
	testCmd := &cobra.Command{Use: "chart", Run: RunChart}

	fmt.Println(testCmd)	
	// assert.Equal(t, "defaultChart", testConfig.ChartVar)
}

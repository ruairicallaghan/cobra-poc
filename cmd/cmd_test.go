package cmd

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func Test_ChartCmd(t *testing.T) {
	cmd := NewChartCommand()
	cmd.SetArgs([]string{"--Env", "staging"})
	cmd.Execute()

	y := viper.Get("Env")
	assert.Equal(t, "staging", y)
}

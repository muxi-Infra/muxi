package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "v1.1.0"

var rootCmd = &cobra.Command{
	Use:     "muxi",
	Short:   "Muxi: A tool for Muxi Go microservice",
	Long:    `Muxi: A tool for Muxi Go microservice`,
	Version: version,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help() // 显示帮助信息
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

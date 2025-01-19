package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var initCmd = &cobra.Command{
	Use:   "init [project name]",
	Short: "Initialize a new project from Muxi RPC template",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		fmt.Printf("Initializing project: %s\n", projectName)

		// 创建项目文件夹
		if err := os.MkdirAll(projectName, 0755); err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
			return
		}

		// GitHub 仓库 URL
		repoURL := "https://github.com/muxi-Infra/Muxi-Micro-Layout"

		// 使用 git clone 克隆模板
		cmdClone := exec.Command("git", "clone", repoURL, projectName)
		cmdClone.Stdout = os.Stdout
		cmdClone.Stderr = os.Stderr
		if err := cmdClone.Run(); err != nil {
			fmt.Printf("Error cloning repo: %v\n", err)
			return
		}

		// 输出成功信息
		fmt.Printf("Project initialized from template in %s\n", projectName)
	},
}

// 注册 initCmd 到 rootCmd
func init() {
	rootCmd.AddCommand(initCmd)
}

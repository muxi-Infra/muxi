package cmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

var (
	repoURL string
	timeout string
)

func init() {
	repoURL = "https://github.com/muxi-Infra/Muxi-Micro-Layout"
	timeout = "60s"

	initCmd.Flags().StringVarP(&repoURL, "repo-url", "r", repoURL, "layout repo")
	initCmd.Flags().StringVarP(&timeout, "timeout", "t", timeout, "time out")
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init [project name]",
	Short: "Initialize a new project from Muxi RPC template",
	Long:  "Initialize a new project using a template from GitHub.",
	Run:   runInit,
}

func runInit(cmd *cobra.Command, args []string) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	t, err := time.ParseDuration(timeout)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()
	name := ""
	if len(args) == 0 {
		fmt.Println("Project name is required")
		return
	} else {
		name = args[0]
	}
	projectDir := filepath.Join(wd, name)
	fmt.Printf("🚀 Initializing project %s, layout repo is %s, please wait a moment.\n\n", name, repoURL)
	if err := cloneRepo(ctx, projectDir); err != nil {
		fmt.Printf("🚫 Failed to initialize project: %v\n", err)
		return
	}
	if err := initGoMod(projectDir, name); err != nil {
		fmt.Printf("🚫 Failed to initialize go mod: %v\n", err)
		return
	}

	fmt.Printf("\n🍺 Project initialization succeeded %s\n", name)
	fmt.Print("💻 Use the following command to start the project 👇:\n\n")
	fmt.Println("$ cd", name)
	fmt.Println("$ go mod tidy")
	fmt.Println("🤝 Thanks for using muxi")
}

func initGoMod(dir, moduleName string) error {
	cmd := exec.Command("go", "mod", "init", moduleName)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	fmt.Println(string(out))
	if err != nil {
		return err
	}
	return nil
}

// 暂时不使用这个函数，由于可能存在网络问题，让使用者自己 go mod tidy
func tidyGoMod(dir string) error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	fmt.Println(string(out))
	if err != nil {
		return err
	}
	return nil
}

func cloneRepo(ctx context.Context, to string) error {
	cmd := exec.CommandContext(ctx, "git", "clone", repoURL, to)
	out, err := cmd.CombinedOutput()
	fmt.Println(string(out))
	if err != nil {
		return err
	}

	// 删除 .git 文件夹
	gitDir := filepath.Join(to, ".git")
	if err := os.RemoveAll(gitDir); err != nil {
		fmt.Printf("🚫 Failed to remove .git directory: %v\n", err)
		fmt.Printf("❗️ Please remove .git directory manually\n")
		// 这里设置返回 nil
		// 因为删除 .git 不影响使用
		// 返回提示提醒 Muxier 手动删除
		return nil
	}

	return nil
}

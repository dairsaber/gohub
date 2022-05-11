package main

import (
	"fmt"
	"gohub/app/cmd"
	"gohub/app/cmd/make"
	"gohub/bootstrap"
	"gohub/pkg/config"
	"gohub/pkg/console"
	"os"

	btsConfig "gohub/config"

	"github.com/spf13/cobra"
)

func init() {
	// 加载 config 目录下的配置信息
	btsConfig.Initialize()
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   config.Get("app.name"),
		Short: "A simple forum project",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,
		// rootCmd 的所有子命令都会执行以下代码
		PersistentPreRun: func(command *cobra.Command, args []string) {
			// 配置初始化，依赖命令行 --env 参数
			config.InitConfig(cmd.Env)
			// 设置gin的一些配置
			bootstrap.SetupGin()
			// 初始化 Logger
			bootstrap.SetupLogger()
			// 初始化数据库
			bootstrap.SetupDB()
			// 初始化 Redis
			bootstrap.SetupRedis()

			// TODO 初始化缓存

		},
	}

	// 注册子命令
	rootCmd.AddCommand(cmd.CmdServe, cmd.CmdKey, cmd.CmdPlay, make.CmdMake, cmd.CmdMigrate, cmd.CmdDBSeed)

	// 配置默认运行 Web 服务
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

	// 注册全局参数，--env
	cmd.RegisterGlobalFlags(rootCmd)

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}

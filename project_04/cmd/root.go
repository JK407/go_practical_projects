package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// 定义一个命令
var rootCmd = &cobra.Command{
	//  命令名称
	Use: "root",
	//  命令简介
	Short: "root is the root command for the project",
	//  命令详细介绍
	Long: `root is the root command for the project.`,
	//  命令执行
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root cmd run begin")
		//  打印 flag
		fmt.Println(
			//  获取 flag 的值
			cmd.Flags().Lookup("viper").Value,
			cmd.PersistentFlags().Lookup("author").Value,
			cmd.PersistentFlags().Lookup("config").Value,
			cmd.PersistentFlags().Lookup("license").Value,
			cmd.Flags().Lookup("source").Value,
			//  使用 viper 获取配置
			viper.Get("author"),
			viper.Get("license"),
		)
		fmt.Println("root cmd run end")
	},
	//  开放子命令给外部调用
	TraverseChildren: true,
}

// Exec
// @Description 执行命令
// @Author Oberl-Fitzgerald 2024-07-15 16:11:47
func Exec() {
	rootCmd.Execute()
}

// 配置文件
var cfgFile string

// 授权
var userLicense string

// init
// @Description 初始化
// @Author Oberl-Fitzgerald 2024-07-15 16:12:13
func init() {
	cobra.OnInitialize(initConfig)
	//  按名称接收命令行参数
	rootCmd.PersistentFlags().Bool("viper", true, "")
	//	指定flag缩写
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "")
	//  通过指针，将值赋值到字段
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "")
	//  通过指针，将值赋值到字段,并指定flag缩写
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "")
	//  添加本地标志
	rootCmd.Flags().StringP("source", "s", "local", "")

	//  配置绑定
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("license", rootCmd.PersistentFlags().Lookup("license"))
	//  设置默认值
	viper.SetDefault("author", "default author")
	viper.SetDefault("license", "default license")
}

// initConfig
// @Description 初始化配置
// @Author Oberl-Fitzgerald 2024-07-15 16:12:17
func initConfig() {
	//  指定配置文件
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		//  获取用户目录
		home, err := os.UserHomeDir()
		//  检查错误
		cobra.CheckErr(err)
		//  设置配置文件路径
		viper.AddConfigPath(home)
		//  设置配置文件类型
		viper.SetConfigType("yaml")
		//  设置配置文件名称
		viper.SetConfigName(".cobra")
	}
	//  检查环境变量，将配置的键值加载到 Viper 中
	viper.AutomaticEnv()
	//  读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
	}
	fmt.Println("Using config file:", viper.ConfigFileUsed())
}

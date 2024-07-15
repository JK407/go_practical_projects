package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

// 定义一个命令
var curArgs = &cobra.Command{
	//  命令名称
	Use: "currentArgs",
	//  参数校验
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		if len(args) > 2 {
			return errors.New("accepts at most two arg")
		}
		return nil
	},
	//  命令执行
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root currentArgs cmd run begin")
		fmt.Println(args)
		fmt.Println("root currentArgs cmd run end")
	},
}

// 定义一个命令
var argsCheckCmd = &cobra.Command{
	Use:       "argsCheck",
	Args:      cobra.OnlyValidArgs,
	ValidArgs: []string{"a", "b", "123", "abc"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root argsCheck cmd run begin")
		fmt.Println(args)
		fmt.Println("root argsCheck cmd run end")
	},
}

// 初始化
func init() {
	//  添加命令
	rootCmd.AddCommand(curArgs)
	rootCmd.AddCommand(argsCheckCmd)
}

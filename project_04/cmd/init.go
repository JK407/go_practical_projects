package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init is the command to initialize the project",
	Long:  `init is the command to initialize the project.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root init cmd run begin")
		//  打印 flag
		fmt.Println(
			cmd.Flags().Lookup("viper").Value,
			//cmd.Flags().Lookup("author").Value,
			//cmd.Flags().Lookup("config").Value,
			cmd.Flags().Lookup("license").Value,
			//  访问不到 rootCmd.Flags().Lookup("source").Value
			//cmd.Flags().Lookup("source").Value,
			cmd.Parent().Flags().Lookup("source").Value,

			viper.Get("author"),
			viper.Get("license"),
		)
		fmt.Println("root init cmd run end")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

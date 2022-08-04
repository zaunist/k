/*
Copyright © 2022 zaunist <y.bz@foxmail.com>
DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE Version Nil, December 2022.
*/
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/zaunist/k/pkg/use"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "use special version kubectl",
	Long:  `use special version kubectl`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		use.Do(args[0])
	},
}

func init() {
	rootCmd.AddCommand(useCmd)

	useCmd.Flags().StringVar(&version, "version", "", "the kubectl version you want")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// useCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// useCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

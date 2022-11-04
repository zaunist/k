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
	Short: "Switch to specified version",
	Long:  `Switch to specified version`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		use.Do(args[0])
	},
}

func init() {
	rootCmd.AddCommand(useCmd)

	useCmd.Flags().StringVar(&version, "version", "", "the kubectl version you want")
}

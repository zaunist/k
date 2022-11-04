/*
Copyright © 2022 zaunist <y.bz@foxmail.com>
DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE Version Nil, December 2022.
*/
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/zaunist/k/pkg/conf"
	"github.com/zaunist/k/pkg/ls"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List installed versions",
	Long:  `List installed versions`,
	Run: func(cmd *cobra.Command, args []string) {
		ls.List(conf.VersonDir)
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}

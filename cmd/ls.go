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

var versionDir = "./.k/versions/"

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "ls local kubectl version",
	Long:  `ls local kubectl version`,
	Run: func(cmd *cobra.Command, args []string) {
		ls.List(conf.VersonDir)
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

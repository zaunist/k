/*
Copyright © 2022 zaunist <y.bz@foxmail.com>
DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE Version Nil, December 2022.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	version2 "github.com/zaunist/k/pkg/version"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the current version",
	Long:  `Print the current version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version2.DisplayVersion())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

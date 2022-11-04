/*
Copyright © 2022 zaunist <y.bz@foxmail.com>
DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE Version Nil, December 2022.
*/
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/zaunist/k/pkg/clean"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Remove files from the package download directory",
	Long:  `Remove files from the package download directory`,
	Run: func(cmd *cobra.Command, args []string) {
		clean.Do()
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}

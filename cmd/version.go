/*
Copyright © 2022 zaunist <y.bz@foxmail.com>
DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE Version Nil, December 2022.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	v "github.com/zaunist/k/pkg/version"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "The K version",
	Long:  `Display the current version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(v.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

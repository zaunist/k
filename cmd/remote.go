/*
Copyright © 2022 zaunist <y.bz@foxmail.com>
DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE Version Nil, December 2022.
*/
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/zaunist/k/pkg/remote"
)

// lsRemoteCmd represents the lsRemote command
var lsRemoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "Get the official available kubectl versions",
	Long:  `Get the official available kubectl versions`,
	Run: func(cmd *cobra.Command, args []string) {
		remote.List()
	},
}

func init() {
	rootCmd.AddCommand(lsRemoteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsRemoteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsRemoteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
Copyright © 2022 zaunist <y.bz@foxmail.com>
DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE Version Nil, December 2022.
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zaunist/k/pkg/uninstall"
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall kubectl version",
	Args:  cobra.ExactArgs(1),
	Long:  `use k uninstall --version=v1.23.0 to uninstall kubectl`,
	Run: func(cmd *cobra.Command, args []string) {
		uninstall.Do(args[0])
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)

	uninstallCmd.Flags().StringVar(&version, "version", "", "The kubectl version you want uninstall")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uninstallCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uninstallCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

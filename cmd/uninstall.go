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
	Short: "Uninstall a version of kubectl",
	Args:  cobra.ExactArgs(1),
	Long:  `Uninstall a version of kubectl`,
	Run: func(cmd *cobra.Command, args []string) {
		uninstall.Do(args[0])
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)

	uninstallCmd.Flags().StringVar(&version, "version", "", "The kubectl version you want uninstall")
}

/*
Copyright © 2022 zaunist <y.bz@foxmail.com>
DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE Version Nil, December 2022.
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zaunist/k/pkg/install"
)

var version string

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install special kubectl version with --version",
	Long:  `install special kubectl version with --version`,
	Run: func(cmd *cobra.Command, args []string) {
		install.Do(version)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	installCmd.Flags().StringVar(&version, "version", "", "kubectl version you want")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

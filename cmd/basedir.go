/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	log "github.com/taylormonacelli/reactnut/cmd/logging"
)

// basedirCmd represents the basedir command
var basedirCmd = &cobra.Command{
	Use:   "basedir",
	Args:  cobra.ExactArgs(1),
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Logger.Tracef("basedir called with argument: %s", args[0])
	},
}

func init() {
	rootCmd.AddCommand(basedirCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// basedirCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// basedirCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

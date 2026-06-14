/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/yama-is-bocchi/filetree-go/filetree"
)

var (
	depth int
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "filetree-go [path]",
	Short: "Display files in a tree structure",
	Long:  "A Go-based CLI tool that displays any directory in a tree structure.",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := "."
		if len(args) > 0 {
			path = args[0]
		}
		tree := filetree.New(path, depth)
		_, err := tree.WriteTo(os.Stdout)
		return err
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVarP(
		&depth,
		"depth",
		"d",
		1,
		"depth of the path to display",
	)
}

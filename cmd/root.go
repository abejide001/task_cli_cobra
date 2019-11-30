package cmd

import "github.com/spf13/cobra"

// RootCmd variable
var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a command line tool for managing task",
	Long:  "A command line tool for learning purpose",
}

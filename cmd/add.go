package cmd

import (
	"fmt"
	"strings"
	"task/db"

	"github.com/spf13/cobra"
)

// add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to the task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, "")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Task:", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
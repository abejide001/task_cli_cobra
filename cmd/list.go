package cmd

import (
	"fmt"
	"task/db"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all your task",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println(err.Error())
		}
		if len(tasks) == 0 {
			fmt.Println("No task to complete")
			return
		}
       for i, task := range tasks {
		   fmt.Printf("%d. %v \n", i+1, task.Value)
	   }
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

}

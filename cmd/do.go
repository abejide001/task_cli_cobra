package cmd

import (
	"fmt"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed", err)
			}
			ids = append(ids, id)
		}
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong", err.Error())
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid", id)
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark %d \n, as completed %s \n", id, err)
			} else {
				fmt.Printf("Mark as completed %d \n", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}

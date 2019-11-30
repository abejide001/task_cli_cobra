package main

import (
	"path/filepath"
	"task/cmd"
	"task/db"

	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbpath := filepath.Join(home, "tasks.db")
	if err := db.Init(dbpath); err != nil {
		panic(err)
	}
	cmd.RootCmd.Execute()
}

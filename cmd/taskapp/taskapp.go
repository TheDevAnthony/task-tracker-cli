package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/TheDevAnthony/task-tracker-cli/internal/cli"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: taskapp <command> [options]")
		os.Exit(1)
	}

	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	dir := filepath.Dir(exePath)
	fullPath := filepath.Join(dir, "tasks.txt")

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		f, err := os.Create(fullPath)
		if err != nil {
			log.Fatal(err)
		}
		f.Close()
	}

	switch strings.ToLower(os.Args[1]) {
	case "add":
		if err := cli.AddTask(os.Args[2]); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	case "update":
		if err := cli.UpdateTask(os.Args[2:]); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	case "delete":
		if err := cli.DeleteTask(os.Args[2]); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	case "mark-in-progres":
		if err := cli.MarkInProgress(os.Args[2]); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	case "mark-done":
		if err := cli.MarkDone(os.Args[2]); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	case "list":
		if err := cli.ListTasks(os.Args[2]); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	case "help":
		cli.PrintHelp()
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}

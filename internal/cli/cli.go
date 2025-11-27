package cli

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func AddTask(rawDesc string) error {
	errMsg := "Invalid argument. Usage: taskapp add \"new description\""

	desc := strings.TrimSpace(rawDesc)
	if desc == "" {
		return errors.New(errMsg)
	}

	// Update logic

	// fmt.Printf("Task added successfully (ID: )\n", id)
	return nil
}

func UpdateTask(args []string) error {
	errMsg := "Invalid arguments. Usage: taskapp update [id] \"new description\""

	if len(args) < 2 {
		return errors.New(errMsg)
	}

	id, err := strconv.Atoi(args[0])
	desc := strings.TrimSpace(args[1])
	if err != nil || desc == "" {
		return errors.New(errMsg)
	}

	// Update logic

	fmt.Println("Task updated successfully")
	return nil
}

func DeleteTask(strId string) error {
	errMsg := "Invalid argument. Usage: taskapp delete [id]"

	id, err := strconv.Atoi(strId)
	if err != nil {
		return errors.New(errMsg)
	}

	// Delete logic

	fmt.Println("Task deleted successfully")
	return nil
}

func MarkInProgress(strId string) error {
	errMsg := "Invalid argument. Usage: taskapp mark-in-progress [id]"

	id, err := strconv.Atoi(strId)
	if err != nil {
		return errors.New(errMsg)
	}

	// Mark as in progress logic

	return nil
}

func MarkDone(strId string) error {
	errMsg := "Invalid argument. Usage: taskapp mark-done [id]"

	id, err := strconv.Atoi(strId)
	if err != nil {
		return errors.New(errMsg)
	}

	// Mark as done logic

	return nil
}

func ListTasks(op string) error {
	errMsg := "Invalid argument. Usage: taskapp list [option]"

	// Listing logic

	return nil
}

func PrintHelp() {
	fmt.Println(`
taskapp - simple task tracker app

Usage:
    taskapp <command> [options]

Commands:
	task-cli add "Buy groceries"
	task-cli update 1 "Buy groceries and cook dinner"
	task-cli delete 1

	task-cli mark-in-progress 1
	task-cli mark-done 1

	task-cli list
	task-cli list done
	task-cli list todo
	task-cli list in-progress
	`)
}

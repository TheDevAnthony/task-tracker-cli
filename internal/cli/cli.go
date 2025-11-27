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
		if errors.Is(err, strconv.ErrSyntax) || desc == "" {
			return errors.New(errMsg)
		}
		return err
	}

	// Update logic

	fmt.Println("Task updated successfully")
	return nil
}

func DeleteTask(strId string) error {
	errMsg := "Invalid argument. Usage: taskapp delete [id]"

	id, err := strconv.Atoi(strId)
	if err != nil {
		if errors.Is(err, strconv.ErrSyntax) {
			return errors.New(errMsg)
		}
		return err
	}

	// Delete logic

	fmt.Println("Task deleted successfully")
	return nil
}

func MarkInProgress(strId string) error {
	errMsg := "Invalid argument. Usage: taskapp mark-in-progress [id]"

	id, err := strconv.Atoi(strId)
	if err != nil {
		if errors.Is(err, strconv.ErrSyntax) {
			return errors.New(errMsg)
		}
		return err
	}

	// Mark as in progress logic

	return nil
}

func MarkDone(strId string) error {
	errMsg := "Invalid argument. Usage: taskapp mark-done [id]"

	id, err := strconv.Atoi(strId)
	if err != nil {
		if errors.Is(err, strconv.ErrSyntax) {
			return errors.New(errMsg)
		}
		return err
	}

	// Mark as done logic

	return nil
}

func ListTasks(op string) error {
	errMsg := "Invalid argument. Usage: taskapp list OR taskapp list [option]"

	op = strings.ToLower(strings.TrimSpace(op))
	validOps := map[string]bool{
		"":            true,
		"done":        true,
		"todo":        true,
		"in-progress": true,
	}
	if !validOps[op] {
		return errors.New(errMsg)
	}

	// Listing logic

	return nil
}

func PrintHelp() {
	fmt.Println(`
taskapp - simple task tracker app

Usage:
    taskapp <command> [options]

Commands:
	task-cli add			 	"add a new task"
	task-cli update				"update a task"
	task-cli delete				"delete a task"

	task-cli mark-in-progress	"mark a task as in-progress"
	task-cli mark-done			"mark a task as done"

	task-cli list				"list all tasks"
	task-cli list done			"list completed tasks"
	task-cli list todo			"list pending tasks"
	task-cli list in-progress	"list tasks in progress"
	task-cli help				"display this help message"
	`)
}

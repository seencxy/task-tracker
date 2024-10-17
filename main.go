package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

// https://roadmap.sh/projects/task-tracker
func main() {
	flag.Parse()
	args := flag.Args()

	storeData, err := openStoreFile()
	if err != nil {
		fmt.Println("Open store file error: " + err.Error())
		return
	}

switchLabel:
	switch args[0] {
	case "add":
		if checkArgs(2) {
			nextId := 1
			if len(storeData) > 0 {
				nextId = storeData[len(storeData)-1].ID + 1
			}
			storeData = append(storeData, NewTask(nextId, args[1]))
			fmt.Printf("Task added successfully (ID: %d) \n", nextId)
		}
	case "update":
		if checkArgs(3) {
			id, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				fmt.Println("Please enter a valid ID")
				return
			}
			for index, element := range storeData {
				if element.ID == int(id) {
					storeData[index].update(WithDescOpts(args[2]))
					break switchLabel
				}
			}
			fmt.Println("Task not found")
		}
	case "delete":
		if checkArgs(2) {
			id, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				fmt.Println("Please enter a valid ID")
				return
			}
			for index, element := range storeData {
				if element.ID == int(id) {
					storeData = append(storeData[:index], storeData[index+1:]...)
					break switchLabel
				}
			}
			fmt.Println("Task not found")
		}
	case "mark-in-progress":
		if checkArgs(2) {
			id, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				fmt.Println("Please enter a valid ID")
				return
			}
			for index, element := range storeData {
				if element.ID == int(id) {
					storeData[index].update(WithStatusOpts(InProgress))
					break switchLabel
				}
			}
			fmt.Println("Task not found")
		}
	case "mark-done":
		if checkArgs(2) {
			id, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				fmt.Println("Please enter a valid ID")
				return
			}
			for index, element := range storeData {
				if element.ID == int(id) {
					storeData[index].update(WithStatusOpts(Done))
					break switchLabel
				}
			}
			fmt.Println("Task not found")
		}
	case "list":
		if len(args) == 1 {
			for _, element := range storeData {
				fmt.Printf("ID: %d, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n", element.ID, element.Description, element.Status, element.CreatedAt, element.UpdatedAt)
			}
			return
		}
		if len(args) == 2 {
			switch args[1] {
			case "done":
				displayTasks(storeData, Done)
			case "todo":
				displayTasks(storeData, Todo)
			case "in-progress":
				displayTasks(storeData, InProgress)
			default:
				fmt.Println("Please enter the operation command")
			}
			return
		}
		fmt.Println("Please enter the operation command")
	default:
		fmt.Println("Please enter the operation command")
		return
	}
	if err = writeStoreFile(storeData); err != nil {
		fmt.Println("Write store file error: " + err.Error())
		return
	}
}
func checkArgs(count int) bool {
	if len(flag.Args()) != count {
		fmt.Println("Please enter the operation command")
		return false
	}
	return true
}

func openStoreFile() ([]*Task, error) {
	file, err := os.OpenFile("store.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileBody, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if len(fileBody) == 0 {
		return nil, nil
	}

	return unmarshalTasks(fileBody)
}

func writeStoreFile(storeData []*Task) error {
	file, err := os.OpenFile("store.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	bodyBytes, err := json.MarshalIndent(storeData, "", "    ")
	if err != nil {
		return err
	}
	_, err = file.Write(bodyBytes)
	return err
}

func displayTasks(storeData []*Task, status string) {
	for _, element := range storeData {
		if element.Status != status {
			continue
		}
		fmt.Printf("ID: %d, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n", element.ID, element.Description, element.Status, element.CreatedAt, element.UpdatedAt)
	}
}

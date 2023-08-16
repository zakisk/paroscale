package main

import (
	"fmt"
	"os"

	"github.com/shirou/gopsutil/v3/process"
)

func main() {
	processes, err := process.Processes()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	checkArgs()

	fileName := os.Args[1]

	// Open the file for writing, creating it if it doesn't exist
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	output := ""

    for _, process := range processes {
		name, err := process.Name()
		if err != nil {
			name = ""
		}

		userName, err := process.Username()
		if err != nil {
			userName = ""
		}
		output += fmt.Sprintf("%d\t|\t%s\t|\t%s\n", process.Pid, name, userName)
    }

	// Write data to the file
	_, err = file.WriteString(output)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}
}

func checkArgs() {
	if len(os.Args) < 2 {
		err := fmt.Errorf("[ERROR] no argument passed for directory")
		fmt.Println(err)
		os.Exit(1)
	}
}
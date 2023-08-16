
// inspired by https://github.com/kddnewton/tree/blob/main/tree.go
package main

import (
	// "fmt"
	"fmt"
	"os"
	"path"
)


var totalDirs, totalFiels int

func main() {
	checkArgs()
	printDir(os.Args[1], "")
}

func checkArgs() {
	if len(os.Args) < 2 {
		err := fmt.Errorf("[ERROR] no argument passed for directory")
		fmt.Println(err)
		os.Exit(1)
	}
}

func printDir(base string, prefix string) {
	names := getDirsAndFiles(base)

	for index, name := range names {
		fullPath := path.Join(base, name)
		stat, err := os.Stat(fullPath)
		if err != nil {
			err := fmt.Errorf("[ERROR] during getting file info")
			fmt.Println(err)
			os.Exit(1)
		}
		dirDetails := ""
		if stat.IsDir() {
			dirs, files := getDirAndFilesCount(fullPath)
			dirDetails = fmt.Sprintf("(Directories : %d, Files: %d)", dirs, files)
		}
		if index == len(names)-1 {
			fmt.Printf("%s└──%s%s\n", prefix, name, dirDetails)
			if stat.IsDir() {
				printDir(fullPath, prefix+"    ")
			}
		} else {
			fmt.Printf("%s├──%s%s\n", prefix, name, dirDetails)
			if stat.IsDir() {
				printDir(fullPath, prefix+"│   ")
			}
		}
	}
}

func getDirsAndFiles(basePath string) []string {
	file, err := os.Open(basePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	items, err := file.Readdirnames(0)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return items
}

func getDirAndFilesCount(fullPath string) (dirs, files int) {
	names := getDirsAndFiles(fullPath)
	for _, item := range names {
		stat, err := os.Stat(path.Join(fullPath, item))
		if err != nil {
			err := fmt.Errorf("[ERROR] during getting file info")
			fmt.Println(err)
			os.Exit(1)
		}
		if stat.IsDir() {
			dirs++
		} else {
			files++
		}
	}
	return dirs, files
}

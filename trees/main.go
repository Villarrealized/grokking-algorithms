package main

import (
	"container/list"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	printFilesBFS("../")
	fmt.Println("")
	printFilesDFS("../")
}

// print all files recursively beginning with directory
func printFilesBFS(startDir string) {
	dirQ := list.New()
	dirQ.PushBack(startDir)

	dirsSeen := map[string]bool{}
	for dirQ.Len() > 0 {
		dir := dirQ.Front()
		dirQ.Remove(dir)

		dirName := dir.Value.(string)
		// Ignore directories beginning with .git
		if strings.Contains(dirName, ".git") {
			continue
		}

		items, err := os.ReadDir(dirName)
		if err != nil {
			fmt.Printf("Error: %v", err)
			continue
		}
		for _, item := range items {
			fullPath := filepath.Join(dirName, item.Name())
			if item.IsDir() {
				dirQ.PushBack(fullPath)
			} else {
				if !dirsSeen[dirName] {
					if dirName == startDir {
						fmt.Println(dirName)
					} else {
						fmt.Println(strings.Replace(dirName, startDir, "", 1) + "/")
					}
					dirsSeen[dirName] = true
				}
				fmt.Printf("  - %s\n", item.Name())
			}
		}
	}
}

func printFilesDFS(dir string) {
	startDir := dir

	var dfs func(dir string)
	dfs = func(dir string) {
		// Ignore directories beginning with .git
		if strings.Contains(dir, ".git") {
			return
		}

		items, err := os.ReadDir(dir)
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}

		for _, item := range items {
			fullPath := filepath.Join(dir, item.Name())
			if item.IsDir() {
				dfs(fullPath)
			} else {
				if dir == startDir {
					fmt.Println(dir)
				} else {
					fmt.Println(strings.Replace(dir, startDir, "", 1) + "/")
				}
				fmt.Printf("  - %s\n", item.Name())
			}
		}
	}
	dfs(dir)
}

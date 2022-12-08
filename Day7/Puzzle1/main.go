package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type File struct {
	FileSize int
	FileName string
}
type Directory struct {
	Files           []File
	FileSize        int
	DirectoryName   string
	ParentDirectory *Directory
	SubDirectories  []*Directory
}

func (d *Directory) GetFullDirectoryPathName() string {
	result := d.DirectoryName
	parent := d.ParentDirectory
	for parent != nil {
		result = fmt.Sprintf("%s%s", parent.DirectoryName, result)
		parent = parent.ParentDirectory
	}
	return result
}
func IsCommandLine(line string) bool {
	if line[:1] == "$" {
		return true
	} else {
		return false
	}
}
func WorkingDirectory() string {
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mydir)
	return mydir
}

func ProcessCommand(currentDirectory *Directory, line string) *Directory {
	result := currentDirectory
	if line[2:4] == "cd" {
		directoryName := line[5:]

		if directoryName == "/" {
			fmt.Println("-----ProcessCommand: this is root directory")
		} else if directoryName == ".." {
			result = currentDirectory.ParentDirectory

		} else {
			var newDirectory Directory
			newDirectory.DirectoryName = fmt.Sprintf("%s/", directoryName)
			newDirectory.ParentDirectory = currentDirectory
			currentDirectory.SubDirectories = append(currentDirectory.SubDirectories, &newDirectory)
			result = &newDirectory
		}
	}
	return result
}
func GenerateFileSize(directory *Directory) {
	fTotal := 0
	for _, f := range directory.Files {
		fTotal += f.FileSize
	}
	directory.FileSize = fTotal //only does current.
	for _, d := range directory.SubDirectories {
		GenerateFileSize(d)
		directory.FileSize += d.FileSize //this adds the subdirectories to the parent
	}
}
func PrintDirectory(directory *Directory) int {
	fmt.Printf("Directory %s has size %d ", directory.GetFullDirectoryPathName(), directory.FileSize)
	total := 0
	if directory.FileSize < 100000 {
		fmt.Printf(">>>>>> this is < 100000\n")
		total += directory.FileSize
	} else {
		fmt.Printf("\n")
	}
	for _, d := range directory.SubDirectories {
		total += PrintDirectory(d)
	}
	return total
}
func main() {
	fmt.Println("Starting Day 7, Puzzle 1")
	file, err := os.Open(fmt.Sprintf("%s/Day7/PuzzleInput.txt", WorkingDirectory()))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	//seed the root directory.
	var RootDirectory Directory
	RootDirectory.DirectoryName = "/"
	currentDirectory := &RootDirectory
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if IsCommandLine(line) {
			currentDirectory = ProcessCommand(currentDirectory, line)
		} else {
			parts := strings.Split(line, " ")
			//always have parts:
			//123 file.txt
			//dir name
			if parts[0][0:3] == "dir" {
			} else {
				var newFile File
				newFile.FileName = parts[1]
				newFile.FileSize, _ = strconv.Atoi(parts[0])
				currentDirectory.Files = append(currentDirectory.Files, newFile)
			}
		}
	}
	GenerateFileSize(&RootDirectory)
	total := PrintDirectory(&RootDirectory)
	fmt.Printf("Magic number: %d\n", total)
}

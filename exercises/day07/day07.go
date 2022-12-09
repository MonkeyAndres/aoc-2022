package day07

import (
	"strconv"
	"strings"
)

type File struct {
	name string
	size int
}

type ContentTree struct {
	files       map[string]File
	directories map[string]*Directory
}

type Directory struct {
	name     string
	size     int
	parent   *Directory
	contents ContentTree
}

func (dir *Directory) addFile(file File) {
	dir.contents.files[file.name] = file
}

func (dir *Directory) addDirectory(directory *Directory) {
	dir.contents.directories[directory.name] = directory
}

func calculateDirectorySize(directory *Directory) {
	for _, dir := range directory.contents.directories {
		calculateDirectorySize(dir)
		directory.size += dir.size
	}

	for _, file := range directory.contents.files {
		directory.size += file.size
	}
}

func buildFilesystem(terminalOutput []string, rootDirectory *Directory) {
	currentDirectory := rootDirectory

	for _, line := range terminalOutput {
		parts := strings.Split(line, " ")

		if parts[0] == "$" {
			if parts[1] == "ls" {
				continue
			}

			// NOTE: `cd` command

			if parts[2] == "/" {
				currentDirectory = rootDirectory
			} else if parts[2] == ".." {
				currentDirectory = currentDirectory.parent
			} else {
				currentDirectory = currentDirectory.contents.directories[parts[2]]
			}
		} else {
			if parts[0] == "dir" {
				newDirectory := Directory{
					parts[1],
					0,
					currentDirectory,
					ContentTree{make(map[string]File), make(map[string]*Directory)},
				}
				currentDirectory.addDirectory(&newDirectory)
			} else {
				numericSize, err := strconv.Atoi(parts[0])

				if err != nil {
					panic("Error converting file size to int")
				}

				newFile := File{parts[1], numericSize}
				currentDirectory.addFile(newFile)
			}

		}
	}

	calculateDirectorySize(rootDirectory)
}

const DELETION_SIZE_THRESHOLD = 100000

func sumDeletionCandidates(directory *Directory, sum int) int {
	for _, dir := range directory.contents.directories {
		if dir.size <= DELETION_SIZE_THRESHOLD {
			sum += dir.size
		}

		sum = sumDeletionCandidates(dir, sum)
	}

	return sum
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	rootDirectory := Directory{
		"/",
		0,
		nil,
		ContentTree{make(map[string]File), make(map[string]*Directory)},
	}

	buildFilesystem(lines, &rootDirectory)

	return sumDeletionCandidates(&rootDirectory, 0)
}

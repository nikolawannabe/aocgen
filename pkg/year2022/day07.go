package year2022

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Day07 struct{}

const (
	Dir               = "dir"
	File              = "file"
	Cd                = "cd"
	Ls                = "ls"
	limit             = 100000
	filesystemSize    = 70000000
	unusedSpaceNeeded = 30000000
	disSpaceTaken     = 44376732
)

type Directory struct {
	Name        string
	Files       []int
	Parent      *Directory
	Directories map[string]Directory
	Size        int
}

type CommandLine struct {
	Command string
	Operand *string
}

type Output struct {
	NodeType string
	FileSize *int
	Name     string
}

type HistoryEntry struct {
	Index   int
	Command *CommandLine
	Output  *Output
}

func parseCommand(i int, line string) HistoryEntry {
	commandEntry := line[2:]
	cmd := CommandLine{}
	parts := strings.Split(commandEntry, " ")
	if len(parts) > 2 {
		log.Printf("too many parts in command line: %s", commandEntry)
	}
	if len(parts) == 1 {
		cmd.Command = parts[0]
	}
	if len(parts) == 2 {
		cmd.Command = parts[0]
		p := parts[1]
		cmd.Operand = &p
	}
	return HistoryEntry{Index: i, Command: &cmd}
}

func parseOutput(line string) Output {
	out := Output{}
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		log.Printf("wrong len file/dir parts: %s", line)
	}
	if parts[0] == Dir {
		out.NodeType = Dir
		out.Name = parts[1]
		return out
	}

	var fileSize int
	var fileName string
	_, err := fmt.Sscanf(line, "%d %s", &fileSize, &fileName)
	if err != nil {
		log.Printf("couldn't read file parts: %v, %s", err, line)
	}
	out.FileSize = &fileSize
	out.Name = fileName
	out.NodeType = File
	return out
}

func NewDirectory(name string, parent *Directory) Directory {
	initialDirectories := make(map[string]Directory, 0)
	initialFiles := make([]int, 0)
	return Directory{Name: name, Directories: initialDirectories, Files: initialFiles, Parent: parent}
}

func enter(curRoot *Directory, name string) *Directory {
	log.Printf("old root: %v", curRoot)

	for _, dir := range curRoot.Directories {
		actualDir := &dir
		if actualDir.Name == name {
			actualDir.Parent = curRoot
			log.Printf("entered directory: %s", actualDir.Name)
			curRoot = actualDir
			log.Printf("New root: %v", curRoot)
			return curRoot
		}
	}
	return nil
}

func pop(curRoot *Directory) (*Directory, []int) {
	dirsToKeep := make([]int, 0)

	dirSize := 0
	for _, dir := range curRoot.Directories {
		dirSize = dirSize + dir.Size
	}
	for _, fileSize := range curRoot.Files {
		dirSize = dirSize + fileSize
	}
	if dirSize <= limit {
		log.Printf("size at most limit: %d vs %d", dirSize, limit)
		dirsToKeep = append(dirsToKeep, dirSize)
	}
	name := curRoot.Name
	curRoot.Size = dirSize
	log.Printf("old root: %v", curRoot)
	log.Printf("left directory: %s, found files: %d with size %d and %d dirs", curRoot.Name, len(curRoot.Files), dirSize, len(curRoot.Directories))
	oldRoot := &curRoot.Parent
	curRoot = *oldRoot
	dir, _ := curRoot.Directories[name]
	dir.Size = dirSize
	curRoot.Directories[name] = dir
	log.Printf("New root: %v", curRoot)
	return curRoot, dirsToKeep
}

func otherpop(curRoot *Directory) (*Directory, []int) {
	dirsToKeep := make([]int, 0)

	dirSize := getDirSize(curRoot)
	if dirSize >= (unusedSpaceNeeded+disSpaceTaken)-filesystemSize {
		log.Printf("size at most limit: %d vs %d", dirSize, limit)
		dirsToKeep = append(dirsToKeep, dirSize)
	}
	name := curRoot.Name
	curRoot.Size = dirSize
	log.Printf("old root: %v", curRoot)
	log.Printf("left directory: %s, found files: %d with size %d and %d dirs", curRoot.Name, len(curRoot.Files), dirSize, len(curRoot.Directories))
	oldRoot := &curRoot.Parent
	curRoot = *oldRoot
	dir, _ := curRoot.Directories[name]
	dir.Size = dirSize
	curRoot.Directories[name] = dir
	log.Printf("New root: %v", curRoot)
	return curRoot, dirsToKeep
}

func getDirSize(curRoot *Directory) int {
	dirSize := 0
	for _, dir := range curRoot.Directories {
		dirSize = dirSize + dir.Size
	}
	for _, fileSize := range curRoot.Files {
		dirSize = dirSize + fileSize
	}
	return dirSize

}

func (p Day07) PartA(lines []string) any {
	dirsToKeep := make([]int, 0)

	entries := make([]HistoryEntry, 0)
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		if line[0] == '$' {
			entry := parseCommand(i, line)
			entries = append(entries, entry)
			continue
		}
		output := parseOutput(line)
		entry := HistoryEntry{Index: i, Output: &output}
		entries = append(entries, entry)
	}

	initialDirectories := make(map[string]Directory, 0)
	initialFiles := make([]int, 0)
	initialRoot := Directory{Name: "/", Directories: initialDirectories, Files: initialFiles}
	curRoot := &initialRoot

	for _, entry := range entries {
		bytes, err := json.Marshal(entry)
		if err != nil {
			log.Printf("couldn't marshal %v", err)
		}
		log.Printf("%s", string(bytes))

		if entry.Command != nil && entry.Command.Command == Cd && *entry.Command.Operand != ".." {
			if *entry.Command.Operand == "/" {
				continue
			}
			curRoot = enter(curRoot, *entry.Command.Operand)
		}

		if entry.Command != nil && entry.Command.Command == Cd && *entry.Command.Operand == ".." {
			newRoot, newDirs := pop(curRoot)
			curRoot = newRoot
			dirsToKeep = append(dirsToKeep, newDirs...)
			continue
		}

		if entry.Output != nil && entry.Output.NodeType == File {
			curRoot.Files = append(curRoot.Files, *entry.Output.FileSize)
			curRoot.Size += *entry.Output.FileSize
		}
		if entry.Output != nil && entry.Output.NodeType == Dir {
			curRoot.Directories[*&entry.Output.Name] = NewDirectory(*&entry.Output.Name, curRoot)
		}
	}

	for curRoot.Name != "/" && curRoot.Parent != nil {
		newRoot, newDirs := pop(curRoot)
		curRoot = newRoot
		dirsToKeep = append(dirsToKeep, newDirs...)
	}

	totalSize := 0
	for _, dirSize := range dirsToKeep {
		totalSize += dirSize
	}
	return totalSize
}

func (p Day07) PartB(lines []string) any {
	dirsToDelete := make([]int, 0)

	entries := make([]HistoryEntry, 0)
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		if line[0] == '$' {
			entry := parseCommand(i, line)
			entries = append(entries, entry)
			continue
		}
		output := parseOutput(line)
		entry := HistoryEntry{Index: i, Output: &output}
		entries = append(entries, entry)
	}

	initialDirectories := make(map[string]Directory, 0)
	initialFiles := make([]int, 0)
	initialRoot := Directory{Name: "/", Directories: initialDirectories, Files: initialFiles}
	curRoot := &initialRoot

	for _, entry := range entries {
		bytes, err := json.Marshal(entry)
		if err != nil {
			log.Printf("couldn't marshal %v", err)
		}
		log.Printf("%s", string(bytes))

		if entry.Command != nil && entry.Command.Command == Cd && *entry.Command.Operand != ".." {
			if *entry.Command.Operand == "/" {
				continue
			}
			curRoot = enter(curRoot, *entry.Command.Operand)
		}

		if entry.Command != nil && entry.Command.Command == Cd && *entry.Command.Operand == ".." {
			newRoot, newDirs := otherpop(curRoot)
			curRoot = newRoot
			dirsToDelete = append(dirsToDelete, newDirs...)
			continue
		}

		if entry.Output != nil && entry.Output.NodeType == File {
			curRoot.Files = append(curRoot.Files, *entry.Output.FileSize)
			curRoot.Size += *entry.Output.FileSize
		}
		if entry.Output != nil && entry.Output.NodeType == Dir {
			curRoot.Directories[*&entry.Output.Name] = NewDirectory(*&entry.Output.Name, curRoot)
		}
	}

	for curRoot.Name != "/" && curRoot.Parent != nil {
		newRoot, newDirs := otherpop(curRoot)
		curRoot = newRoot
		dirsToDelete = append(dirsToDelete, newDirs...)
	}

	log.Printf("disk size: %d", getDirSize(curRoot))
	least := dirsToDelete[0]
	for _, dirSize := range dirsToDelete {
		if dirSize < least {
			least = dirSize
		}
	}
	log.Printf("%v", dirsToDelete)
	return least
}

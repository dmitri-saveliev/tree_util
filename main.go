package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, showFiles bool) error {
	fi, err := os.Lstat(path)
	if err != nil || !fi.IsDir() {
		return fmt.Errorf("an error occured while opening file")
	}

	return printDir(out, path, showFiles, 0, 0)
}

func printDir(out io.Writer, path string, showFiles bool, level int, lastDirLvl int) error {
	sep := string(os.PathSeparator)

	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	if len(entries) == 0 {
		return nil
	}

	if !showFiles {
		entries = getOnlyDirs(entries)
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})

	for i, entry := range entries {
		isLast := i == len(entries)-1
		if entry.IsDir() {
			printPrefix(out, level, isLast, lastDirLvl)
			fmt.Fprintln(out, entry.Name())
			if isLast {
				lastDirLvl++
			}
			err = printDir(out, path+sep+entry.Name(), showFiles, level+1, lastDirLvl)
			if err != nil {
				return err
			}
		} else if showFiles {
			printPrefix(out, level, isLast, lastDirLvl)
			fi, err := entry.Info()
			if err == nil && fi != nil {
				size := fi.Size()
				sizeStr := ""
				if size == 0 {
					sizeStr = "empty"
				} else {
					sizeStr = fmt.Sprintf("%db", size)
				}
				fmt.Fprintf(out, "%s (%s)\n", fi.Name(), sizeStr)
			} else {
				return err
			}
		}
	}
	return nil
}

func printPrefix(out io.Writer, level int, isLast bool, lastDirLvl int) {
	for i := 0; i < level; i++ {
		if level != 0 && !(i >= level-lastDirLvl) {
			fmt.Fprint(out, "│")
		}
		fmt.Fprint(out, "\t")
	}
	if isLast {
		fmt.Fprint(out, "└───")
	} else {
		fmt.Fprint(out, "├───")
	}
}

func getOnlyDirs(entries []os.DirEntry) []os.DirEntry {
	var onlyDirs []os.DirEntry
	for _, entry := range entries {
		if entry.IsDir() {
			onlyDirs = append(onlyDirs, entry)
		}
	}
	return onlyDirs
}

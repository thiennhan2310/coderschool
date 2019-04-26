package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] [<dir>]\n", os.Args[0])
		flag.PrintDefaults()
	}

	dept := flag.Int("L", 0, "Max display depth of the directory tree.")
	printDirectorOnly := flag.Bool("d", false, "List directories only.")
	printHelp := flag.Bool("h", false, " Prints usage information")
	out := flag.String("o", "", " Send output to filename")
	flag.Parse()
	fmt.Println(dept, *printDirectorOnly, *printHelp, *out)

	if *printHelp == true {
		flag.Usage()
		return
	}

	curDir, _ := os.Getwd()
	var writer io.Writer
	if writer = os.Stdout; *out != "" {
		f, _ := os.Create(*out)
		defer f.Close()
		writer = f
	}
	printTree(writer, curDir, 0, *printDirectorOnly, *dept)
}

func printNested(writer io.Writer, level int) {
	for i := 0; i < level; i++ {
		fmt.Fprintf(writer, "--")
	}
}

func printTree(writer io.Writer, dir string, level int, printDirectorOnly bool, dept int) {
	listFiles, _ := ioutil.ReadDir(dir)
	for _, file := range listFiles {
		if dept != 0 && dept == level {
			return
		}

		if printDirectorOnly {
			if file.IsDir() {
				printNested(writer, level)
				fmt.Fprintf(writer, "%v\n", file.Name())
			}
		} else {
			printNested(writer, level)
			fmt.Fprintf(writer, "%v\n", file.Name())
		}

		if file.IsDir() {
			printTree(writer, file.Name(), level+1, printDirectorOnly, dept)
		}
	}
}

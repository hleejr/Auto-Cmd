package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Cmd struct {
	// path of a program or executable
	Path string
	// arguments to invoke program or executable with
	Args []string
	// environment variables for the execution
	Env []string
	// current working directory of the execution
	Dir string
	// standard input for the execution
	Stdin io.Reader
	// standard output for the execution
	Stdout io.Writer
	// standard error for the execution
	Stderr io.Writer
	// Process is the underlying process, once started
	Process *os.Process
}

func readFile(name string) string {
	fileContents, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

func runCommands(filename string, message string, branch string) {

	result := readFile(filename)
	cmds := strings.Split(result, "\n")

	for _, cmd := range cmds {
		cmdArray := strings.Fields(cmd)
		gitExec, _ := exec.LookPath(cmdArray[0])

		command := &exec.Cmd{
			Path:   gitExec,
			Args:   cmdArray,
			Stdout: os.Stdout,
			Stderr: os.Stdout,
		}

		if strings.Contains(cmd, "commit") {
			command.Args = append(command.Args, message)
		}
		if strings.Contains(cmd, "push") {
			command.Args = append(command.Args, branch)
		}

		err := command.Run()
		if err != nil {
			panic(err)
		}

		fmt.Println("ran command", strconv.Quote(strings.Join(command.Args, " ")))
	}

}

func main() {
	inputFile := flag.String("file", "", "what file has the commands you want to run?")
	inputMsg := flag.String("msg", "commit made", "do you have a message for your commit?")
	inputBranch := flag.String("branch", "master", "what branch are you committing to?")
	flag.Parse()

	runCommands(*inputFile, *inputMsg, *inputBranch)
}

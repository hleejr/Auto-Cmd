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

func runCommands(filename string) {
	inputMsg := flag.String("msg", "commit made", "do you have a message for your commit?")
	inputBranch := flag.String("branch", "master", "what branch are you committing to?")
	flag.Parse()

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
			command.Args = append(command.Args, *inputMsg)
		}
		if strings.Contains(cmd, "push") {
			command.Args = append(command.Args, *inputBranch)
		}

		err := command.Run()
		if err != nil {
			panic(err)
		}

		fmt.Println("ran command %s", strconv.Quote(cmd))
	}

}

func main() {

	runCommands("gitCommands.txt")
}

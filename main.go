package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/shlex"
)

func main() {
	command := flag.String("command", "go test ./...", "The command (with arguments) to run when a .go file is saved.")
	flag.Parse()

	working, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	args, err := shlex.Split(*command)
	if err != nil {
		log.Fatal(err)
	}
	if len(args) < 1 {
		log.Fatal("Please provide something to run.")
	}

	scanner := &Scanner{working: working}
	runner := &Runner{working: working, command: args}
	ticker := time.NewTicker(time.Millisecond * 250)
	defer ticker.Stop()
	for range ticker.C {
		if scanner.Scan() {
			runner.Run()
		}
	}
}

////////////////////////////////////////////////////////////////////////////

type Scanner struct {
	state   int64
	working string
}

func (s *Scanner) Scan() bool {
	newState := s.checksum()
	defer func() { s.state = newState }()
	return newState != s.state
}

func (s *Scanner) checksum() int64 {
	var sum int64 = 0
	err := filepath.Walk(s.working, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			sum++
		} else if strings.HasSuffix(info.Name(), ".go") {
			sum += info.Size() + info.ModTime().Unix()
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return sum
}

////////////////////////////////////////////////////////////////////////////

type Runner struct {
	command []string
	working string
}

func (r *Runner) Run() {
	write(clearScreen)
	output, success := r.run()
	if success {
		write(greenColor)
	} else {
		write(redColor)
	}
	write(string(output))
	write(resetColor)
}

func write(a ...interface{}) {
	fmt.Fprint(os.Stdout, a...)
	os.Stdout.Sync()
}

func (r *Runner) run() (output []byte, success bool) {
	command := exec.Command(r.command[0])
	if len(r.command) > 1 {
		command.Args = append(command.Args, r.command[1:]...)
	}
	command.Dir = r.working

	defer func(t time.Time) { fmt.Println(time.Since(t)) }(time.Now())

	output, err := command.CombinedOutput()
	if err != nil {
		output = append(output, []byte(err.Error())...)
	}
	return output, command.ProcessState.Success()
}

////////////////////////////////////////////////////////////////////////////

var (
	clearScreen = "\033[2J\033[H" // clear the screen and put the cursor at top-left
	greenColor  = "\033[32m"
	redColor    = "\033[31m"
	resetColor  = "\033[0m"
)

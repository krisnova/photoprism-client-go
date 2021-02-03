package sampleapp

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/kris-nova/logger"
)

type Script struct {
	commands []string
}

func NewScript(str string) *Script {
	script := &Script{}
	spl := strings.Split(str, "\n")
	//logger.Info("Script lines: %d", len(spl))
	for _, line := range spl {
		script.commands = append(script.commands, line)
	}
	return script
}

func (s *Script) Interpret() error {
	//logger.Info("Running script...")
	chResult := make(chan *ExecResult)
	chError := make(chan error)
	chBreak := make(chan bool)
	defer close(chResult)
	defer close(chError)
	defer close(chBreak)
	for i, cmdStr := range s.commands {
		// Exec will hang for output

		// Ignore newlines
		// Ignore comments starting with #
		// Ignore comments starting with //
		if cmdStr == "\n" || strings.HasPrefix(cmdStr, "#") || strings.HasPrefix(cmdStr, "//") {
			continue
		}
		//logger.Info("Executing: [%s]", cmdStr)
		result, err := Exec(cmdStr)
		if err != nil {
			return fmt.Errorf("error executing  running command [%s] on line [%d]\n%v\n", cmdStr, i+1, err)
		} else if result.exitCode != 0 {
			return fmt.Errorf("non zero exit code running command [%s] on line [%d]\n%s\n%s\n", cmdStr, i+1, result.Stdout(), result.stderr)
		}
		// Here is where we log STDOUT from a "script"
		// Right now it is set to DEBUG which can be enabled by
		// setting logger.Level = 4
		logger.Debug(result.Stdout())
	}
	return nil
}

type ExecResult struct {
	stderr   string
	stdout   string
	exitCode int
	execErr  exec.ExitError
}

func (e *ExecResult) Stdout() string {
	return e.stdout
}

func (e *ExecResult) Stderr() string {
	return e.stderr
}

func (e *ExecResult) ExitCode() int {
	if e == nil {
		return 0
	}
	return e.exitCode
}

func (e *ExecResult) ExecError() exec.ExitError {
	return e.execErr
}

// Exec will take an arbitrary executable string
// and hang until the command exits
func Exec(str string) (*ExecResult, error) {
	//logger.Info("Exec [%s]", str)
	var cmdstr string
	var args []string
	var l int
	spl := strings.Split(str, " ")
	l = len(spl)
	if l == 1 {
		// <cmd>
		cmdstr = spl[0]
	} else if l > 1 {
		// <cmd> <arg>...
		cmdstr = spl[0]
		for i := 1; i < l; i++ {
			args = append(args, spl[i])
		}
	} else if l < 1 {
		return nil, fmt.Errorf("invalid Exec() string %s", str)
	}
	fqpcmd, err := exec.LookPath(cmdstr)
	if err != nil {
		return nil, fmt.Errorf("unable to find fully qualified path for executable %s: %v", cmdstr, err)
	}

	//logger.Info("Command: %s", fqpcmd)
	//logger.Info("Args: %v", args)

	stdoutBuffer := bytes.Buffer{}
	stderrBuffer := bytes.Buffer{}
	e := []string{fqpcmd, fmt.Sprint(args)}
	cmd := exec.Command(e[0], e[1:]...)
	cmd.Stdout = &stdoutBuffer
	cmd.Stderr = &stderrBuffer
	result := &ExecResult{}
	err = cmd.Run()
	if err != nil {
		if eerr, ok := err.(*exec.ExitError); ok {
			result.stderr = stderrBuffer.String()
			result.stdout = stdoutBuffer.String()
			result.exitCode = eerr.ExitCode()
			result.execErr = *eerr
			return result, nil
		}
		return nil, fmt.Errorf("major error running command [%s]: %v", str, err)
	}
	result.stderr = stderrBuffer.String()
	result.stdout = stdoutBuffer.String()
	result.exitCode = 0
	return result, nil
}

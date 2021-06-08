//  Copyright © 2021 Kris Nóva <kris@nivenly.com>
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.
package sampleapp

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/kris-nova/logger"
)

// Script is a set of commands delimited by newlines
// Comments # and // are ignored.
type Script struct {
	commands []string
}

// NewScriptFromPath is used to build an executable script from a path of disk.
func NewScriptFromPath(path string) (*Script, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("unable to calculate fully qualified path for path: %s: %v", path, err)
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read from path %s: %v", path, err)
	}
	content := string(bytes)
	return NewScriptFromString(content), nil
}

const (
	// IgnoreSpacesBash is the amount of (spaces - 1) that we see in common \\n delimited commands
	IgnoreSpacesBash string = "    "

	// IgnoreTabs is the tab character
	IgnoreTabs string = "\t"
)

// NewScriptFromString is used to build an executable script from the content in string form.
func NewScriptFromString(str string) *Script {
	script := &Script{}
	removeRuleF := func(str string, rs []string) string {
		for _, r := range rs {
			str = strings.Replace(str, r, "", -1)
		}
		return str
	}
	str = strings.Replace(str, "\\\n", "", -1)
	str = removeRuleF(str, []string{IgnoreSpacesBash, IgnoreTabs})
	spl := strings.Split(str, "\n")
	//logger.Info("Script lines: %d", len(spl))
	for _, line := range spl {
		script.commands = append(script.commands, line)
	}
	return script
}

// Interpret is used to procedurally execute a script. The script will execute each line independently
// and can error at any point in the executation path.
func (s *Script) Interpret() error {
	//logger.Info("Running script...")
	for i, cmdStr := range s.commands {
		// Exec will hang for output
		// Ignore newlines
		// Ignore comments starting with #
		// Ignore comments starting with //
		if cmdStr == "\n" || cmdStr == "" || strings.HasPrefix(cmdStr, "#") || strings.HasPrefix(cmdStr, "//") {
			continue
		}
		logger.Info("Executing: [%s]", cmdStr)
		result, err := Exec(cmdStr)
		if err != nil {
			return fmt.Errorf("error executing  running command [%s] on line [%d]\n%v\n", cmdStr, i+1, err)
		} else if result.exitCode != 0 {
			return fmt.Errorf("non zero exit code running command [%s] on line [%d]\n%s\n%s\n", cmdStr, i+1, result.Stdout(), result.Stderr())
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
		logger.Debug("unable to find fully qualified path for executable %s: %v", cmdstr, err)
	}

	//logger.Info("Command: %s", fqpcmd)
	//logger.Info("Args: %v", args)

	stdoutBuffer := bytes.Buffer{}
	stderrBuffer := bytes.Buffer{}
	e := []string{fqpcmd}
	for _, arg := range args {
		e = append(e, arg)
	}
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

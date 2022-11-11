// Copyright © 2021 Kris Nóva <kris@nivenly.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package sampleapp

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/kris-nova/logger"
)

const (
	// IgnoreSpacesBash is the amount of (spaces - 1) that we see in common \\n delimited commands
	IgnoreSpacesBash string = "    "

	// IgnoreTabs is the tab character
	IgnoreTabs string = "\t"
)

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

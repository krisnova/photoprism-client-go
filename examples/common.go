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
package main

import (
	"fmt"
	"os"

	photoprism "github.com/drummonds/photoprism-client-go"

	"github.com/kris-nova/logger"
)

// format string, a ...interface{}
func halt(code int, msg string, a ...interface{}) {
	str := fmt.Sprintf(msg, a...)
	logger.Critical(str)
	os.Exit(code)
}

func auth() photoprism.ClientAuthenticator {
	user := os.Getenv("PHOTOPRISM_USER")
	if user == "" {
		halt(1, "Missing PHOTOPRISM_USER")
	}
	pass := os.Getenv("PHOTOPRISM_PASS")
	if pass == "" {
		halt(2, "Missing PHOTOPRISM_PASS")
	}
	auth := photoprism.NewClientAuthLogin(user, pass)
	return auth
}

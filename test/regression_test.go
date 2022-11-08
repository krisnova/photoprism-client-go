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
package test

import (
	"testing"

	"github.com/drummonds/photoprism-client-go"
)

// Trailing slash issue
// https://github.com/drummonds/photoprism-client-go/issues/2
func TestRegressionIssue2(t *testing.T) {
	testStrings := []string{"localhost/", "localhost///////", "localhost//"}
	goal := "localhost"
	for _, str := range testStrings {
		client := photoprism.New(str)
		if client.ConnectionString() != goal {
			t.Error("Failed to trim suffix / in client connection string")
			t.FailNow()
		}
	}

}

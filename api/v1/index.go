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
package api

// Index is used to sync the backend storage with
// the database meta information
//
// POST /api/v1/index
func (v1 *V1Client) Index() error {
	resp := v1.POST(nil, "/api/v1/index")
	return resp.Error
}

// CancelIndex can be used to attempt to cancel a running index
// operation
//
// DELETE /api/v1/index
func (v1 *V1Client) CancelIndex() error {
	resp := v1.DELETE(nil, "/api/v1/index")
	return resp.Error
}

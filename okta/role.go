/*
* Copyright 2018 - Present Okta, Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

// AUTO-GENERATED!  DO NOT EDIT FILE DIRECTLY

package okta

import (
	"time"
)

type Role struct {
	Embedded    interface{} `json:"_embedded,omitempty"`
	Created     *time.Time  `json:"created,omitempty"`
	Description string      `json:"description,omitempty"`
	Id          string      `json:"id,omitempty"`
	Label       string      `json:"label,omitempty"`
	LastUpdated *time.Time  `json:"lastUpdated,omitempty"`
	Status      string      `json:"status,omitempty"`
	Type        string      `json:"type,omitempty"`
}

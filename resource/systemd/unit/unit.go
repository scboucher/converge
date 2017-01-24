// Copyright © 2016 Asteris, LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package unit

import "fmt"

type Unit struct {
	Name           string
	Description    string
	ActiveState    string
	Path           string
	Type           UnitType
	Properties     map[string]interface{}
	TypeProperties map[string]interface{}
}

func (u *Unit) IsServiceUnit() bool {
	return UnitTypeService == UnitTypeFromName(u.Path)
}

func PPUnit(u *Unit) string {
	knownTypeProps := make(map[string]string)
	for p, v := range u.TypeProperties {
		knownTypeProps[p] = fmt.Sprintf("%T", v)
	}
	fmtStr := `
Unit
=================
Name:        %s
Description: %s
ActiveState: %s
Path:        %s
-----------------
Properties:
%v
-----------------
Typed Properties:
%v
=================
`
	return fmt.Sprintf(fmtStr, u.Name, u.Description, u.ActiveState, u.Path, u.Properties, knownTypeProps)
}

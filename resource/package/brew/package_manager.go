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

package brew

import (
	"fmt"

	"github.com/asteris-llc/converge/resource/package"
)

// Outputs from Brew
const (
	// NotInstalled means the package is not installed
	NotInstalled = "Not Installed"
)

type BrewManager struct {
	Sys pkg.SysCaller
}

func (b *BrewManager) InstalledVersion(p string) (pkg.PackageVersion, bool) {
	result, err := b.Sys.Run(fmt.Sprintf("brew info %s", p))
}

func (b *BrewManager) RemovePackage(pkg string) (string, error) {
	res, err := b.Sys.Run(fmt.Sprintf("brew uninstall --force %s", pkg))
	return string(res), err
}

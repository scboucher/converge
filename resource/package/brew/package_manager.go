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
	"strings"

	"github.com/asteris-llc/converge/resource/package"
)

// Outputs from Brew
const (
	// NotInstalled means the package is not installed
	NotInstalled = "Not Installed"
	BrewConfig   = "brew config"
	BrewCellar   = "HOMEBREW_CELLAR"
)

// BrewManager is a concrete implemantation of PackageManager for homebrew
type BrewManager struct {
	Sys pkg.SysCaller
}

func (b *BrewManager) InstalledVersion(p string) (pkg.PackageVersion, bool) {
	var version string
	var installed bool
	var cellar string
	result, err := b.Sys.Run(fmt.Sprintf("brew info %s", p))
	line3 := strings.Split(strings.TrimSpace(string(result)), "\n")[3]
	if strings.Compare(strings.TrimSpace(line3), NotInstalled) == 0 {
		return "", false
	}
	config, err := b.Sys.Run(fmt.Sprintf(BrewConfig))
	brewConfigLines := strings.Split(strings.TrimSpace(string(config)), "\n")
	for _, line := range brewConfigLines {
		if strings.Contains(line, BrewCellar) {
			cellar = strings.TrimSpace(strings.Split(line, ": ")[1])
			version = strings.Split(strings.Split(line3, " ")[0], "%s/%s", cellar, p)
			installed = true

		}
	}

	return version, installed

}
func (b *BrewManager) InstallPackage(p string) (string, error) {
	if _, isInstalled := b.InstalledVersion(p); isInstalled {
		return "already installed", nil
	}
	res, err := b.Sys.Run(fmt.Sprintf("brew install %s", p))
	return string(res), err
}

func (b *BrewManager) RemovePackage(p string) (string, error) {
	switch _, isInstalled := b.InstalledVersion(p); isInstalled {
	case true:
		res, err := b.Sys.Run(fmt.Sprintf("brew uninstall --force %s", p))
		return string(res), err
	default:
		return "package is not installed", nil
	}
}

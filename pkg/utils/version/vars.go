/*
Copyright 2023 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package version

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"sigs.k8s.io/kwok/pkg/consts"
)

// DisplayVersion is the version string for the current build
func DisplayVersion() string {
	return fmt.Sprintf("%s %s (%s/%s)", AddPrefixV(consts.Version), runtime.Version(), runtime.GOOS, runtime.GOARCH)
}

// DefaultUserAgent returns a User-Agent string built from static global vars.
func DefaultUserAgent() string {
	return fmt.Sprintf("%s/%s (%s/%s)", adjustCommand(os.Args[0]), AddPrefixV(consts.Version), runtime.GOOS, runtime.GOARCH)
}

// adjustCommand returns the last component of the
// OS-specific command path for use in User-Agent.
func adjustCommand(p string) string {
	// Unlikely, but better than returning "".
	if len(p) == 0 {
		return "unknown"
	}
	return filepath.Base(p)
}

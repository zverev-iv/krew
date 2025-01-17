// Copyright 2019 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package constants

import "os"

const (
	CurrentAPIVersion = "krew.googlecontainertools.github.com/v1alpha2"
	PluginKind        = "Plugin"
	ManifestExtension = ".yaml"
	KrewPluginName    = "krew" // plugin name of krew itself

	// DefaultIndexName is a magic string that's used for a plugin name specified without an index.
	DefaultIndexName = "default"

	// defaultIndexURI is the default krew-index URI that gets used unless it's overridden
	defaultIndexURI = "https://github.com/kubernetes-sigs/krew-index.git"
)

// DefaultIndexURI points to the upstream index.
var DefaultIndexURI = defaultIndexURI

func init() {
	if uri := os.Getenv("KREW_DEFAULT_INDEX_URI"); uri != "" {
		DefaultIndexURI = uri
	}
}

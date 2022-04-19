// Copyright  observIQ, Inc.
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

package pluginreceiver

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const pluginDirPath = "../../plugins"

// TestValidateSuppliedPlugins ensures each plugin that ships with the collector loads with the current
// version of the receiver
func TestValidateSuppliedPlugins(t *testing.T) {
	entries, err := os.ReadDir(pluginDirPath)
	require.NoError(t, err)

	for _, entry := range entries {
		t.Run(fmt.Sprintf("Loading %s", entry.Name()), func(t *testing.T) {
			t.Parallel()
			fullFilePath, err := filepath.Abs(filepath.Join(pluginDirPath, entry.Name()))
			assert.NoError(t, err, "Failed to determine path of file %s", entry.Name())

			// Load the plugin
			plugin, err := LoadPlugin(fullFilePath)
			assert.NoError(t, err, "Failed to load file %s", entry.Name())

			_, err = plugin.RenderComponents(map[string]interface{}{})
			assert.NoError(t, err, "Failed to render components for plugin %s", entry.Name())
		})
	}
}
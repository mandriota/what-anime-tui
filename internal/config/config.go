// Copyright 2023 Mark Mandriota
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
package config

import "github.com/charmbracelet/lipgloss"

type BorderConfig struct {
	Foreground lipgloss.Color
}

type AppearanceConfig struct {
	Background lipgloss.Color
	Foreground lipgloss.Color
	Border     BorderConfig
}

// Warning: this fetcher config is not stable and may be changed in future releases.
type FetcherConfig struct {
	ApiUrlByUrl  string
	ApiUrlByFile string
}

type GeneralConfig struct {
	Appearance AppearanceConfig
	Fetcher FetcherConfig
}


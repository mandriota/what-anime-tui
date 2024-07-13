// Copyright 2023 Mark Mandriota
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
package fetcher

import (
	"fmt"
	"math"
)

func formatSecondsToDuration(s float64) string {
	return fmt.Sprintf("%02d:%02d:%02.1f",
		int(s/3600),
		int(s/60)%60,
		math.Mod(s, 60),
	)
}

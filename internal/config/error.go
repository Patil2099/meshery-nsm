// Copyright 2020 Layer5, Inc.
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

import (
	"fmt"

	"github.com/layer5io/meshkit/errors"
)

const (
	// ErrEmptyConfigCode represents the error when the configuration is either empty
	// or is invalid
	ErrEmptyConfigCode = "11300"

	// ErrGetLatestReleasesCode represents the error which occurs during the process of getting
	// latest releases
	ErrGetLatestReleasesCode = "nsm_mesh_test_code"

	// ErrGetLatestReleaseNamesCode represents the error which occurs during the process of extracting
	// release names
	ErrGetLatestReleaseNamesCode = "nsm_mesh_test_code"
)

var (
	// ErrEmptyConfig error is the error when config is invalid
	ErrEmptyConfig = errors.NewDefault(ErrEmptyConfigCode, "Config is empty")
)

// ErrGetLatestReleases is the error for fetching nsm-mesh releases
func ErrGetLatestReleases(err error) error {
	return errors.NewDefault(ErrGetLatestReleasesCode, fmt.Sprintf("unable to fetch release info: %s", err.Error()))
}

// ErrGetLatestReleaseNames is the error for fetching nsm-mesh releases
func ErrGetLatestReleaseNames(err error) error {
	return errors.NewDefault(ErrGetLatestReleaseNamesCode, fmt.Sprintf("failed to extract release names: %s", err.Error()))
}

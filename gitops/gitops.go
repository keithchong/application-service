//
// Copyright 2021-2022 Red Hat, Inc.
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

package gitops

import (
	"fmt"
	"path/filepath"

	appstudiov1alpha1 "github.com/redhat-appstudio/application-service/api/v1alpha1"
	"github.com/redhat-appstudio/application-service/gitops/prepare"
	"github.com/spf13/afero"
)

func GenerateTektonBuild(outputPath string, component appstudiov1alpha1.Component, appFs afero.Afero, context string, gitopsConfig prepare.GitopsConfig) error {
	componentName := component.Name
	repoPath := filepath.Join(outputPath, componentName)
	gitopsFolder := filepath.Join(repoPath, context)
	componentPath := filepath.Join(gitopsFolder, "components", componentName, "base")
	if err := GenTektonBuild(appFs, componentPath, component, gitopsConfig); err != nil {
		return fmt.Errorf("failed to generate tekton build in %q for component %q: %s", componentPath, componentName, err)
	}
	return nil
}

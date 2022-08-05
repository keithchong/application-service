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
	"path/filepath"

	appstudiov1alpha1 "github.com/redhat-appstudio/application-service/api/v1alpha1"
	"github.com/redhat-appstudio/application-service/gitops/prepare"
	gitopsgen "github.com/redhat-developer/gitops-generator/pkg"
	"github.com/spf13/afero"
)

const (
	kustomizeFileName       = "kustomization.yaml"
	deploymentFileName      = "deployment.yaml"
	deploymentPatchFileName = "deployment-patch.yaml"
	serviceFileName         = "service.yaml"
	routeFileName           = "route.yaml"
	repositoryFileName      = "repository.yaml"
)

func GenTektonBuild(fs afero.Afero, outputFolder string, component appstudiov1alpha1.Component, gitopsConfig prepare.GitopsConfig) error {

	if component.Spec.Source.GitSource != nil && component.Spec.Source.GitSource.URL != "" {
		tektonResourcesDirName := ".tekton"

		if err := GenerateBuild(fs, filepath.Join(outputFolder, tektonResourcesDirName), component, gitopsConfig); err != nil {
			return err
		}
		// Update the kustomize file and return
		return gitopsgen.UpdateExistingKustomize(fs, outputFolder)
	}
	return nil
}

/*
Copyright 2017 The Nuclio Authors.

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

package factory

import (
	"github.com/nuclio/nuclio/pkg/common"
	"github.com/nuclio/nuclio/pkg/containerimagebuilderpusher"
	"github.com/nuclio/nuclio/pkg/platform"
	"github.com/nuclio/nuclio/pkg/platform/kube"
	"github.com/nuclio/nuclio/pkg/platform/local"
	"github.com/nuclio/nuclio/pkg/platformconfig"

	"github.com/nuclio/errors"
	"github.com/nuclio/logger"
)

// CreatePlatform creates a platform based on a requested type (platformType) and configuration it receives
// and probes
func CreatePlatform(parentLogger logger.Logger,
	platformType string,
	platformConfiguration *platformconfig.Config,
	defaultNamespace string) (platform.Platform, error) {

	var newPlatform platform.Platform
	var err error

	platformConfiguration.ContainerBuilderConfiguration = containerimagebuilderpusher.NewContainerBuilderConfiguration()

	switch platformType {
	case "local":
		newPlatform, err = local.NewPlatform(parentLogger, platformConfiguration)

	case "kube":
		newPlatform, err = kube.NewPlatform(parentLogger, platformConfiguration)

	case "auto":

		// kubeconfig path is set, or running in kubernetes clsuter
		if common.GetKubeconfigPath(platformConfiguration.Kube.KubeConfigPath) != "" ||
			kube.IsInCluster() {

			// call again, but force kube
			newPlatform, err = CreatePlatform(parentLogger, "kube", platformConfiguration, defaultNamespace)
		} else {

			// call again, force local
			newPlatform, err = CreatePlatform(parentLogger, "local", platformConfiguration, defaultNamespace)
		}

	default:
		return nil, errors.Errorf("Can't create platform - unsupported: %s", platformType)
	}

	if err != nil {
		return nil, errors.Wrapf(err, "Failed to create %s platform", platformType)
	}

	if err = EnsureDefaultProjectExistence(parentLogger, newPlatform, defaultNamespace); err != nil {
		return nil, errors.Wrap(err, "Failed to ensure default project existence")
	}

	return newPlatform, nil
}

func EnsureDefaultProjectExistence(parentLogger logger.Logger, p platform.Platform, defaultNamespace string) error {
	resolvedNamespace := p.ResolveDefaultNamespace(defaultNamespace)

	projects, err := p.GetProjects(&platform.GetProjectsOptions{
		Meta: platform.ProjectMeta{
			Name:      platform.DefaultProjectName,
			Namespace: resolvedNamespace,
		},
	})
	if err != nil {
		return errors.Wrap(err, "Failed to get projects")
	}

	if len(projects) == 0 {

		// if we're here the default project doesn't exist. create it
		projectConfig := platform.ProjectConfig{
			Meta: platform.ProjectMeta{
				Name:      platform.DefaultProjectName,
				Namespace: resolvedNamespace,
			},
			Spec: platform.ProjectSpec{},
		}
		newProject, err := platform.NewAbstractProject(parentLogger, p, projectConfig)
		if err != nil {
			return errors.Wrap(err, "Failed to create abstract default project")
		}

		if err := p.CreateProject(&platform.CreateProjectOptions{
			ProjectConfig: newProject.GetConfig(),
		}); err != nil {
			return errors.Wrap(err, "Failed to create default project")
		}

	} else if len(projects) > 1 {
		return errors.New("Something went wrong. There's more than one default project")
	}

	return nil
}

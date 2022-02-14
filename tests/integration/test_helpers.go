// Copyright 2021 VMware
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

package integration

import (
	"io"
	"os"
	"path/filepath"

	"sigs.k8s.io/controller-runtime/pkg/envtest"
)

const DebugControlPlane = false

// GetWaitTimeout returns the time to wait for a call to the APIServer
// If in CI, we should give the API server a lot of time to respond.
// When testing, we keep it short for faster iterations, however
// you can use CI=true when running your tests if flakes are causing you
// grief.
func GetWaitTimeout() string {
	if os.Getenv("CI") == "true" {
		return "10s"
	}
	return "1s"
}

func CreateTestEnv(workingDir string, out io.Writer) *envtest.Environment {
	testEnv := &envtest.Environment{
		WebhookInstallOptions: envtest.WebhookInstallOptions{
			Paths: []string{filepath.Join("..", "..", "..", "config", "webhook")},
		},
		CRDDirectoryPaths: []string{
			filepath.Join(workingDir, "..", "..", "..", "config", "crd", "bases"),
			filepath.Join(workingDir, "..", "..", "resources", "crds"),
		},
		AttachControlPlaneOutput: DebugControlPlane, // Set to true for great debug logging
	}

	if DebugControlPlane {
		apiServer := testEnv.ControlPlane.GetAPIServer()
		apiServer.Out = out
		apiServer.Err = out
		apiServer.Configure().
			Append("audit-policy-file", filepath.Join(workingDir, "..", "policy.yaml")).
			Append("audit-log-path", "-")
	}

	return testEnv
}

/*
 * Copyright 2018 SUSE LINUX GmbH, Nuernberg, Germany..
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package config

const (
	// DefaultClusterAdminRole admin role
	DefaultClusterAdminRole = "cluster-admin"
)

const (
	// DefaultConfigMapFilename Configmap (in the conatiner)
	DefaultConfigMapFilename = "/etc/dex/cfg/config.yaml"

	// DefaultNodePort Default Dex port
	DefaultNodePort = 32000

	// DefaultCertsDir the directory where certs are stored (in the container)
	DefaultCertsDir = "/etc/dex/tls"

	// DefaultSharedPasswordLen the length (in bytes) for random passwords
	DefaultSharedPasswordLen = 16
)

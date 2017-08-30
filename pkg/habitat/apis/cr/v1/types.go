// Copyright (c) 2017 Chef Software Inc. and/or applicable contributors
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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	HabitatResourcePlural = "habitats"
	HabitatLabel          = "habitat"

	TopologyLabel = "topology"
)

type Habitat struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              HabitatSpec   `json:"spec"`
	Status            HabitatStatus `json:"status,omitempty"`
}

type HabitatSpec struct {
	// Count is the amount of Services to start in this Service Group.
	Count int `json:"count"`
	// Image is the Docker image of the Habitat Service.
	Image   string  `json:"image"`
	Service Service `json:"service"`
}

type HabitatStatus struct {
	State   HabitatState `json:"state,omitempty"`
	Message string       `json:"message,omitempty"`
}

type HabitatState string

type Service struct {
	// Group is the value of the --group flag for the hab client.
	// Optional. Defaults to `default`.
	Group string `json:"group"`
	// Topology is the value of the --topology flag for the hab client.
	Topology `json:"topology"`
	// Config is the name of the Secret that the user has previously created.
	// The file with this name is mounted inside of the pod. Habitat will
	// use it for initial configuration of the service.
	Config string `json:"config"`
	// The name of the secret that contains the ring key.
	// Optional.
	RingSecretName string `json:"ringSecretName,omitempty"`
}

type Topology string

func (t Topology) String() string {
	return string(t)
}

const (
	HabitatStateCreated   HabitatState = "Created"
	HabitatStateProcessed HabitatState = "Processed"

	TopologyStandalone Topology = "standalone"
	TopologyLeader     Topology = "leader"
)

type HabitatList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Habitat `json:"items"`
}

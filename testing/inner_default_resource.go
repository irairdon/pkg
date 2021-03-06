/*
Copyright 2019 The Knative Authors

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

package testing

import (
	"context"

	"/apis"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// InnerDefaultResource is a simple resource that's compatible with our webhook. It differs from
// Resource by not omitting empty `spec`, so can change when it round trips
// JSON -> Golang type -> JSON.
type InnerDefaultResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Note that this does _not_ have omitempty. So when JSON is round tripped through the Golang
	// type, `spec: {}` will automatically be injected.
	Spec InnerDefaultSpec `json:"spec"`
}

// InnerDefaultSpec is the spec for InnerDefaultResource.
type InnerDefaultSpec struct {
	Generation int64 `json:"generation,omitempty"`

	FieldWithDefault string `json:"fieldWithDefault,omitempty"`
}

// Check that ImmutableDefaultResource may be validated and defaulted.
var _ apis.Validatable = (*InnerDefaultResource)(nil)
var _ apis.Defaultable = (*InnerDefaultResource)(nil)

// SetDefaults sets default values.
func (i *InnerDefaultResource) SetDefaults(ctx context.Context) {
	i.Spec.SetDefaults(ctx)
}

// SetDefaults sets default values.
func (cs *InnerDefaultSpec) SetDefaults(ctx context.Context) {
	if cs.FieldWithDefault == "" {
		cs.FieldWithDefault = "I'm a default."
	}
}

// Validate validates the resource.
func (*InnerDefaultResource) Validate(ctx context.Context) *apis.FieldError {
	return nil
}

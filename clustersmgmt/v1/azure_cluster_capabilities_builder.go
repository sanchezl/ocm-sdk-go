/*
Copyright (c) 2020 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

// AzureClusterCapabilitiesBuilder contains the data and logic needed to build 'azure_cluster_capabilities' objects.
type AzureClusterCapabilitiesBuilder struct {
	bitmap_  uint32
	disabled []AzureOptionalClusterCapability
}

// NewAzureClusterCapabilities creates a new builder of 'azure_cluster_capabilities' objects.
func NewAzureClusterCapabilities() *AzureClusterCapabilitiesBuilder {
	return &AzureClusterCapabilitiesBuilder{}
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *AzureClusterCapabilitiesBuilder) Empty() bool {
	return b == nil || b.bitmap_ == 0
}

// Disabled sets the value of the 'disabled' attribute to the given values.
func (b *AzureClusterCapabilitiesBuilder) Disabled(values ...AzureOptionalClusterCapability) *AzureClusterCapabilitiesBuilder {
	b.disabled = make([]AzureOptionalClusterCapability, len(values))
	copy(b.disabled, values)
	b.bitmap_ |= 1
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *AzureClusterCapabilitiesBuilder) Copy(object *AzureClusterCapabilities) *AzureClusterCapabilitiesBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	if object.disabled != nil {
		b.disabled = make([]AzureOptionalClusterCapability, len(object.disabled))
		copy(b.disabled, object.disabled)
	} else {
		b.disabled = nil
	}
	return b
}

// Build creates a 'azure_cluster_capabilities' object using the configuration stored in the builder.
func (b *AzureClusterCapabilitiesBuilder) Build() (object *AzureClusterCapabilities, err error) {
	object = new(AzureClusterCapabilities)
	object.bitmap_ = b.bitmap_
	if b.disabled != nil {
		object.disabled = make([]AzureOptionalClusterCapability, len(b.disabled))
		copy(object.disabled, b.disabled)
	}
	return
}

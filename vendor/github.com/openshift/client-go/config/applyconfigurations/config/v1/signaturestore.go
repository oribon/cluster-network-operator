// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// SignatureStoreApplyConfiguration represents an declarative configuration of the SignatureStore type for use
// with apply.
type SignatureStoreApplyConfiguration struct {
	URL *string                                   `json:"url,omitempty"`
	CA  *ConfigMapNameReferenceApplyConfiguration `json:"ca,omitempty"`
}

// SignatureStoreApplyConfiguration constructs an declarative configuration of the SignatureStore type for use with
// apply.
func SignatureStore() *SignatureStoreApplyConfiguration {
	return &SignatureStoreApplyConfiguration{}
}

// WithURL sets the URL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the URL field is set to the value of the last call.
func (b *SignatureStoreApplyConfiguration) WithURL(value string) *SignatureStoreApplyConfiguration {
	b.URL = &value
	return b
}

// WithCA sets the CA field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CA field is set to the value of the last call.
func (b *SignatureStoreApplyConfiguration) WithCA(value *ConfigMapNameReferenceApplyConfiguration) *SignatureStoreApplyConfiguration {
	b.CA = value
	return b
}
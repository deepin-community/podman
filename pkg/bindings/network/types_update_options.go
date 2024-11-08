// Code generated by go generate; DO NOT EDIT.
package network

import (
	"net/url"

	"github.com/containers/podman/v5/pkg/bindings/internal/util"
)

// Changed returns true if named field has been set
func (o *UpdateOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams formats struct fields to be passed to API service
func (o *UpdateOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}

// WithAddDNSServers set field AddDNSServers to given value
func (o *UpdateOptions) WithAddDNSServers(value []string) *UpdateOptions {
	o.AddDNSServers = value
	return o
}

// GetAddDNSServers returns value of field AddDNSServers
func (o *UpdateOptions) GetAddDNSServers() []string {
	if o.AddDNSServers == nil {
		var z []string
		return z
	}
	return o.AddDNSServers
}

// WithRemoveDNSServers set field RemoveDNSServers to given value
func (o *UpdateOptions) WithRemoveDNSServers(value []string) *UpdateOptions {
	o.RemoveDNSServers = value
	return o
}

// GetRemoveDNSServers returns value of field RemoveDNSServers
func (o *UpdateOptions) GetRemoveDNSServers() []string {
	if o.RemoveDNSServers == nil {
		var z []string
		return z
	}
	return o.RemoveDNSServers
}
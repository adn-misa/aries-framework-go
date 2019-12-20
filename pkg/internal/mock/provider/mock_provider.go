/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package provider

import (
	"github.com/hyperledger/aries-framework-go/pkg/didcomm/common/didconnection"
	"github.com/hyperledger/aries-framework-go/pkg/didcomm/dispatcher"
	"github.com/hyperledger/aries-framework-go/pkg/didcomm/packer"
	vdriapi "github.com/hyperledger/aries-framework-go/pkg/framework/aries/api/vdri"
	"github.com/hyperledger/aries-framework-go/pkg/kms"
	"github.com/hyperledger/aries-framework-go/pkg/storage"
)

// Provider mocks provider needed for did exchange service initialization
type Provider struct {
	ServiceValue                  interface{}
	ServiceErr                    error
	KMSValue                      kms.KeyManager
	InboundEndpointValue          string
	StorageProviderValue          storage.Provider
	TransientStorageProviderValue storage.Provider
	PackerList                    []packer.Packer
	PackerValue                   packer.Packer
	OutboundDispatcherValue       dispatcher.Outbound
	VDRIRegistryValue             vdriapi.Registry
	ConnectionStoreValue          didconnection.Store
}

// Service return service
func (p *Provider) Service(id string) (interface{}, error) {
	return p.ServiceValue, p.ServiceErr
}

// KMS returns a KMS instance
func (p *Provider) KMS() kms.KeyManager {
	return p.KMSValue
}

// InboundTransportEndpoint returns the inbound transport endpoint
func (p *Provider) InboundTransportEndpoint() string {
	return p.InboundEndpointValue
}

// StorageProvider returns the storage provider
func (p *Provider) StorageProvider() storage.Provider {
	return p.StorageProviderValue
}

// TransientStorageProvider returns the transient storage provider
func (p *Provider) TransientStorageProvider() storage.Provider {
	return p.TransientStorageProviderValue
}

// Packers returns the available Packer services
func (p *Provider) Packers() []packer.Packer {
	return p.PackerList
}

// PrimaryPacker returns the main Packer service
func (p *Provider) PrimaryPacker() packer.Packer {
	return p.PackerValue
}

// OutboundDispatcher return outbound dispatcher
func (p *Provider) OutboundDispatcher() dispatcher.Outbound {
	return p.OutboundDispatcherValue
}

// VDRIRegistry return vdri registry
func (p *Provider) VDRIRegistry() vdriapi.Registry {
	return p.VDRIRegistryValue
}

// DIDConnectionStore returns a didconnection.Store service.
func (p *Provider) DIDConnectionStore() didconnection.Store {
	return p.ConnectionStoreValue
}

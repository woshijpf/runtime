// Copyright (c) 2017 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

package vcmock

import (
	vc "github.com/kata-containers/runtime/virtcontainers"
)

// ID implements the VCSandbox function of the same name.
func (p *Sandbox) ID() string {
	return p.MockID
}

// Annotations implements the VCSandbox function of the same name.
func (p *Sandbox) Annotations(key string) (string, error) {
	return p.MockAnnotations[key], nil
}

// SetAnnotations implements the VCSandbox function of the same name.
func (p *Sandbox) SetAnnotations(annotations map[string]string) error {
	return nil
}

// GetAnnotations implements the VCSandbox function of the same name.
func (p *Sandbox) GetAnnotations() map[string]string {
	return p.MockAnnotations
}

// GetAllContainers implements the VCSandbox function of the same name.
func (p *Sandbox) GetAllContainers() []vc.VCContainer {
	var ifa = make([]vc.VCContainer, len(p.MockContainers))

	for i, v := range p.MockContainers {
		ifa[i] = v
	}

	return ifa
}

// GetContainer implements the VCSandbox function of the same name.
func (p *Sandbox) GetContainer(containerID string) vc.VCContainer {
	for _, c := range p.MockContainers {
		if c.MockID == containerID {
			return c
		}
	}
	return &Container{}
}

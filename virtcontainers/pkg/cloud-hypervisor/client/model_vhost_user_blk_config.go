/*
 * Cloud Hypervisor API
 *
 * Local HTTP based API for managing and inspecting a cloud-hypervisor virtual machine.
 *
 * API version: 0.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// VhostUserBlkConfig struct for VhostUserBlkConfig
type VhostUserBlkConfig struct {
	Wce bool `json:"wce"`
	VuCfg VhostUserConfig `json:"vu_cfg"`
}

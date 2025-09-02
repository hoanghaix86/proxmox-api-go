package core

import (
	"github.com/hoanghaix86/proxmox-api-go/proxmox/attributes"
)

type UPID string

type QEMU struct {
	Id   uint64 `json:"id"`
	Node string `json:"node"`

	Agent *attributes.Agent `json:"agent,omitempty"`

	Bios string `json:"bios,omitempty"`

	Boot    string `json:"boot,omitempty"`    // computed
	VmGenId string `json:"vmgenid,omitempty"` // computed
	Meta    string `json:"meta,omitempty"`    // computed
	SmBios1 string `json:"smbios1,omitempty"` // computed
}

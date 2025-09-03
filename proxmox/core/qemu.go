package core

import (
	"github.com/hoanghaix86/proxmox-api-go/proxmox/attributes"
)

type UPID string

type Hardware struct {
	Memory         uint64               `json:"memory"`
	Cpu            string               `json:"cpu"`
	Cores          uint64               `json:"cores"`
	Bios           string               `json:"bios"`
	Vga            string               `json:"vga"`
	Machine        string               `json:"machine"`
	ScsiController string               `json:"scsihw"`
	Ide2           *attributes.Ide      `json:"ide2"`
	Scsi0          string               `json:"scsi0"`
	EfiDisk0       *attributes.EfIdisk  `json:"efidisk0"`
	TpmState0      *attributes.TpmState `json:"tpmstate0"`
	Net0           *attributes.Network  `json:"net0"`
	Serial0        string               `json:"serial0"`
}

type Cloudinit struct {
	CiUser       string `json:"ciuser"`
	CiPassword   string `json:"cipassword"`
	SearchDomain string `json:"searchdomain"`
	NameServer   string `json:"nameserver"`
	SshKeys      string `json:"sshkeys"`
	CiUpgrade    bool   `json:"ciupgrade"`
	CiCustom     string `json:"cicustom"`
}

type Options struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Startup     string            `json:"startup"`
	OsType      string            `json:"ostype"`
	Boot        string            `json:"boot"`
	Agent       *attributes.Agent `json:"agent,omitempty"`
	SmBios1     string            `json:"smbios1,omitempty"` // computed
}

type QEMU struct {
	Id   uint64 `json:"id"`
	Node string `json:"node"`

	Hardware
	Cloudinit
	Options

	VmGenId string `json:"vmgenid,omitempty"` // computed
	Meta    string `json:"meta,omitempty"`    // computed
}

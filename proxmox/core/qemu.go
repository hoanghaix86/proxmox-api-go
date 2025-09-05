package core

import (
	"github.com/hoanghaix86/proxmox-api-go/proxmox/attributes"
)

type UPID string

type Hardware struct {
	Memory         uint64               `json:"memory"`
	Cpu            attributes.CpuType   `json:"cpu"`
	Cores          uint64               `json:"cores"`
	Bios           attributes.BiosType  `json:"bios"`
	Vga            *attributes.Vga      `json:"vga"`
	Machine        attributes.Machine   `json:"machine"`
	ScsiController string               `json:"scsihw"`
	Ide0           *attributes.Ide      `json:"ide0"`
	Ide2           *attributes.Ide      `json:"ide2"`
	Scsi0          *attributes.Scsi     `json:"scsi0"`
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
	Startup     string            `json:"startup"` // Format: [[order=]\d+] [,up=\d+] [,down=\d+]
	OsType      attributes.OsType `json:"ostype"`
	Boot        string            `json:"boot"` // Format: order=<device[;device...]>
	Agent       *attributes.Agent `json:"agent"`
	OnBoot      bool              `json:"onboot"`
	SmBios1     string            `json:"smbios1"`
	VmGenId     string            `json:"vmgenid"`
	Meta        string            `json:"meta"`
	Template    bool              `json:"template"`
}

type QEMU struct {
	Id   uint64 `json:"id"`
	Node string `json:"node"`
	Hardware
	Cloudinit
	Options
}

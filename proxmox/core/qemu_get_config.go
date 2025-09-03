package core

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hoanghaix86/proxmox-api-go/proxmox/attributes"
	"github.com/hoanghaix86/proxmox-api-go/proxmox/client"
)

type GetConfigQEMUResponse struct {
	//Hardware//
	Memory         string `json:"memory,omitempty"`
	Cpu            string `json:"cpu,omitempty"`
	Cores          uint64 `json:"cores,omitempty"`
	Bios           string `json:"bios,omitempty"`
	Vga            string `json:"vga,omitempty"`
	Machine        string `json:"machine,omitempty"`
	ScsiController string `json:"scsihw,omitempty"`
	Ide2           string `json:"ide2,omitempty"`
	Scsi0          string `json:"scsi0,omitempty"`
	EfiDisk0       string `json:"efidisk0,omitempty"`
	TpmState0      string `json:"tpmstate0,omitempty"`
	Net0           string `json:"net0,omitempty"`
	Serial0        string `json:"serial0,omitempty"`
	//Cloudinit//
	CiUser       string `json:"ciuser,omitempty"`
	CiPassword   string `json:"cipassword,omitempty"`
	SearchDomain string `json:"searchdomain,omitempty"`
	NameServer   string `json:"nameserver,omitempty"`
	SshKeys      string `json:"sshkeys,omitempty"`
	CiUpgrade    bool   `json:"ciupgrade,omitempty"`
	CiCustom     string `json:"cicustom,omitempty"`
	//Options//
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Startup     string `json:"startup,omitempty"`
	OsType      string `json:"ostype,omitempty"`
	Boot        string `json:"boot,omitempty"`
	Agent       string `json:"agent,omitempty"`
	SmBios1     string `json:"smbios1,omitempty"`
	VmGenId     string `json:"vmgenid,omitempty"`
	Meta        string `json:"meta,omitempty"`
}

func (q *QEMU) GetConfig(ctx context.Context, c *client.Client) (*QEMU, error) {
	path := fmt.Sprintf("/nodes/%s/qemu/%d/config", q.Node, q.Id)

	raw, err := client.Get[GetConfigQEMUResponse](ctx, c, path, nil, nil)
	if err != nil {
		return nil, err
	}

	// Hardware //
	mem, err := strconv.Atoi(raw.Memory)
	if err != nil {
		q.Hardware.Memory = 0
	} else {
		q.Hardware.Memory = uint64(mem)
	}

	q.Hardware.Cores = raw.Cores
	q.Hardware.Bios = attributes.BiosType(raw.Bios)
	q.Hardware.Vga = (&attributes.Vga{}).ToDomain(raw.Vga)

	return q, nil
}

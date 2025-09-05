package core

import (
	"context"
	"errors"
	"fmt"

	"github.com/hoanghaix86/proxmox-api-go/proxmox/client"
)

type CreateQEMURequest struct {
	Id   uint64 `json:"vmid"`
	Node string `json:"node"`
	//Hardware//
	Memory         uint64 `json:"memory,omitempty"`
	Cpu            string `json:"cpu,omitempty"`
	Cores          uint64 `json:"cores,omitempty"`
	Bios           string `json:"bios,omitempty"`
	Vga            string `json:"vga,omitempty"`
	Machine        string `json:"machine,omitempty"`
	ScsiController string `json:"scsihw,omitempty"`
	Ide0           string `json:"ide0,omitempty"`
	Ide2           string `json:"ide2,omitempty"`
	Scsi0          string `json:"scsi0,omitempty"`
	EfiDisk0       string `json:"efidisk0,omitempty"`
	TpmState0      string `json:"tpmstate0,omitempty"`
	Net0           string `json:"net0,omitempty"`
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
	OnBoot      bool   `json:"onboot,omitempty"`
	Template    bool   `json:"template,omitempty"`
	Serial0     string `json:"serial0,omitempty"`
}

func (q *QEMU) ToCreateQEMURequest() *CreateQEMURequest {
	return &CreateQEMURequest{
		Id:   q.Id,
		Node: q.Node,

		//Hardware//
		Memory:         q.Hardware.Memory,
		Cpu:            string(q.Hardware.Cpu),
		Cores:          q.Hardware.Cores,
		Bios:           string(q.Hardware.Bios),
		Vga:            q.Hardware.Vga.ToApi(),
		Machine:        string(q.Hardware.Machine),
		ScsiController: q.Hardware.ScsiController,
		Ide0:           q.Hardware.Ide0.ToApi(),
		Ide2:           q.Hardware.Ide2.ToApi(),
		Scsi0:          q.Hardware.Scsi0.ToApi(),
		EfiDisk0:       q.Hardware.EfiDisk0.ToApi(),
		TpmState0:      q.Hardware.TpmState0.ToApi(),
		Net0:           q.Hardware.Net0.ToApi(),
		//Cloudinit//
		CiUser:       q.Cloudinit.CiUser,
		CiPassword:   q.Cloudinit.CiPassword,
		SearchDomain: q.Cloudinit.SearchDomain,
		NameServer:   q.Cloudinit.NameServer,
		SshKeys:      q.Cloudinit.SshKeys,
		CiUpgrade:    q.Cloudinit.CiUpgrade,
		CiCustom:     q.Cloudinit.CiCustom,
		//Options//
		Name:        q.Options.Name,
		Description: q.Options.Description,
		Startup:     q.Options.Startup,
		OsType:      string(q.Options.OsType),
		Boot:        q.Options.Boot,
		Agent:       q.Agent.ToApi(),
		OnBoot:      q.Options.OnBoot,
		Template:    q.Options.Template,
		Serial0:     q.Serial0,
	}
}

func (q *QEMU) Create(ctx context.Context, c *client.Client) (*UPID, error) {
	if _, err := q.GetConfig(ctx, c); err == nil {
		return nil, errors.New("VM already exists")
	}

	path := fmt.Sprintf("/nodes/%s/qemu", q.Node)
	upid, err := client.Post[UPID](ctx, c, path, nil, q.ToCreateQEMURequest())
	if err != nil {
		return nil, err
	}

	return upid, nil
}

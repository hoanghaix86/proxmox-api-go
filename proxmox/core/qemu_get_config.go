package core

import (
	"context"
	"fmt"

	"github.com/hoanghaix86/proxmox-api-go/proxmox/attributes"
	"github.com/hoanghaix86/proxmox-api-go/proxmox/client"
)

type GetConfigQEMUResponse struct {
	// Hardware //
	TpmState0 string `json:"tpmstate0,omitempty"`
	Net0      string `json:"net0,omitempty"`

	Boot    string `json:"boot,omitempty"`    // computed
	VmGenId string `json:"vmgenid,omitempty"` // computed
	Meta    string `json:"meta,omitempty"`    // computed
	SmBios1 string `json:"smbios1,omitempty"` // computed
	Ide2    string `json:"ide2,omitempty"`
	Agent   string `json:"agent,omitempty"`

	Bios string `json:"bios,omitempty"`
}

func (q *QEMU) GetConfig(ctx context.Context, c *client.Client) (*QEMU, error) {
	path := fmt.Sprintf("/nodes/%s/qemu/%d/config", q.Node, q.Id)

	raw, err := client.Get[GetConfigQEMUResponse](ctx, c, path, nil, nil)
	if err != nil {
		return nil, err
	}

	// Hardware //
	q.Hardware.Ide2 = (&attributes.Ide{}).ToDomain(raw.Ide2)
	q.Hardware.TpmState0 = (&attributes.TpmState{}).ToDomain(raw.TpmState0)
	q.Hardware.Net0 = (&attributes.Network{}).ToDomain(raw.Net0)

	q.Boot = raw.Boot
	q.Meta = raw.Meta
	q.VmGenId = raw.VmGenId
	q.SmBios1 = raw.SmBios1

	q.Agent = (&attributes.Agent{}).ToDomain(raw.Agent)
	q.Bios = raw.Bios

	return q, nil
}

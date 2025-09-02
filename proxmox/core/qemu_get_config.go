package core

import (
	"context"
	"fmt"

	"github.com/hoanghaix86/proxmox-api-go/proxmox/attributes"
	"github.com/hoanghaix86/proxmox-api-go/proxmox/client"
)

type GetConfigQEMUResponse struct {
	Boot    string `json:"boot,omitempty"`    // computed
	VmGenId string `json:"vmgenid,omitempty"` // computed
	Meta    string `json:"meta,omitempty"`    // computed
	SmBios1 string `json:"smbios1,omitempty"` // computed

	Agent string `json:"agent,omitempty"`
}

func (q *QEMU) GetConfig(ctx context.Context, c *client.Client) (*QEMU, error) {
	path := fmt.Sprintf("/nodes/%s/qemu/%d/config", q.Node, q.Id)

	raw, err := client.Get[GetConfigQEMUResponse](ctx, c, path, nil, nil)
	if err != nil {
		return nil, err
	}

	q.Boot = raw.Boot
	q.Meta = raw.Meta
	q.VmGenId = raw.VmGenId
	q.SmBios1 = raw.SmBios1

	q.Agent = (&attributes.Agent{}).ToDomain(raw.Agent)

	return q, nil
}

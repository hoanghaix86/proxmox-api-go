package core

import (
	"context"
	"fmt"

	"github.com/hoanghaix86/proxmox-api-go/proxmox/client"
)

type UpdateConfigQEMURequest struct {
	Agent string `json:"agent,omitempty"`
	Bios  string `json:"bios,omitempty"`
}

func (q *QEMU) ToUpdateConfigQEMURequest() *UpdateConfigQEMURequest {
	return &UpdateConfigQEMURequest{
		Agent: q.Agent.ToApi(),
		Bios:  q.Bios,
	}
}

func (q *QEMU) UpdateConfig(ctx context.Context, c *client.Client) (*UPID, error) {
	path := fmt.Sprintf("/nodes/%s/qemu/%d/config", q.Node, q.Id)

	body := q.ToUpdateConfigQEMURequest()

	upid, err := client.Post[UPID](ctx, c, path, nil, body)
	if err != nil {
		return nil, err
	}
	return upid, nil
}

func (q *QEMU) EnableAgent(ctx context.Context, c *client.Client) (*UPID, error) {
	q.Agent.Enabled = true
	return q.UpdateConfig(ctx, c)
}

func (q *QEMU) DisableAgent(ctx context.Context, c *client.Client) (*UPID, error) {
	q.Agent.Enabled = false
	return q.UpdateConfig(ctx, c)
}

func (q *QEMU) SetBios(ctx context.Context, c *client.Client, bios string) (*UPID, error) {
	q.Bios = bios
	return q.UpdateConfig(ctx, c)
}

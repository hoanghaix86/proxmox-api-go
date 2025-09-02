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

	Agent string `json:"agent,omitempty"`
}

func (q *QEMU) ToCreateQEMURequest() *CreateQEMURequest {
	return &CreateQEMURequest{
		Id:   q.Id,
		Node: q.Node,

		Agent: q.Agent.ToApi(),
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

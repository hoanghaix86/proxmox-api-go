package core

import (
	"context"
	"errors"
	"fmt"

	"github.com/hoanghaix86/proxmox-api-go/proxmox/client"
)

type CloneQEMURequest struct {
	ToId        uint64 `json:"newid"`
	ToNode      string `json:"target"`
	Full        bool   `json:"full,omitempty"`
	Storage     string `json:"storage,omitempty"`
	Format      string `json:"format,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

func (q *QEMU) ToCloneQEMURequest() *CloneQEMURequest {
	return &CloneQEMURequest{
		ToId:        q.Id,
		ToNode:      q.Node,
		Full:        true,
		Storage:     q.Options.Storage,
		Format:      "qcow2",
		Name:        q.Options.Name,
		Description: q.Options.Description,
	}
}

func (q *QEMU) Clone(ctx context.Context, c *client.Client, clone *QEMU) (*UPID, error) {
	if clone == nil {
		return nil, errors.New("target vm is nil")
	}

	// check target vm exists
	if _, err := clone.GetConfig(ctx, c); err == nil {
		return nil, errors.New("vm clone already exists")
	}

	// check template exists
	if _, err := q.GetConfig(ctx, c); err != nil {
		return nil, errors.New("template not found")
	}

	path := fmt.Sprintf("/nodes/%s/qemu/%d/clone", q.Node, q.Id)
	body := clone.ToCloneQEMURequest()

	upid, err := client.Post[UPID](ctx, c, path, nil, body)
	if err != nil {
		return nil, err
	}

	return upid, nil
}

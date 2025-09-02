package core

import (
	"context"
	"fmt"

	"github.com/hoanghaix86/proxmox-api-go/proxmox/client"
)

type UPID string

type QEMU struct {
	Id   uint64 `json:"id"`
	Node string `json:"node"`

	Boot    string `json:"boot,omitempty"`    // computed
	VmGenId string `json:"vmgenid,omitempty"` // computed
	Meta    string `json:"meta,omitempty"`    // computed
	SmBios1 string `json:"smbios1,omitempty"` // computed
}

type CreateQEMURequest struct {
	Id   uint64 `json:"vmid"`
	Node string `json:"node"`
}

func (q *QEMU) ToCreateQEMURequest() *CreateQEMURequest {
	return &CreateQEMURequest{
		Id:   q.Id,
		Node: q.Node,
	}
}

func (q *QEMU) Create(ctx context.Context, c *client.Client) *UPID {
	path := fmt.Sprintf("/nodes/%s/qemu", q.Node)
	upid := client.Post[UPID](ctx, c, path, nil, q.ToCreateQEMURequest())
	return upid
}

type DeleteQEMUArgs struct {
	DestroyUnreferencedDisks bool `url:"destroy-unreferenced-disks,omitempty"`
	Purge                    bool `url:"purge,omitempty"`
	SkipLock                 bool `url:"skiplock,omitempty"`
}

func NewDeleteQEMUArgs() *DeleteQEMUArgs {
	return &DeleteQEMUArgs{
		DestroyUnreferencedDisks: true,
		Purge:                    true,
		SkipLock:                 false,
	}
}

func (q *QEMU) Delete(ctx context.Context, c *client.Client, args *DeleteQEMUArgs) *UPID {
	path := fmt.Sprintf("/nodes/%s/qemu/%d", q.Node, q.Id)
	if args == nil {
		args = NewDeleteQEMUArgs()
	}
	upid := client.Delete[UPID](ctx, c, path, args, nil)
	return upid
}

func (q *QEMU) GetConfig(ctx context.Context, c *client.Client) (*UPID, error) {
	return nil, nil
}

func (q *QEMU) UpdateConfig(ctx context.Context, c *client.Client) (*UPID, error) {
	return nil, nil
}

package core

import (
	"context"
	"errors"
	"fmt"

	"github.com/hoanghaix86/proxmox-api-go/proxmox/client"
)

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

func (q *QEMU) Delete(ctx context.Context, c *client.Client, args *DeleteQEMUArgs) (*UPID, error) {
	if _, err := q.GetConfig(ctx, c); err != nil {
		return nil, errors.New("VM does not exist")
	}

	path := fmt.Sprintf("/nodes/%s/qemu/%d", q.Node, q.Id)

	if args == nil {
		args = NewDeleteQEMUArgs()
	}

	upid, err := client.Delete[UPID](ctx, c, path, args, nil)
	if err != nil {
		return nil, err
	}
	return upid, nil
}

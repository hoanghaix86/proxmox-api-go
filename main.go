package main

import (
	"context"
	"fmt"

	"github.com/hoanghaix86/proxmox-api-go/proxmox/client"
	"github.com/hoanghaix86/proxmox-api-go/proxmox/core"
)

func main() {
	fmt.Println("PROXMOX API GO BETA")

	ctx := context.Background()

	client := client.NewClient(nil, nil)
	client.AuthWithToken(nil, nil)

	vm := core.QEMU{
		Id:   300,
		Node: "proxmox",
	}

	// upid := vm.Create(ctx, client)
	// fmt.Println(*upid)
	upid := vm.Delete(ctx, client, nil)
	fmt.Println(*upid)
}

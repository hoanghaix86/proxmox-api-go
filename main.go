package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hoanghaix86/proxmox-api-go/proxmox/attributes"
	"github.com/hoanghaix86/proxmox-api-go/proxmox/client"
	"github.com/hoanghaix86/proxmox-api-go/proxmox/core"
)

func PrintJson(v any) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal json: %v", err)
	}
	fmt.Println(string(b))
}

func main() {
	fmt.Println("PROXMOX API GO BETA")

	ctx := context.Background()

	client := client.NewClient(nil, nil)
	client.AuthWithToken(nil, nil)

	vm := core.QEMU{
		Id:   300,
		Node: "proxmox",
		Agent: &attributes.Agent{
			Enabled:           true,
			FreezeFsOnBackup:  true,
			FstrimClonedDisks: true,
			Type:              "virtio",
		},
	}

	// upid, err := vm.Create(ctx, client)
	// if err != nil {
	// 	log.Fatalf("%s", err.Error())
	// }
	// fmt.Println(*upid)
	// upid, err := vm.Delete(ctx, client, nil)
	// if err != nil {
	// 	log.Fatalf("getconfig: %s", err.Error())
	// }
	// fmt.Println(*upid)
	// config, err := vm.GetConfig(ctx, client)
	// if err != nil {
	// 	log.Fatalf("getconfig: %s", err.Error())
	// }
	// PrintJson(config)

	// update config
	upid, err := vm.UpdateConfig(ctx, client)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	fmt.Println(*upid)
}

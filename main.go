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
		Hardware: core.Hardware{
			Memory:         1024,
			Cpu:            attributes.CpuTypeX8664V2AES,
			Cores:          1,
			Bios:           attributes.BiosTypeOvmf,
			Vga:            attributes.NewDefaultVga(attributes.VgaTypeQxl),
			Machine:        "q35",
			ScsiController: "virtio-scsi-single",
			Ide2: &attributes.Ide{
				Volume: "local",
				Iso:    "iso/ubuntu-24.04.2-live-server-amd64.iso",
			},
			Scsi0:     "local-lvm:16,iothread=on,ssd=1,discard=on",
			EfiDisk0:  attributes.NewDefaultEfIdisk("local-lvm"),
			TpmState0: attributes.NewDefaultTpmState("local-lvm"),
			Net0:      attributes.NewDefaultNetwork("vmbr0"),
		},
	}

	// PrintJson(vm)

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

	// get config
	// time.Sleep(4 * time.Second)
	config, err := vm.GetConfig(ctx, client)
	if err != nil {
		log.Fatalf("getconfig: %s", err.Error())
	}
	PrintJson(config)

	// update config
	// upid, err := vm.UpdateConfig(ctx, client)
	// if err != nil {
	// 	log.Fatalf("%s", err.Error())
	// }
	// fmt.Println(*upid)
}

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
			Machine:        attributes.MachineQ35,
			ScsiController: "virtio-scsi-single",
			Ide0:           attributes.NewIdeCloudinit("local-lvm"),
			Ide2:           attributes.NewIdeIso("local", "iso/ubuntu-24.04.2-live-server-amd64.iso"),
			Scsi0:          attributes.NewScsi("local-lvm", 16),
			EfiDisk0:       attributes.NewDefaultEfIdisk("local-lvm"),
			TpmState0:      attributes.NewDefaultTpmState("local-lvm"),
			Net0:           attributes.NewDefaultNetwork("vmbr0"),
		},
		Options: core.Options{
			Name:        "testing",
			Description: "this is a testing",
			Startup:     "order=1,up=10,down=10",
			OsType:      attributes.OsTypeL26,
			Boot:        "order=scsi0;ide2;net0",
			Agent:       attributes.NewAgent(),
			OnBoot:      true,
		},
	}

	upid, err := vm.Create(ctx, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*upid)
	// upid, err := vm.Delete(ctx, client, nil)
	// if err != nil {
	// 	log.Fatalf("getconfig: %s", err.Error())
	// }
	// fmt.Println(*upid)

	// get config
	// time.Sleep(4 * time.Second)
	// config, err := vm.GetConfig(ctx, client)
	// if err != nil {
	// 	log.Fatalf("getconfig: %s", err.Error())
	// }
	// PrintJson(config)

	// update config
	// upid, err := vm.UpdateConfig(ctx, client)
	// if err != nil {
	// 	log.Fatalf("%s", err.Error())
	// }
	// fmt.Println(*upid)
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

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
	}

	targetVM := core.QEMU{
		Id:   400,
		Node: "proxmox",
		Options: core.Options{
			Name:    "testing-clone",
			Storage: "local-lvm",
		},
	}

	upid, err := vm.Clone(ctx, client, &targetVM)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*upid)

	// upid, err := vm.Create(ctx, client)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(*upid)
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

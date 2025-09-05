# PROXMOX API GO

## Example

1. Create VM

```go
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
		},
	}

	upid, err := vm.Create(ctx, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*upid)
}
```
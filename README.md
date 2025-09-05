# PROXMOX API GO

## Example

1. Create VM with ISO

```go
func main() {
	ctx := context.Background()

	client := client.NewClient(nil, nil)
	client.AuthWithToken(nil, nil)

	vm := core.QEMU{
		Id:   300,
		Node: "proxmox",
		Hardware: core.Hardware{
			Memory:         4096,
			Cpu:            attributes.CpuTypeX8664V2AES,
			Cores:          4,
			Bios:           attributes.BiosTypeOvmf,
			Vga:            attributes.NewVga(attributes.VgaTypeStd),
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
			OsType:      attributes.OsTypeL26,
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

2. Clone VM from template
# PROXMOX API GO

## Example

### 1. Create VM with ISO

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

### 2. Clone VM from template

1. Create cloudinit

```bash
mkdir -p /var/lib/vz/snippets

tee /var/lib/vz/snippets/vendor.yml <<EOF
#cloud-config
apt:
  preserve_sources_list: false
  primary:
    - arches:
        - default
      uri: http://mirror.viettelcloud.vn/ubuntu

runcmd:
  - apt-get update
  - apt-get install -y qemu-guest-agent
  - systemctl enable qemu-guest-agent --now
EOF
```

2. Download image

```bash
wget https://cloud-images.ubuntu.com/noble/current/noble-server-cloudimg-amd64.img
```

3. Create VM with qm cli

```bash
qm create 300 \
  --memory 4096 \
  --cpu x86-64-v2-AES \
  --cores 4 \
  --bios ovmf \
  --vga serial0 \
  --serial0 socket \
  --machine q35 \
  --scsihw virtio-scsi-single \
  --ide2 local-lvm:cloudinit \
  --scsi0 local-lvm:0,import-from=/root/noble-server-cloudimg-amd64.img,iothread=1,ssd=1,discard=on,format=qcow2 \
  --efidisk0 local-lvm:0 \
  --tpmstate0 local-lvm:0,version=v2.0 \
  --net0 virtio,bridge=vmbr0,firewall=0 \
  --ciuser runner \
  --cipassword runner \
  --ciupgrade 1 \
  --cicustom vendor=local:snippets/vendor.yml \
  --ipconfig0 ip=dhcp \
  --name noble-server \
  --ostype l26 \
  --agent '1,freeze-fs-on-backup=1,fstrim_cloned_disks=1,type=virtio' \
  --template 1
```
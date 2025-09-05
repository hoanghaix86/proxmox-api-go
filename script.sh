#!/usr/bin/env bash
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

qm create 300 \
  --memory 2048 \
  --cpu x86-64-v2-AES \
  --cores 2 \
  --bios ovmf \
  --vga serial0 \
  --serial0 'socket' \
  --machine q35 \
  --scsihw virtio-scsi-single \
  --ide2 local-lvm:cloudinit \
  --scsi0 local-lvm:0,import-from=/root/noble-server-cloudimg-amd64.img,iothread=1,ssd=1,discard=on,format=qcow2 \
  --efidisk0 local-lvm:0 \
  --tpmstate0 local-lvm:0,version=v2.0 \
  --net0 'virtio,bridge=vmbr0,firewall=0' \
  --ciuser 'runner' \
  --cipassword '1' \
  --nameserver '1.1.1.1' \
  --ipconfig0 'ip=dhcp' \
  --cicustom vendor=local:snippets/qemu-guest-agent.yml \
  --name ubuntu-noble-server \
  --description 'This is a template' \
  --ostype l26 \
  --boot 'order=scsi0;net0' \
  --agent '1,freeze-fs-on-backup=1,fstrim_cloned_disks=1' \
  --onboot 0 \
  --template 0


# XDP eBPF Packet Filter (Go + C)

## 🚀 Overview

This project demonstrates an **eBPF-based XDP packet filtering program** written in C with a Go userspace loader using the Cilium eBPF library.

The program attaches an XDP (eXpress Data Path) hook to a network interface and drops TCP packets targeting a configurable port (default: 4040).

---

## ⚙️ Features

- ✔ XDP-based kernel-level packet filtering
- ✔ Drops TCP packets on a configurable port
- ✔ Dynamic configuration using BPF maps
- ✔ Userspace control using Go (Cilium eBPF library)
- ✔ Lightweight and high-performance packet processing

---

## 🧠 Architecture


Go Userspace Program
|
| (updates BPF map)
↓
BPF Map (port configuration)
↓
eBPF XDP Program (kernel space)
↓
Network Interface (eth0)
↓
Packets filtered at kernel level


---

## 📌 How it Works

1. Go program loads compiled eBPF object (`xdp.o`)
2. It updates a BPF map with blocked port (e.g., 4040)
3. XDP program reads this value from kernel space
4. Incoming TCP packets are inspected
5. If destination port matches → packet is dropped

---

## 📦 Requirements

- Linux (WSL2 / Ubuntu)
- clang + llvm
- Go 1.23+
- libbpf-dev
- kernel with eBPF support

Install dependencies:

```bash
sudo apt update
sudo apt install -y clang llvm libbpf-dev gcc make
🛠 Build Instructions
1. Compile eBPF program
clang -O2 -g -target bpf -c xdp.c -o xdp.o
2. Run Go loader
sudo go run main.go
🧪 Testing

Start a simple HTTP server:

python3 -m http.server 4040 --bind 0.0.0.0

Then test:

curl http://<IP>:4040

If working on native Linux:

Traffic to port 4040 will be dropped

⚠️ Note: On WSL2, XDP packet dropping may not fully behave as expected due to virtualized networking limitations.

⚠️ Known Limitation (WSL2)

WSL2 does not fully support real XDP packet enforcement in all cases.

So:

Program attaches successfully ✔
Map updates work ✔
Packet drop behavior may not be visible ❌

For full testing, use:

Native Linux VM
or Cloud Ubuntu instance
🧩 Key Learning
eBPF XDP programming model
Kernel vs userspace interaction
BPF maps for dynamic configuration
Network packet processing at kernel level
Go-based eBPF control plane

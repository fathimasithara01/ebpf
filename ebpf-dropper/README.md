# eBPF XDP TCP Packet Dropper

## Overview

This project demonstrates a high-performance TCP packet filtering system using eBPF and XDP (eXpress Data Path).

The application attaches an XDP program to a network interface and drops TCP packets targeting a configurable port (default: `4040`) directly at the kernel networking layer.

The userspace controller is implemented in Go using the Cilium eBPF library, while the XDP program is written in C.

---

# Architecture

```text id="a1"
Go Userspace Loader
        │
        ▼
Loads eBPF Program (xdp.o)
        │
        ▼
Updates BPF Map (blocked port)
        │
        ▼
XDP Hook attached to eth0
        │
        ▼
Kernel-level packet inspection
        │
        ▼
Matching TCP packets dropped
```

---

# Features

* XDP-based packet filtering
* Kernel-level TCP packet dropping
* Configurable blocked port using BPF maps
* Go userspace control plane
* Lightweight and high-performance architecture
* Early packet processing at driver level

---

# Tech Stack

* Go
* eBPF
* XDP
* Cilium eBPF library
* C
* Linux Networking
* clang/LLVM

---

# Project Structure

```text id="a2"
ebpf-dropper/
│
├── main.go      # Go userspace loader
├── xdp.c        # XDP/eBPF packet filter
├── xdp.o        # Compiled eBPF object
├── go.mod
├── go.sum
└── README.md
```

---

# Build Instructions

## Install Dependencies

```bash id="a3"
sudo apt update
sudo apt install -y clang llvm libbpf-dev gcc make
```

## Compile eBPF Program

```bash id="a4"
clang -O2 -g -target bpf -c xdp.c -o xdp.o
```

## Run Userspace Loader

```bash id="a5"
sudo go run main.go
```

Expected Output:

```text id="a6"
XDP attached: TCP port filtering active
```

---

# Verification

Check whether XDP is attached:

```bash id="a7"
ip link show eth0
```

Expected:

```text id="a8"
prog/xdp
```

---

# Testing

Start a test HTTP server:

```bash id="a9"
python3 -m http.server 4040
```

Send traffic:

```bash id="a10"
curl http://localhost:4040
```

---

# WSL2 Limitation

This project was developed inside WSL2 Ubuntu.

While XDP attachment succeeds correctly, WSL2 virtualized networking may not fully enforce real packet dropping behavior.

For complete validation, testing on:

* Native Linux
* Linux VM
* Cloud Ubuntu instance

is recommended.

---

# Key Learnings

* eBPF/XDP programming model
* Kernel-space packet processing
* Userspace-to-kernel interaction
* BPF maps
* Linux networking internals
* High-performance packet filtering

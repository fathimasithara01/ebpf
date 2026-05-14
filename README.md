# AccuKnox eBPF Assignment Submission

## Overview

This repository contains solutions and technical explorations for the AccuKnox backend/eBPF assignment focused on:

* eBPF
* XDP
* Linux networking
* kernel-space packet filtering
* Go concurrency concepts

The assignment demonstrates understanding of:

* Linux kernel networking primitives
* eBPF/XDP architecture
* userspace-to-kernel interaction
* concurrent programming in Go
* systems-level debugging and exploration

---

# Repository Structure

```text id="root1"
ebpf-accuknox-assignment/
│
├── ebpf-dropper/
│   ├── README.md
│   ├── main.go
│   ├── xdp.c
│   └── xdp.o
│
├── ebpf-process-filter/
│   ├── README.md
│   ├── main.go
│   ├── xdp.c
│   └── xdp.o
│
├── go-concurrency/
│   ├── README.md
│   └── explanation.md
│
└── README.md
```

---

# Problem 1 — eBPF XDP TCP Packet Dropper

## Objective

Implement an eBPF/XDP program that drops TCP packets targeting a configurable port (default: `4040`).

## Implementation Summary

* XDP program written in C
* Userspace controller written in Go
* Dynamic configuration using BPF maps
* XDP attached to network interface (`eth0`)
* TCP packet inspection at kernel level

## Key Features

* Kernel-level packet filtering
* Configurable blocked port
* High-performance XDP hook
* Go + eBPF integration
* Lightweight networking architecture

## Verification

XDP attachment verified using:

```bash id="root2"
ip link show eth0
```

Observed:

```text id="root3"
prog/xdp
```

which confirms successful kernel attachment.

---

# Problem 2 — Process-aware Traffic Filtering Exploration

## Objective

Allow traffic only on a specific TCP port (`4040`) for a given process while blocking traffic to all other ports for that process.

## Exploration Summary

The implementation explored process-aware filtering using XDP and identified important architectural limitations.

### Key Findings

* XDP executes very early in packet processing
* Process context is unavailable at XDP layer
* Helpers like:

```c id="root4"
bpf_get_current_comm()
```

are not supported in XDP programs

## Alternative Approaches Explored

* cgroup socket hooks
* socket-layer eBPF
* tc-based filtering
* process-aware socket filtering

## Outcome

The project evolved into a deeper exploration of:

* Linux networking layers
* eBPF verifier restrictions
* hook-specific capabilities
* process vs packet context separation

---

# Problem 3 — Go Concurrency Analysis

## Objective

Analyze and explain a Go concurrency code snippet involving:

* goroutines
* buffered channels
* worker pools
* scheduling behavior

## Topics Covered

* Worker pool architecture
* Concurrent task execution
* Buffered channels
* Goroutine scheduling
* Synchronization issues
* Runtime behavior analysis

A detailed explanation is included in:

```text id="root5"
go-concurrency/explanation.md
```

---

# Technologies Used

* Go
* eBPF
* XDP
* Linux Networking
* C
* clang/LLVM
* Cilium eBPF library

---

# Environment

Development environment:

* Ubuntu (WSL2)
* Linux kernel with eBPF support
* Go 1.23+
* clang/LLVM

---

# WSL2 Note

The projects were developed and tested inside WSL2 Ubuntu.

While XDP attachment and eBPF loading worked successfully, real packet dropping behavior may not always be fully observable due to WSL2 virtualized networking limitations.

For complete validation:

* Native Linux
* Linux VM
* Cloud Ubuntu instance

are recommended.

---

# Key Learnings

* eBPF/XDP programming model
* Kernel-space networking
* BPF maps and dynamic configuration
* Linux networking internals
* Hook-specific limitations in eBPF
* Go concurrency patterns
* Goroutine scheduling behavior
* Systems-level debugging

---

# Author

Fathima Sithara

Submitted as part of the AccuKnox backend/eBPF evaluation assignment.

# Process-aware Traffic Filtering Exploration using eBPF

## Overview

This project explores process-aware TCP traffic filtering using eBPF.

The goal was to allow traffic only on a specific TCP port (`4040`) for a given process while blocking traffic to all other ports for that process.

During implementation, the project evolved into an exploration of Linux networking hooks, eBPF verifier restrictions, and process context limitations inside XDP programs.

---

# Objective

Allow:

```text id="b1"
Process → Port 4040
```

Block:

```text id="b2"
Process → Any other TCP port
```

---

# Initial Approach

The initial implementation attempted to use XDP-based packet filtering.

However, during development it was identified that:

* XDP executes extremely early in the networking stack
* Process/task context is unavailable at XDP layer
* Helper functions such as:

```c id="b3"
bpf_get_current_comm()
```

are not supported inside XDP programs

---

# Technical Findings

## Why Process Filtering is Difficult in XDP

XDP programs operate:

* before socket creation
* before process association
* directly at driver-level packet receive path

At that stage:

* packets are not yet associated with a userspace process

Therefore:

* process-aware filtering cannot be reliably implemented using XDP alone.

---

# Alternative Approaches Explored

* cgroup socket hooks
* socket-level eBPF
* tc-based filtering
* process-aware socket filtering

---

# Tech Stack

* Go
* eBPF
* Linux Networking
* XDP
* clang/LLVM
* Cilium eBPF library

---

# Project Structure

```text id="b4"
ebpf-process-filter/
│
├── main.go
├── xdp.c
├── xdp.o
├── go.mod
├── go.sum
└── README.md
```

---

# Build Instructions

## Compile eBPF Program

```bash id="b5"
clang -O2 -g -target bpf -c xdp.c -o xdp.o
```

## Run Program

```bash id="b6"
sudo go run main.go
```

---

# Current Status

## Successfully Demonstrated

* eBPF compilation
* XDP attachment
* Kernel hook installation
* TCP packet inspection
* Exploration of process-aware filtering architecture

## Limitations Identified

True process-aware packet filtering requires:

* cgroup hooks
* socket-layer programs
* or tc/socket filters

rather than XDP alone.

---

# Key Learnings

* Differences between XDP and socket-layer hooks
* Process context availability in Linux networking
* eBPF verifier restrictions
* Hook-specific helper support
* Linux kernel packet processing flow
* Architectural tradeoffs in eBPF

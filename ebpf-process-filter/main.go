package main

import (
	"log"
	"net"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/rlimit"
)

func main() {
	// Allow loading large eBPF programs
	rlimit.RemoveMemlock()

	// Load compiled eBPF program
	spec, err := ebpf.LoadCollectionSpec("xdp.o")
	if err != nil {
		log.Fatal(err)
	}

	coll, err := ebpf.NewCollection(spec)
	if err != nil {
		log.Fatal(err)
	}
	defer coll.Close()

	prog := coll.Programs["xdp_filter"]
	if prog == nil {
		log.Fatal("program not found")
	}

	iface, err := net.InterfaceByName("eth0")
	if err != nil {
		log.Fatal(err)
	}

	l, err := link.AttachXDP(link.XDPOptions{
		Program:   prog,
		Interface: iface.Index,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	log.Println("XDP attached: TCP port filtering active")
	select {}
}

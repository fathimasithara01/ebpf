package main

import (
	"log"
	"net"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/rlimit"
)

type Objects struct {
	DropPort *ebpf.Program `ebpf:"drop_port"`
	PortMap  *ebpf.Map     `ebpf:"port_map"`
}

func main() {
	// allow loading eBPF programs
	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatal(err)
	}

	// load compiled object
	spec, err := ebpf.LoadCollectionSpec("xdp.o")
	if err != nil {
		log.Fatalf("loading spec: %v", err)
	}

	var objs Objects

	if err := spec.LoadAndAssign(&objs, nil); err != nil {
		log.Fatalf("loading objects: %v", err)
	}

	// set blocked port (configurable)
	var key uint32 = 0
	var port uint32 = 4040

	if err := objs.PortMap.Put(key, port); err != nil {
		log.Fatalf("updating map: %v", err)
	}

	// attach to interface
	iface, err := net.InterfaceByName("eth0")
	if err != nil {
		log.Fatal(err)
	}

	l, err := link.AttachXDP(link.XDPOptions{
		Program:   objs.DropPort,
		Interface: iface.Index,
	})

	if err != nil {
		log.Fatalf("attach xdp: %v", err)
	}

	defer l.Close()

	log.Println("XDP attached. Blocking port:", port)

	select {}
}

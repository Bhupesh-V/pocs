package main

import (
	"fmt"
	"log"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <interface> <filter>")
		fmt.Println("Example: go run main.go eth0 'host server_a or host server_b'")
		os.Exit(1)
	}

	iface := os.Args[1]  // Network interface to capture on, e.g., "eth0"
	filter := os.Args[2] // BPF filter expression, e.g., "host server_a or host server_b"

	fmt.Printf("Capturing on interface: %s with filter: %s\n", iface, filter)

	handle, err := pcap.OpenLive(iface, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatalf("Error opening device %s: %v", iface, err)
	}
	defer handle.Close()

	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatalf("Error setting BPF filter: %v", err)
	}

	fmt.Println("Starting packet capture...")

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		fmt.Println("Packet captured:", packet)
		transportLayer := packet.TransportLayer()
		if transportLayer != nil {
			if tcpLayer, ok := transportLayer.(*layers.TCP); ok {
				fmt.Printf("Captured TCP packet: %s:%d -> %s:%d\n",
					packet.NetworkLayer().NetworkFlow().Src(), tcpLayer.SrcPort,
					packet.NetworkLayer().NetworkFlow().Dst(), tcpLayer.DstPort)
			} else {
				fmt.Println("Not a TCP packet")
			}
		} else {
			fmt.Println("No transport layer found in packet")
		}
	}
}

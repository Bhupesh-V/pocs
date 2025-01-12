# packet-watcher

A random utility I built to capture packets on a network and display them with no end goal

## Usage

```bash
sudo go run main.go lo0 'tcp port 8080'
```

> `lo0` is used to capture packets on the loopback interface. All localhost traffic is captured using this interface.

In another terminal fire up a server on port 8080

```bash
nc -l 8080
```

Send a message to our demo server

```bash
curl http://localhost:8080
```

You should see the packet in the packet watcher terminal. I am sure the output is not clean but I wasn't motivated enough to do anything after I got the packets. That's it.

```
Capturing on interface: lo0 with filter: tcp port 8080
Starting packet capture...
Packet captured: PACKET: 88 bytes, wire length 88 cap length 88 @ 2025-01-12 13:51:11.928227 +0530 IST
- Layer 1 (04 bytes) = Loopback {Contents=[30, 0, 0, 0] Payload=[..84..] Family=IPv6}
- Layer 2 (40 bytes) = IPv6     {Contents=[..40..] Payload=[..44..] Version=6 TrafficClass=0 FlowLabel=854528 Length=44 NextHeader=TCP HopLimit=64 SrcIP=::1 DstIP=::1 HopByHop=nil}
- Layer 3 (44 bytes) = TCP      {Contents=[..44..] Payload=[] SrcPort=50049 DstPort=8080(http-alt) Seq=2528661141 Ack=0 DataOffset=11 FIN=false SYN=true RST=false PSH=false ACK=false URG=false ECE=false CWR=false NS=false Window=65535 Checksum=52 Urgent=0 Options=[..8..] Padding=[0]}

Captured TCP packet: ::1:50049 -> ::1:8080
Packet captured: PACKET: 64 bytes, wire length 64 cap length 64 @ 2025-01-12 13:51:11.928281 +0530 IST
- Layer 1 (04 bytes) = Loopback {Contents=[30, 0, 0, 0] Payload=[..60..] Family=IPv6}
- Layer 2 (40 bytes) = IPv6     {Contents=[..40..] Payload=[..20..] Version=6 TrafficClass=0 FlowLabel=854528 Length=20 NextHeader=TCP HopLimit=64 SrcIP=::1 DstIP=::1 HopByHop=nil}
- Layer 3 (20 bytes) = TCP      {Contents=[..20..] Payload=[] SrcPort=8080(http-alt) DstPort=50049 Seq=0 Ack=2528661142 DataOffset=5 FIN=false SYN=false RST=true PSH=false ACK=true URG=false ECE=false CWR=false NS=false Window=0 Checksum=28 Urgent=0 Options=[] Padding=[]}

Captured TCP packet: ::1:8080 -> ::1:50049
Packet captured: PACKET: 68 bytes, wire length 68 cap length 68 @ 2025-01-12 13:51:11.928414 +0530 IST
- Layer 1 (04 bytes) = Loopback {Contents=[2, 0, 0, 0] Payload=[..64..] Family=IPv4}
- Layer 2 (20 bytes) = IPv4     {Contents=[..20..] Payload=[..44..] Version=4 IHL=5 TOS=0 Length=64 Id=0 Flags=DF FragOffset=0 TTL=64 Protocol=TCP Checksum=0 SrcIP=127.0.0.1 DstIP=127.0.0.1 Options=[] Padding=[]}
- Layer 3 (44 bytes) = TCP      {Contents=[..44..] Payload=[] SrcPort=50050 DstPort=8080(http-alt) Seq=3014653227 Ack=0 DataOffset=11 FIN=false SYN=true RST=false PSH=false ACK=false URG=false ECE=false CWR=false NS=false Window=65535 Checksum=65076 Urgent=0 Options=[..8..] Padding=[0]}

Captured TCP packet: 127.0.0.1:50050 -> 127.0.0.1:8080
Packet captured: PACKET: 68 bytes, wire length 68 cap length 68 @ 2025-01-12 13:51:11.928606 +0530 IST
- Layer 1 (04 bytes) = Loopback {Contents=[2, 0, 0, 0] Payload=[..64..] Family=IPv4}
- Layer 2 (20 bytes) = IPv4     {Contents=[..20..] Payload=[..44..] Version=4 IHL=5 TOS=0 Length=64 Id=0 Flags=DF FragOffset=0 TTL=64 Protocol=TCP Checksum=0 SrcIP=127.0.0.1 DstIP=127.0.0.1 Options=[] Padding=[]}
- Layer 3 (44 bytes) = TCP      {Contents=[..44..] Payload=[] SrcPort=8080(http-alt) DstPort=50050 Seq=2460021093 Ack=3014653228 DataOffset=11 FIN=false SYN=true RST=false PSH=false ACK=true URG=false ECE=false CWR=false NS=false Window=65535 Checksum=65076 Urgent=0 Options=[..8..] Padding=[0]}

Captured TCP packet: 127.0.0.1:8080 -> 127.0.0.1:50050
Packet captured: PACKET: 56 bytes, wire length 56 cap length 56 @ 2025-01-12 13:51:11.928632 +0530 IST
- Layer 1 (04 bytes) = Loopback {Contents=[2, 0, 0, 0] Payload=[..52..] Family=IPv4}
- Layer 2 (20 bytes) = IPv4     {Contents=[..20..] Payload=[..32..] Version=4 IHL=5 TOS=0 Length=52 Id=0 Flags=DF FragOffset=0 TTL=64 Protocol=TCP Checksum=0 SrcIP=127.0.0.1 DstIP=127.0.0.1 Options=[] Padding=[]}
- Layer 3 (32 bytes) = TCP      {Contents=[..32..] Payload=[] SrcPort=50050 DstPort=8080(http-alt) Seq=3014653228 Ack=2460021094 DataOffset=8 FIN=false SYN=false RST=false PSH=false ACK=true URG=false ECE=false CWR=false NS=false Window=6379 Checksum=65064 Urgent=0 Options=[TCPOption(NOP:), TCPOption(NOP:), TCPOption(Timestamps:376113632/1423538061 0x166b09e054d9778d)] Padding=[]}

Captured TCP packet: 127.0.0.1:50050 -> 127.0.0.1:8080
Packet captured: PACKET: 56 bytes, wire length 56 cap length 56 @ 2025-01-12 13:51:11.928646 +0530 IST
- Layer 1 (04 bytes) = Loopback {Contents=[2, 0, 0, 0] Payload=[..52..] Family=IPv4}
- Layer 2 (20 bytes) = IPv4     {Contents=[..20..] Payload=[..32..] Version=4 IHL=5 TOS=0 Length=52 Id=0 Flags=DF FragOffset=0 TTL=64 Protocol=TCP Checksum=0 SrcIP=127.0.0.1 DstIP=127.0.0.1 Options=[] Padding=[]}
- Layer 3 (32 bytes) = TCP      {Contents=[..32..] Payload=[] SrcPort=8080(http-alt) DstPort=50050 Seq=2460021094 Ack=3014653228 DataOffset=8 FIN=false SYN=false RST=false PSH=false ACK=true URG=false ECE=false CWR=false NS=false Window=6379 Checksum=65064 Urgent=0 Options=[TCPOption(NOP:), TCPOption(NOP:), TCPOption(Timestamps:1423538061/376113632 0x54d9778d166b09e0)] Padding=[]}

Captured TCP packet: 127.0.0.1:8080 -> 127.0.0.1:50050
Packet captured: PACKET: 133 bytes, wire length 133 cap length 133 @ 2025-01-12 13:51:11.92869 +0530 IST
- Layer 1 (04 bytes) = Loopback {Contents=[2, 0, 0, 0] Payload=[..129..] Family=IPv4}
- Layer 2 (20 bytes) = IPv4     {Contents=[..20..] Payload=[..109..] Version=4 IHL=5 TOS=0 Length=129 Id=0 Flags=DF FragOffset=0 TTL=64 Protocol=TCP Checksum=0 SrcIP=127.0.0.1 DstIP=127.0.0.1 Options=[] Padding=[]}
- Layer 3 (32 bytes) = TCP      {Contents=[..32..] Payload=[..77..] SrcPort=50050 DstPort=8080(http-alt) Seq=3014653228 Ack=2460021094 DataOffset=8 FIN=false SYN=false RST=false PSH=true ACK=true URG=false ECE=false CWR=false NS=false Window=6379 Checksum=65141 Urgent=0 Options=[TCPOption(NOP:), TCPOption(NOP:), TCPOption(Timestamps:376113632/1423538061 0x166b09e054d9778d)] Padding=[]}
- Layer 4 (77 bytes) = Payload  77 byte(s)

Captured TCP packet: 127.0.0.1:50050 -> 127.0.0.1:8080
Packet captured: PACKET: 56 bytes, wire length 56 cap length 56 @ 2025-01-12 13:51:11.928713 +0530 IST
- Layer 1 (04 bytes) = Loopback {Contents=[2, 0, 0, 0] Payload=[..52..] Family=IPv4}
- Layer 2 (20 bytes) = IPv4     {Contents=[..20..] Payload=[..32..] Version=4 IHL=5 TOS=0 Length=52 Id=0 Flags=DF FragOffset=0 TTL=64 Protocol=TCP Checksum=0 SrcIP=127.0.0.1 DstIP=127.0.0.1 Options=[] Padding=[]}
- Layer 3 (32 bytes) = TCP      {Contents=[..32..] Payload=[] SrcPort=8080(http-alt) DstPort=50050 Seq=2460021094 Ack=3014653305 DataOffset=8 FIN=false SYN=false RST=false PSH=false ACK=true URG=false ECE=false CWR=false NS=false Window=6378 Checksum=65064 Urgent=0 Options=[TCPOption(NOP:), TCPOption(NOP:), TCPOption(Timestamps:1423538061/376113632 0x54d9778d166b09e0)] Padding=[]}

Captured TCP packet: 127.0.0.1:8080 -> 127.0.0.1:50050
```

To capture packets on a different interface, replace `lo0` with the interface name. You can find the interface name by running `ifconfig` on your terminal.

```bash
sudo go run main.go en0 'host google.com'
```

Then you can send a request to google.com (using curl or a browser) and see the packets being captured by the utility.

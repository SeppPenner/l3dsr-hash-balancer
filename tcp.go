package balancer

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

// TCPState defines the state a TCP connection has.
type TCPState uint8

const (
	_                           = iota
	TCP_STATE_SYN_SENT TCPState = iota
	TCP_STATE_SYN_RECEIVED
	TCP_STATE_ESTABLISHED
	TCP_STATE_FIN_WAIT_1
	TCP_STATE_FIN_WAIT_2
	TCP_STATE_CLOSE_WAIT
	TCP_STATE_CLOSING
	TCP_STATE_LAST_ACK
	TCP_STATE_TIME_WAIT
	TCP_STATE_CLOSED
)

// TCPPacket represents a TCP packet.
type TCPPacket struct {
	ip  *layers.IPv4
	tcp *layers.TCP
}

// NewTCPPacket creates and initializes a new TCP packet.
// The IP layer is needed to calculate the correct TCP checksum.
func NewTCPPacket(ip *layers.IPv4, tcp *layers.TCP) TCPPacket {
	return TCPPacket{
		ip:  ip,
		tcp: tcp,
	}
}

// MarshalBinary returns the binary representation of the packet.
func (p TCPPacket) MarshalBinary() ([]byte, error) {
	p.tcp.SetNetworkLayerForChecksum(p.ip)
	opts := gopacket.SerializeOptions{
		ComputeChecksums: true,
		FixLengths:       true,
	}
	buf := gopacket.NewSerializeBuffer()
	err := gopacket.SerializeLayers(buf, opts, p.tcp, gopacket.Payload(p.tcp.Payload))
	return buf.Bytes(), err
}

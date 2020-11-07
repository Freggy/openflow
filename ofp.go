package main

// Header of each openflow message. Header is 8 byte in lengh
type Header struct {
	// OFP_VERSION
	Version uint8 // OFP_VERSION

	// One of the OFPT_ constants
	Type uint8

	// Length of the message including this Lenght
	Length uint16

	// Transaction id associated with this packet.
	// Replies use the same id as was in the request to faciliate pairing
	XID uint32
}

// Openflow message
type Message struct {
	Header  Header
	Payload []byte
}

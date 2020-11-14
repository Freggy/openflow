package ofp15

import "github.com/freggy/openflow"

//
// OFP_HELLO
//

// OFP_HELLO. This message includes zero or more hello elements
// having variable size. Unknown elements must be ignored/skipped, to allow
// for future extensions.
type OfpHello struct {
	Header openflow.Header

	Elements []OfpHelloElemVersionBitmap
}

type OfpHelloElemType int

const (
	OfphetVersionBitmap OfpHelloElemType = 1
)

// Common header for all Hello Elements
type OfpHelloElemHeader struct {
	// One of Ofphet*
	Type uint16

	// Lenght in bytes of the element, including this header, excluding padding
	Length uint16
}

// Version bitmap Hello Element
type OfpHelloElemVersionBitmap struct {
	Header OfpHelloElemHeader

	// Followed by:
	// - Exactly (length - 4) bytes containing the bitmaps then
	// - Exactly (length + 7)/8*8 - (length) (between 0 and 7)
	//   bytes of all zero bytes
	Bitmaps []uint32
}

//
// OFP_ECHO_REQUEST
//

// An Echo Request Message consists of an OpenFlow header plus an
// arbitrary length data field (in this case a byte slice) to verify
// liveness between the switch and controller.
type OfpEchoRequest struct {
	Header  openflow.Header
	Payload []byte
}

//
// OFP_ECHO_REPLY
//

// An Echo Reply message consists of an OpenFlow header and the
// unmodified data field of an Echo Request message.
type OfpEchoReply struct {
	Header  openflow.Header
	Payload []byte
}

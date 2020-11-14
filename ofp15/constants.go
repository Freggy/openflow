package ofp15

const (
	//
	// Immutable message
	//

	// Symmetric message
	OFPTHello = iota

	// Symmetric message
	OFPTError

	// Symmetric message
	OFPTEchoRequest

	// Symmetric message
	OFPTEchoReply

	// Symmetric message
	OFPTExperimenter

	//
	// Switch configuration messages
	//

	OFPTFeaturesRequest
	OFPTFeaturesReply
	OFPTGetConfigRequest
	OFPTGetConfigReply
	OFPTSetConfig

	//
	// Asynchronous messages
	//

	OFPTPacketIn
	OFPTFlowMod
	OFPTGroupMod
	OFPTPortMod
	OFPTTableMod

	//
	// Multipart messages
	//

	OFPTMultipartRequest
	OFPTMultipartReply

	//
	// Barrier messages
	//

	OFPTBarrierRequest
	OFPTBarrierReply

	//
	// Controller role change request message
	//

	OFPTRoleRequest
	OFPTRoleReply

	//
	// Asynchronous message configuration
	//

	OFPTGetAsyncRequest
	OFPTGetAsyncReply
	OFPTSetAsync

	//
	// Meters and rate limiters configuration messages
	//

	OFPTMeterMod

	//
	// Controller role change event messages
	//

	OFPTRoleStatus

	//
	// Asynchronous messages
	//

	OFPTTableStatus

	//
	// Request forwarding by the switch
	//

	OFPTRequestForward

	//
	// Bundle operations (multiple messages as single operation)
	//

	OFPTBundleControl
	OFPTBundleAddMessage

	//
	// Controller Status async message
	//

	OFPTControllerStatus
)

package protocol

type MSGAuthReconnectProof struct {
	Opcode       AuthCmd
	R1           [16]byte
	R2           any // SHA1
	R3           any // SHA1
	NumberOfKeys byte
}

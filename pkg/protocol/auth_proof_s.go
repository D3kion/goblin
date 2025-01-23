package protocol

type SMSGAuthProof struct {
	Opcode AuthCmd
	Error  byte
	M2     any // SHA1
	Unk1   uint32
	Unk2   uint32
	Unk3   uint16
}

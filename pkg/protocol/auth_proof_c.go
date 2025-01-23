package protocol

type CMSGAuthProof struct {
	Opcode        AuthCmd
	A             any // SRP6 key
	ClientM       any // SHA1
	CrcHash       any // SHA1
	NumberOfKeys  byte
	SecurityFlags byte
}

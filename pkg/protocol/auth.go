// TODO: xfer packets
package protocol

type AuthCmd byte

const (
	AuthCmdChallenge          AuthCmd = 0x00
	AuthCmdProof              AuthCmd = 0x01
	AuthCmdReconnectChallenge AuthCmd = 0x02
	AuthCmdReconnectProof     AuthCmd = 0x03
	AuthCmdRealmList          AuthCmd = 0x10
	// AuthCmdXferInitiate       AuthCmd = 0x30
	// AuthCmdXferData           AuthCmd = 0x31
	// AuthCmdXferAccept         AuthCmd = 0x32
	// AuthCmdXferResume         AuthCmd = 0x33
	// AuthCmdXferCancel         AuthCmd = 0x34
)

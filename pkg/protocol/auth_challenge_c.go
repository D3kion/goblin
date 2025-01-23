package protocol

import (
	"encoding/binary"
	"net"
)

type C_MSGAuthChallenge struct {
	Opcode   AuthCmd
	Error    uint8 // wtf?
	Size     uint16
	GameName string
	Version1 byte
	Version2 byte
	Version3 byte
	Build    uint16
	Platform string
	Os       string
	Locale   string
	Timezone uint32
	Ip       net.IP
	LoginLen byte
	Login    string
}

func (msg *C_MSGAuthChallenge) Read(buf []byte) (n int, err error) {
	_, err = binary.Decode(buf[n:1], binary.LittleEndian,
		&msg.Opcode)
	if err != nil {
		return
	}
	n++
	_, err = binary.Decode(buf[n:n+1], binary.LittleEndian,
		&msg.Error)
	if err != nil {
		return
	}
	n++
	_, err = binary.Decode(buf[n:n+2], binary.LittleEndian,
		&msg.Size)
	if err != nil {
		return
	}
	n += 2
	gameName := make([]byte, 4)
	_, err = binary.Decode(buf[n:n+4], binary.LittleEndian,
		&gameName)
	if err != nil {
		return
	}
	msg.GameName = bytesToStr(gameName)
	n += 4
	_, err = binary.Decode(buf[n:n+1], binary.LittleEndian,
		&msg.Version1)
	if err != nil {
		return
	}
	n++
	_, err = binary.Decode(buf[n:n+1], binary.LittleEndian,
		&msg.Version2)
	if err != nil {
		return
	}
	n++
	_, err = binary.Decode(buf[n:n+1], binary.LittleEndian,
		&msg.Version3)
	if err != nil {
		return
	}
	n++
	_, err = binary.Decode(buf[n:n+2], binary.LittleEndian,
		&msg.Build)
	if err != nil {
		return
	}
	n += 2
	platform := make([]byte, 4)
	_, err = binary.Decode(buf[n:n+4], binary.LittleEndian,
		&platform)
	if err != nil {
		return
	}
	msg.Platform = bytesToStr(platform)
	n += 4
	os := make([]byte, 4)
	_, err = binary.Decode(buf[n:n+4], binary.LittleEndian,
		&os)
	if err != nil {
		return
	}
	msg.Os = bytesToStr(os)
	n += 4
	locale := make([]byte, 4)
	_, err = binary.Decode(buf[n:n+4], binary.LittleEndian,
		&locale)
	if err != nil {
		return
	}
	msg.Locale = bytesToStr(locale)
	n += 4
	_, err = binary.Decode(buf[n:n+4], binary.LittleEndian,
		&msg.Timezone)
	if err != nil {
		return
	}
	n += 4
	ip := make([]byte, 4)
	_, err = binary.Decode(buf[n:n+4], binary.LittleEndian,
		&ip)
	if err != nil {
		return
	}
	n += 4
	msg.Ip = net.IPv4(ip[0], ip[1], ip[2], ip[3])
	_, err = binary.Decode(buf[n:n+1], binary.LittleEndian,
		&msg.LoginLen)
	if err != nil {
		return
	}
	n++
	login := make([]byte, msg.LoginLen)
	_, err = binary.Decode(buf[n:n+int(msg.LoginLen)], binary.LittleEndian,
		&login)
	if err != nil {
		return
	}
	n += int(msg.LoginLen)
	msg.Login = string(login)
	return
}

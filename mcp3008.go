package mcp3008

import (
	"periph.io/x/conn/v3/physic"
	"periph.io/x/conn/v3/spi"
	"periph.io/x/conn/v3/spi/spireg"
	"periph.io/x/host/v3"
)

type Mcp3008 struct {
	Conn        spi.Conn
	VRef        float64
	Name        string `json:"name"`
	Description string `json:"description"`
}

func New(device string, vRef float64, name string, description string) (Mcp3008, error) {
	mcp3008 := Mcp3008{nil, vRef, name, description}

	if _, err := host.Init(); err != nil {
		return Mcp3008{}, err
	}

	if port, err := spireg.Open(device); err != nil {
		return Mcp3008{}, err
	} else {
		if mcp3008.Conn, err = port.Connect(physic.MegaHertz, spi.Mode0, 8); err != nil {
			return Mcp3008{}, err
		}
	}

	return mcp3008, nil
}

func ReadConversionRegister(module Mcp3008, channel int) float64 {
	if channel > 7 || channel < 0 {
		return -1
	}

	write := []byte{1, byte((8 + channel) << 4), 0}
	read := make([]byte, 3)
	if err := module.Conn.Tx(write, read); err != nil {
		return -1.0
	}
	// Use read.
	return float64((int(read[1])<<8)+int(read[2])) * module.VRef / 1024
}

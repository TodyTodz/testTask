package handler_helpers

import (
	"encoding/xml"
	"os"
)

type HandlerSettings struct {
	HostField      string `xml:"host"`
	PortHTTPField  uint16 `xml:"port-http"`
}

func NewHttpConfig(path string) (*HandlerSettings, error) {
	hSettings := &HandlerSettings{}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }()

	if err = xml.NewDecoder(f).Decode(hSettings); err != nil {
		return nil, err
	}

	return hSettings, nil
}

func (hs HandlerSettings) Host() string {
	return hs.HostField
}

func (hs HandlerSettings) PortHttp() uint16 {
	return hs.PortHTTPField
}

//Package for HTTP settings
package handler_helpers

import (
	"encoding/xml"
	"os"
)

//HandlerSettings - structure for store HTTP settings
type HandlerSettings struct {
	HostField      string `xml:"host"`
	PortHTTPField  uint16 `xml:"port-http"`
}

//NewHttpConfig parse HTTP settings from xml file (see the example in the folder CONFIGS)
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

//The method for getting an IP for the http server.
func (hs HandlerSettings) Host() string {
	return hs.HostField
}

//The method for getting a port for the http server.
func (hs HandlerSettings) PortHttp() uint16 {
	return hs.PortHTTPField
}

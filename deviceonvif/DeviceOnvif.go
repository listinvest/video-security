package deviceonvif

import (
	"encoding/xml"
	"io/ioutil"
	"strings"

	"github.com/yakovlevdmv/goonvif"
	"github.com/yakovlevdmv/goonvif/Media"
	"github.com/yakovlevdmv/goonvif/xsd/onvif"
)

//----------------------------------------------------------------
type AnyNode struct {
	XMLName xml.Name
	Attrs   []xml.Attr `xml:",any,attr"`
	Nodes   []AnyNode  `xml:",any"`
	Content string     `xml:",chardata"`
}

func walkAnyNode(path string, aNode *AnyNode) []AnyNode {
	found := []AnyNode{}
	ss := strings.SplitN(path, "/", 2)
	if aNode.XMLName.Local == ss[0] {
		if len(ss) > 1 {
			for i := range aNode.Nodes {
				found = append(found, walkAnyNode(ss[1], &aNode.Nodes[i])...)
			}
		} else {
			return append(found, *aNode)
		}
	}
	return found
}

//----------------------------------------------------------------

type CallMethodType func(method interface{}, pathToData string) ([]AnyNode, error)
type DeviceOnvif struct {
	 CallMethod CallMethodType 
}

//----------------------------------------------------------------
func (dev *DeviceOnvif) NewDevice(address string, login string, password string) *DeviceOnvif {

	devInternal, _ := goonvif.NewDevice(address)
	devInternal.Authenticate(login, password)

	return &DeviceOnvif{ CallMethod: CallMethodType(func(method interface{}, pathToData string) ([]AnyNode, error) {
		response, err := devInternal.CallMethod(method)
		if err == nil {
			bodyBytes, err := ioutil.ReadAll(response.Body)
			if err == nil {
				var aNode AnyNode
				err := xml.Unmarshal(bodyBytes, &aNode)
				if err == nil {
					found := walkAnyNode(pathToData, &aNode)
					return found, nil
				}
			}
		}
		return make([]AnyNode, 0), err
	})}
}

//----------------------------------------------------------------
func (dev *DeviceOnvif) GetProfiles() []Media.GetProfilesResponse {
	method := Media.GetProfiles{}
	result := make([]Media.GetProfilesResponse, 0)
	var data Media.GetProfilesResponse

	found, _ := dev.CallMethod(method, "Envelope/Body/GetProfilesResponse")

	for i := range found {
		tempBytes, _ := xml.Marshal(found[i])
		xml.Unmarshal(tempBytes, &data)
		result = append(result, data)
	}
	return result
}

//----------------------------------------------------------------
func (dev *DeviceOnvif) GetStreamUri(token onvif.ReferenceToken) []Media.GetStreamUriResponse {
	method := Media.GetStreamUri{ProfileToken: token}
	result := make([]Media.GetStreamUriResponse, 0)
	var data Media.GetStreamUriResponse

	found, _ := dev.CallMethod(method, "Envelope/Body/GetStreamUriResponse")

	for i := range found {
		tempBytes, _ := xml.Marshal(found[i])
		xml.Unmarshal(tempBytes, &data)
		result = append(result, data)
	}
	return result
}

//----------------------------------------------------------------
func (dev *DeviceOnvif) GetSnapshotUri(token onvif.ReferenceToken) []Media.GetSnapshotUriResponse {
	method := Media.GetSnapshotUri{ProfileToken: token}
	result := make([]Media.GetSnapshotUriResponse, 0)
	var data Media.GetSnapshotUriResponse

	found, _ := dev.CallMethod(method, "Envelope/Body/GetSnapshotUriResponse")

	for i := range found {
		tempBytes, _ := xml.Marshal(found[i])
		xml.Unmarshal(tempBytes, &data)
		result = append(result, data)
	}
	return result
}

//----------------------------------------------------------------

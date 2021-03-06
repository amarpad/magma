// Code generated by go-swagger; DO NOT EDIT.

package wifi_meshes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// GetWifiNetworkIDMeshesReader is a Reader for the GetWifiNetworkIDMeshes structure.
type GetWifiNetworkIDMeshesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWifiNetworkIDMeshesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWifiNetworkIDMeshesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetWifiNetworkIDMeshesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWifiNetworkIDMeshesOK creates a GetWifiNetworkIDMeshesOK with default headers values
func NewGetWifiNetworkIDMeshesOK() *GetWifiNetworkIDMeshesOK {
	return &GetWifiNetworkIDMeshesOK{}
}

/*GetWifiNetworkIDMeshesOK handles this case with default header values.

List of meshes in the network
*/
type GetWifiNetworkIDMeshesOK struct {
	Payload []models.MeshID
}

func (o *GetWifiNetworkIDMeshesOK) Error() string {
	return fmt.Sprintf("[GET /wifi/{network_id}/meshes][%d] getWifiNetworkIdMeshesOK  %+v", 200, o.Payload)
}

func (o *GetWifiNetworkIDMeshesOK) GetPayload() []models.MeshID {
	return o.Payload
}

func (o *GetWifiNetworkIDMeshesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWifiNetworkIDMeshesDefault creates a GetWifiNetworkIDMeshesDefault with default headers values
func NewGetWifiNetworkIDMeshesDefault(code int) *GetWifiNetworkIDMeshesDefault {
	return &GetWifiNetworkIDMeshesDefault{
		_statusCode: code,
	}
}

/*GetWifiNetworkIDMeshesDefault handles this case with default header values.

Unexpected Error
*/
type GetWifiNetworkIDMeshesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get wifi network ID meshes default response
func (o *GetWifiNetworkIDMeshesDefault) Code() int {
	return o._statusCode
}

func (o *GetWifiNetworkIDMeshesDefault) Error() string {
	return fmt.Sprintf("[GET /wifi/{network_id}/meshes][%d] GetWifiNetworkIDMeshes default  %+v", o._statusCode, o.Payload)
}

func (o *GetWifiNetworkIDMeshesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetWifiNetworkIDMeshesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

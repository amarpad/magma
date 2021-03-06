// Code generated by go-swagger; DO NOT EDIT.

package subscribers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// GetLTENetworkIDMsisdnsReader is a Reader for the GetLTENetworkIDMsisdns structure.
type GetLTENetworkIDMsisdnsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLTENetworkIDMsisdnsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLTENetworkIDMsisdnsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLTENetworkIDMsisdnsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLTENetworkIDMsisdnsOK creates a GetLTENetworkIDMsisdnsOK with default headers values
func NewGetLTENetworkIDMsisdnsOK() *GetLTENetworkIDMsisdnsOK {
	return &GetLTENetworkIDMsisdnsOK{}
}

/*GetLTENetworkIDMsisdnsOK handles this case with default header values.

List of all MSISDNS in the network, mapped to their subscriber ID
*/
type GetLTENetworkIDMsisdnsOK struct {
	Payload map[string]models.SubscriberID
}

func (o *GetLTENetworkIDMsisdnsOK) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/msisdns][%d] getLteNetworkIdMsisdnsOK  %+v", 200, o.Payload)
}

func (o *GetLTENetworkIDMsisdnsOK) GetPayload() map[string]models.SubscriberID {
	return o.Payload
}

func (o *GetLTENetworkIDMsisdnsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLTENetworkIDMsisdnsDefault creates a GetLTENetworkIDMsisdnsDefault with default headers values
func NewGetLTENetworkIDMsisdnsDefault(code int) *GetLTENetworkIDMsisdnsDefault {
	return &GetLTENetworkIDMsisdnsDefault{
		_statusCode: code,
	}
}

/*GetLTENetworkIDMsisdnsDefault handles this case with default header values.

Unexpected Error
*/
type GetLTENetworkIDMsisdnsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get LTE network ID msisdns default response
func (o *GetLTENetworkIDMsisdnsDefault) Code() int {
	return o._statusCode
}

func (o *GetLTENetworkIDMsisdnsDefault) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/msisdns][%d] GetLTENetworkIDMsisdns default  %+v", o._statusCode, o.Payload)
}

func (o *GetLTENetworkIDMsisdnsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLTENetworkIDMsisdnsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

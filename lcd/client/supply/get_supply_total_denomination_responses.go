// Code generated by go-swagger; DO NOT EDIT.

package supply

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetSupplyTotalDenominationReader is a Reader for the GetSupplyTotalDenomination structure.
type GetSupplyTotalDenominationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSupplyTotalDenominationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSupplyTotalDenominationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSupplyTotalDenominationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSupplyTotalDenominationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSupplyTotalDenominationOK creates a GetSupplyTotalDenominationOK with default headers values
func NewGetSupplyTotalDenominationOK() *GetSupplyTotalDenominationOK {
	return &GetSupplyTotalDenominationOK{}
}

/*GetSupplyTotalDenominationOK handles this case with default header values.

OK
*/
type GetSupplyTotalDenominationOK struct {
	Payload string
}

func (o *GetSupplyTotalDenominationOK) Error() string {
	return fmt.Sprintf("[GET /supply/total/{denomination}][%d] getSupplyTotalDenominationOK  %+v", 200, o.Payload)
}

func (o *GetSupplyTotalDenominationOK) GetPayload() string {
	return o.Payload
}

func (o *GetSupplyTotalDenominationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSupplyTotalDenominationBadRequest creates a GetSupplyTotalDenominationBadRequest with default headers values
func NewGetSupplyTotalDenominationBadRequest() *GetSupplyTotalDenominationBadRequest {
	return &GetSupplyTotalDenominationBadRequest{}
}

/*GetSupplyTotalDenominationBadRequest handles this case with default header values.

Invalid coin denomination
*/
type GetSupplyTotalDenominationBadRequest struct {
}

func (o *GetSupplyTotalDenominationBadRequest) Error() string {
	return fmt.Sprintf("[GET /supply/total/{denomination}][%d] getSupplyTotalDenominationBadRequest ", 400)
}

func (o *GetSupplyTotalDenominationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSupplyTotalDenominationInternalServerError creates a GetSupplyTotalDenominationInternalServerError with default headers values
func NewGetSupplyTotalDenominationInternalServerError() *GetSupplyTotalDenominationInternalServerError {
	return &GetSupplyTotalDenominationInternalServerError{}
}

/*GetSupplyTotalDenominationInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetSupplyTotalDenominationInternalServerError struct {
}

func (o *GetSupplyTotalDenominationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /supply/total/{denomination}][%d] getSupplyTotalDenominationInternalServerError ", 500)
}

func (o *GetSupplyTotalDenominationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
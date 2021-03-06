///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package drivers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// GetDriverTypeReader is a Reader for the GetDriverType structure.
type GetDriverTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDriverTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDriverTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetDriverTypeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetDriverTypeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetDriverTypeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetDriverTypeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDriverTypeOK creates a GetDriverTypeOK with default headers values
func NewGetDriverTypeOK() *GetDriverTypeOK {
	return &GetDriverTypeOK{}
}

/*GetDriverTypeOK handles this case with default header values.

Successful operation
*/
type GetDriverTypeOK struct {
	Payload *v1.EventDriverType
}

func (o *GetDriverTypeOK) Error() string {
	return fmt.Sprintf("[GET /drivertypes/{driverTypeName}][%d] getDriverTypeOK  %+v", 200, o.Payload)
}

func (o *GetDriverTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.EventDriverType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDriverTypeBadRequest creates a GetDriverTypeBadRequest with default headers values
func NewGetDriverTypeBadRequest() *GetDriverTypeBadRequest {
	return &GetDriverTypeBadRequest{}
}

/*GetDriverTypeBadRequest handles this case with default header values.

Invalid Name supplied
*/
type GetDriverTypeBadRequest struct {
	Payload *v1.Error
}

func (o *GetDriverTypeBadRequest) Error() string {
	return fmt.Sprintf("[GET /drivertypes/{driverTypeName}][%d] getDriverTypeBadRequest  %+v", 400, o.Payload)
}

func (o *GetDriverTypeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDriverTypeNotFound creates a GetDriverTypeNotFound with default headers values
func NewGetDriverTypeNotFound() *GetDriverTypeNotFound {
	return &GetDriverTypeNotFound{}
}

/*GetDriverTypeNotFound handles this case with default header values.

Driver not found
*/
type GetDriverTypeNotFound struct {
	Payload *v1.Error
}

func (o *GetDriverTypeNotFound) Error() string {
	return fmt.Sprintf("[GET /drivertypes/{driverTypeName}][%d] getDriverTypeNotFound  %+v", 404, o.Payload)
}

func (o *GetDriverTypeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDriverTypeInternalServerError creates a GetDriverTypeInternalServerError with default headers values
func NewGetDriverTypeInternalServerError() *GetDriverTypeInternalServerError {
	return &GetDriverTypeInternalServerError{}
}

/*GetDriverTypeInternalServerError handles this case with default header values.

Internal server error
*/
type GetDriverTypeInternalServerError struct {
	Payload *v1.Error
}

func (o *GetDriverTypeInternalServerError) Error() string {
	return fmt.Sprintf("[GET /drivertypes/{driverTypeName}][%d] getDriverTypeInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDriverTypeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDriverTypeDefault creates a GetDriverTypeDefault with default headers values
func NewGetDriverTypeDefault(code int) *GetDriverTypeDefault {
	return &GetDriverTypeDefault{
		_statusCode: code,
	}
}

/*GetDriverTypeDefault handles this case with default header values.

Unknown error
*/
type GetDriverTypeDefault struct {
	_statusCode int

	Payload *v1.Error
}

// Code gets the status code for the get driver type default response
func (o *GetDriverTypeDefault) Code() int {
	return o._statusCode
}

func (o *GetDriverTypeDefault) Error() string {
	return fmt.Sprintf("[GET /drivertypes/{driverTypeName}][%d] getDriverType default  %+v", o._statusCode, o.Payload)
}

func (o *GetDriverTypeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

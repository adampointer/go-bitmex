// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/adampointer/go-bitmex/swagger/models"
)

// UserGetWalletReader is a Reader for the UserGetWallet structure.
type UserGetWalletReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserGetWalletReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserGetWalletOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserGetWalletOK creates a UserGetWalletOK with default headers values
func NewUserGetWalletOK() *UserGetWalletOK {
	return &UserGetWalletOK{}
}

/*UserGetWalletOK handles this case with default header values.

Request was successful
*/
type UserGetWalletOK struct {
	Payload *models.Wallet
}

func (o *UserGetWalletOK) Error() string {
	return fmt.Sprintf("[GET /user/wallet][%d] userGetWalletOK  %+v", 200, o.Payload)
}

func (o *UserGetWalletOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Wallet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
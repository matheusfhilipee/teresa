package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/luizalabs/teresa-api/models"
)

// AddUserToTeamReader is a Reader for the AddUserToTeam structure.
type AddUserToTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *AddUserToTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddUserToTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewAddUserToTeamDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewAddUserToTeamOK creates a AddUserToTeamOK with default headers values
func NewAddUserToTeamOK() *AddUserToTeamOK {
	return &AddUserToTeamOK{}
}

/*AddUserToTeamOK handles this case with default header values.

Team details
*/
type AddUserToTeamOK struct {
	Payload *models.Team
}

func (o *AddUserToTeamOK) Error() string {
	return fmt.Sprintf("[POST /teams/{team_name}/users][%d] addUserToTeamOK  %+v", 200, o.Payload)
}

func (o *AddUserToTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Team)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserToTeamDefault creates a AddUserToTeamDefault with default headers values
func NewAddUserToTeamDefault(code int) *AddUserToTeamDefault {
	return &AddUserToTeamDefault{
		_statusCode: code,
	}
}

/*AddUserToTeamDefault handles this case with default header values.

Error
*/
type AddUserToTeamDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the add user to team default response
func (o *AddUserToTeamDefault) Code() int {
	return o._statusCode
}

func (o *AddUserToTeamDefault) Error() string {
	return fmt.Sprintf("[POST /teams/{team_name}/users][%d] addUserToTeam default  %+v", o._statusCode, o.Payload)
}

func (o *AddUserToTeamDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddUserToTeamBody add user to team body

swagger:model AddUserToTeamBody
*/
type AddUserToTeamBody struct {

	/* email

	Required: true
	Min Length: 1
	*/
	Email *strfmt.Email `json:"email"`
}

// Validate validates this add user to team body
func (o *AddUserToTeamBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddUserToTeamBody) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("user"+"."+"email", "body", o.Email); err != nil {
		return err
	}

	if err := validate.MinLength("user"+"."+"email", "body", string(*o.Email), 1); err != nil {
		return err
	}

	return nil
}

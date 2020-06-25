// Code generated by go-swagger; DO NOT EDIT.

package foods

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"food_app/food_app/models"
)

// UpdateFoodOKCode is the HTTP code returned for type UpdateFoodOK
const UpdateFoodOKCode int = 200

/*UpdateFoodOK update a food

swagger:response updateFoodOK
*/
type UpdateFoodOK struct {

	/*
	  In: Body
	*/
	Payload *models.Food `json:"body,omitempty"`
}

// NewUpdateFoodOK creates UpdateFoodOK with default headers values
func NewUpdateFoodOK() *UpdateFoodOK {

	return &UpdateFoodOK{}
}

// WithPayload adds the payload to the update food o k response
func (o *UpdateFoodOK) WithPayload(payload *models.Food) *UpdateFoodOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update food o k response
func (o *UpdateFoodOK) SetPayload(payload *models.Food) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateFoodOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
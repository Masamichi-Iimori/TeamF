// Code generated by go-swagger; DO NOT EDIT.

package scenepicks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostTagHandlerFunc turns a function with the right signature into a post tag handler
type PostTagHandlerFunc func(PostTagParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostTagHandlerFunc) Handle(params PostTagParams) middleware.Responder {
	return fn(params)
}

// PostTagHandler interface for that can handle valid post tag params
type PostTagHandler interface {
	Handle(PostTagParams) middleware.Responder
}

// NewPostTag creates a new http.Handler for the post tag operation
func NewPostTag(ctx *middleware.Context, handler PostTagHandler) *PostTag {
	return &PostTag{Context: ctx, Handler: handler}
}

/*PostTag swagger:route POST /tag postTag

PostTag post tag API

*/
type PostTag struct {
	Context *middleware.Context
	Handler PostTagHandler
}

func (o *PostTag) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostTagParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostTagBody post tag body
//
// swagger:model PostTagBody
type PostTagBody struct {

	// name
	Name string `json:"name,omitempty"`

	// “title“, “author“, “speaker“ “other“のうちのどれか。typeによる絞り込み。空欄だと全部。
	Type string `json:"type,omitempty"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (o *PostTagBody) UnmarshalJSON(data []byte) error {
	var props struct {

		// name
		Name string `json:"name,omitempty"`

		// “title“, “author“, “speaker“ “other“のうちのどれか。typeによる絞り込み。空欄だと全部。
		Type string `json:"type,omitempty"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	o.Name = props.Name
	o.Type = props.Type
	return nil
}

// Validate validates this post tag body
func (o *PostTagBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostTagBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostTagBody) UnmarshalBinary(b []byte) error {
	var res PostTagBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package logins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetRsaPubKeyHandlerFunc turns a function with the right signature into a get rsa pub key handler
type GetRsaPubKeyHandlerFunc func(GetRsaPubKeyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRsaPubKeyHandlerFunc) Handle(params GetRsaPubKeyParams) middleware.Responder {
	return fn(params)
}

// GetRsaPubKeyHandler interface for that can handle valid get rsa pub key params
type GetRsaPubKeyHandler interface {
	Handle(GetRsaPubKeyParams) middleware.Responder
}

// NewGetRsaPubKey creates a new http.Handler for the get rsa pub key operation
func NewGetRsaPubKey(ctx *middleware.Context, handler GetRsaPubKeyHandler) *GetRsaPubKey {
	return &GetRsaPubKey{Context: ctx, Handler: handler}
}

/*GetRsaPubKey swagger:route GET /cc/v1/getrsapubkey Logins getRsaPubKey

rsa pubkey

get LDAP login's rsa public key

*/
type GetRsaPubKey struct {
	Context *middleware.Context
	Handler GetRsaPubKeyHandler
}

func (o *GetRsaPubKey) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRsaPubKeyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

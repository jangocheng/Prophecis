// Code generated by go-swagger; DO NOT EDIT.

package storages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteStorageByIDHandlerFunc turns a function with the right signature into a delete storage by Id handler
type DeleteStorageByIDHandlerFunc func(DeleteStorageByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteStorageByIDHandlerFunc) Handle(params DeleteStorageByIDParams) middleware.Responder {
	return fn(params)
}

// DeleteStorageByIDHandler interface for that can handle valid delete storage by Id params
type DeleteStorageByIDHandler interface {
	Handle(DeleteStorageByIDParams) middleware.Responder
}

// NewDeleteStorageByID creates a new http.Handler for the delete storage by Id operation
func NewDeleteStorageByID(ctx *middleware.Context, handler DeleteStorageByIDHandler) *DeleteStorageByID {
	return &DeleteStorageByID{Context: ctx, Handler: handler}
}

/*DeleteStorageByID swagger:route DELETE /cc/v1/storages/id/{storageId} Storages deleteStorageById

Returns a Storage.

Optional extended description in Markdown.

*/
type DeleteStorageByID struct {
	Context *middleware.Context
	Handler DeleteStorageByIDHandler
}

func (o *DeleteStorageByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteStorageByIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
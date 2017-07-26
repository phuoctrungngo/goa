// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// storage HTTP server
//
// Command:
// $ goa gen goa.design/goa.v2/examples/cellar/design

package server

import (
	"net/http"

	goa "goa.design/goa.v2"
	"goa.design/goa.v2/examples/cellar/gen/storage"
	goahttp "goa.design/goa.v2/http"
)

// Server lists the storage service endpoint HTTP handlers.
type Server struct {
	Add    http.Handler
	List   http.Handler
	Show   http.Handler
	Remove http.Handler
}

// New instantiates HTTP handlers for all the storage service endpoints.
func New(
	e *storage.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(http.ResponseWriter, *http.Request) (goahttp.Encoder, string),
) *Server {
	return &Server{
		Add:    NewAddHandler(e.Add, mux, dec, enc),
		List:   NewListHandler(e.List, mux, dec, enc),
		Show:   NewShowHandler(e.Show, mux, dec, enc),
		Remove: NewRemoveHandler(e.Remove, mux, dec, enc),
	}
}

// Mount configures the mux to serve the storage endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountAddHandler(mux, h.Add)
	MountListHandler(mux, h.List)
	MountShowHandler(mux, h.Show)
	MountRemoveHandler(mux, h.Remove)
}

// MountAddHandler configures the mux to serve the "storage" service "add"
// endpoint.
func MountAddHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/storage", f)
}

// NewAddHandler creates a HTTP handler which loads the HTTP request and calls
// the "storage" service "add" endpoint.
func NewAddHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(http.ResponseWriter, *http.Request) (goahttp.Encoder, string),
) http.Handler {
	var (
		decodeRequest  = DecodeAddRequest(mux, dec)
		encodeResponse = EncodeAddResponse(enc)
		encodeError    = goahttp.EncodeError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload, err := decodeRequest(r)
		if err != nil {
			encodeError(w, r, err)
			return
		}

		res, err := endpoint(r.Context(), payload)

		if err != nil {
			encodeError(w, r, err)
			return
		}
		if err := encodeResponse(w, r, res); err != nil {
			encodeError(w, r, err)
		}
	})
}

// MountListHandler configures the mux to serve the "storage" service "list"
// endpoint.
func MountListHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/storage", f)
}

// NewListHandler creates a HTTP handler which loads the HTTP request and calls
// the "storage" service "list" endpoint.
func NewListHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(http.ResponseWriter, *http.Request) (goahttp.Encoder, string),
) http.Handler {
	var (
		encodeResponse = EncodeListResponse(enc)
		encodeError    = goahttp.EncodeError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res, err := endpoint(r.Context(), nil)

		if err != nil {
			encodeError(w, r, err)
			return
		}
		if err := encodeResponse(w, r, res); err != nil {
			encodeError(w, r, err)
		}
	})
}

// MountShowHandler configures the mux to serve the "storage" service "show"
// endpoint.
func MountShowHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/storage/{id}", f)
}

// NewShowHandler creates a HTTP handler which loads the HTTP request and calls
// the "storage" service "show" endpoint.
func NewShowHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(http.ResponseWriter, *http.Request) (goahttp.Encoder, string),
) http.Handler {
	var (
		decodeRequest  = DecodeShowRequest(mux, dec)
		encodeResponse = EncodeShowResponse(enc)
		encodeError    = EncodeShowError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload, err := decodeRequest(r)
		if err != nil {
			encodeError(w, r, err)
			return
		}

		res, err := endpoint(r.Context(), payload)

		if err != nil {
			encodeError(w, r, err)
			return
		}
		if err := encodeResponse(w, r, res); err != nil {
			encodeError(w, r, err)
		}
	})
}

// MountRemoveHandler configures the mux to serve the "storage" service
// "remove" endpoint.
func MountRemoveHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("DELETE", "/storage/{id}", f)
}

// NewRemoveHandler creates a HTTP handler which loads the HTTP request and
// calls the "storage" service "remove" endpoint.
func NewRemoveHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(http.ResponseWriter, *http.Request) (goahttp.Encoder, string),
) http.Handler {
	var (
		decodeRequest  = DecodeRemoveRequest(mux, dec)
		encodeResponse = EncodeRemoveResponse(enc)
		encodeError    = goahttp.EncodeError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload, err := decodeRequest(r)
		if err != nil {
			encodeError(w, r, err)
			return
		}

		res, err := endpoint(r.Context(), payload)

		if err != nil {
			encodeError(w, r, err)
			return
		}
		if err := encodeResponse(w, r, res); err != nil {
			encodeError(w, r, err)
		}
	})
}

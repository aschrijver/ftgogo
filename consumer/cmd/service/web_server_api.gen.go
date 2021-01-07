// Package main provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi"
)

// ConsumerID defines model for ConsumerID.
type ConsumerID string

// ConsumerIDResponse defines model for ConsumerIDResponse.
type ConsumerIDResponse struct {
	Id string `json:"id"`
}

// ConsumerResponse defines model for ConsumerResponse.
type ConsumerResponse struct {
	ConsumerId string `json:"consumer_id"`
	Name       string `json:"name"`
}

// RegisterConsumerJSONBody defines parameters for RegisterConsumer.
type RegisterConsumerJSONBody struct {
	Name string `json:"name"`
}

// RegisterConsumerRequestBody defines body for RegisterConsumer for application/json ContentType.
type RegisterConsumerJSONRequestBody RegisterConsumerJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Register a new consumer for delivery
	// (POST /consumers)
	RegisterConsumer(w http.ResponseWriter, r *http.Request)
	// Get basic information for an consumer by ID
	// (GET /consumers/{consumerID})
	GetConsumer(w http.ResponseWriter, r *http.Request, consumerID ConsumerID)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// RegisterConsumer operation middleware
func (siw *ServerInterfaceWrapper) RegisterConsumer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	siw.Handler.RegisterConsumer(w, r.WithContext(ctx))
}

// GetConsumer operation middleware
func (siw *ServerInterfaceWrapper) GetConsumer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "consumerID" -------------
	var consumerID ConsumerID

	err = runtime.BindStyledParameter("simple", false, "consumerID", chi.URLParam(r, "consumerID"), &consumerID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter consumerID: %s", err), http.StatusBadRequest)
		return
	}

	siw.Handler.GetConsumer(w, r.WithContext(ctx), consumerID)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerFromMux(si, chi.NewRouter())
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerFromMuxWithBaseURL(si, r, "")
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	r.Group(func(r chi.Router) {
		r.Post(baseURL+"/consumers", wrapper.RegisterConsumer)
	})
	r.Group(func(r chi.Router) {
		r.Get(baseURL+"/consumers/{consumerID}", wrapper.GetConsumer)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7xUUW/bOAz+KwLvHu4AN0p7Lwe/besWBHvo0PWtKAZFZmJ1tqRRdIog8H8f6Nix22ZY",
	"gA0DgiSiRPLj95Hcgw11DB49J8j3EA2ZGhmpO70LPjU10vJaTgUmSy6yCx5ysP2dcoX6p2lc8S9k4OQm",
	"Gi4hA29qnLxbXkMGhN8aR1hAztRgBsmWWBsJvg5UG4YcJBRkwLso3onJ+Q20bSvOKQaf8AW0294sVhs8",
	"o2f5a2KsnDWCVj8mgbyfpIsUIhK7QzBXyPfLlFO49/Lm4QgrrB7R8gHWc1puPkKbHdH9BmwDf19Oghxo",
	"/hn6aZTe59xqxOr8Ogw1GNvV0Mv74W5xs7hRy5QaTJBBQxXkUDLHlGu9cVw2q5kNtU5s7Ncm6TVvwiZo",
	"d3B4lfGuRDWAvUhIW2dRvfm0VC6paIhVWCsuUfV5C6yDT0wdl2rCq7SQ40ogDmKoz2M4yGCLlA45L2fz",
	"2VyghIjeRAc5/NeZsq6XOxn0AOowJiF1JIhQXbplATnc4sYlRhoS9g2Pid+GYvcLPXCexmerSj3OI9Gq",
	"h/lqQl9O3dX8Un7+JlxDDn/pcXno4zt9YjQ7CGvTVDxxn8304ZNKQ1hcbIJ+wpVOEe1sZ+rqB+HfE4Vx",
	"sKS61NS1od1EAWWUx6exvnUgVWDltki7zmWUU+/HDdUKug2ekHaBPFF1uiXvT9MxPpnwAe3DK0Ln5xP6",
	"x+lcIKuVSc4qWQCyoGXKhEzjR3JXOyWlda5I24GV55tgX4bEbb6PgbjVMmQZbA05s6oOTJT9SB0LgypY",
	"U3XmNgPxe379/3x+JVkf2u8BAAD//xbgrwrFBgAA",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}

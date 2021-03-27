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
	"github.com/go-chi/chi/v5"
	externalRef0 "github.com/stackus/ftgogo/delivery/internal/application/queries"
)

// CourierID defines model for CourierID.
type CourierID string

// DeliveryID defines model for DeliveryID.
type DeliveryID string

// CourierAvailableResponse defines model for CourierAvailableResponse.
type CourierAvailableResponse struct {
	Available bool `json:"available"`
}

// DeliveryStatusResponse defines model for DeliveryStatusResponse.
type DeliveryStatusResponse externalRef0.DeliveryStatus

// SetCourierAvailabilityJSONBody defines parameters for SetCourierAvailability.
type SetCourierAvailabilityJSONBody struct {
	Available bool `json:"available"`
}

// SetCourierAvailabilityRequestBody defines body for SetCourierAvailability for application/json ContentType.
type SetCourierAvailabilityJSONRequestBody SetCourierAvailabilityJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Set availability for new or existing couriers
	// (POST /couriers/{courierID}/availability)
	SetCourierAvailability(w http.ResponseWriter, r *http.Request, courierID CourierID)
	// Get delivery status by ID
	// (GET /deliveries/{deliveryID})
	GetDeliveryStatus(w http.ResponseWriter, r *http.Request, deliveryID DeliveryID)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// SetCourierAvailability operation middleware
func (siw *ServerInterfaceWrapper) SetCourierAvailability(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "courierID" -------------
	var courierID CourierID

	err = runtime.BindStyledParameter("simple", false, "courierID", chi.URLParam(r, "courierID"), &courierID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter courierID: %s", err), http.StatusBadRequest)
		return
	}

	siw.Handler.SetCourierAvailability(w, r.WithContext(ctx), courierID)
}

// GetDeliveryStatus operation middleware
func (siw *ServerInterfaceWrapper) GetDeliveryStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "deliveryID" -------------
	var deliveryID DeliveryID

	err = runtime.BindStyledParameter("simple", false, "deliveryID", chi.URLParam(r, "deliveryID"), &deliveryID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter deliveryID: %s", err), http.StatusBadRequest)
		return
	}

	siw.Handler.GetDeliveryStatus(w, r.WithContext(ctx), deliveryID)
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
		r.Post(baseURL+"/couriers/{courierID}/availability", wrapper.SetCourierAvailability)
	})
	r.Group(func(r chi.Router) {
		r.Get(baseURL+"/deliveries/{deliveryID}", wrapper.GetDeliveryStatus)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8xUTW/bOBD9K8TsHnYBRXQ2l4VuadMaRg8pmtyCHGhpbDGVSGY4cmoY+u8FKVkfSYyg",
	"LVD0JnG+Hh/fmwPktnbWoGEP2QGcIlUjI8W/97YhjbS6Cj8F+py0Y20NZJB3IaEL8U/T6OJfSECHgFNc",
	"QgJG1Timra4gAcLHRhMWkDE1mIDPS6xVaL2xVCuGDEInSID3LhR7Jm220LYJXGGld0j716AUfewtLMXY",
	"41fAtKHYO2s8Tkm63CldqXWFX/pgiOXWMBoOn8q5SucqYJYPPgA/TIY6sg6JdddSHXuFn37+2toKlYFu",
	"/hH83ST3fsBq1w+Yc4d1ztX1J5iwecOKG/9TgP8m3EAGaSrTVGrDSEZVclrz2CBp9NI7zNO9qqu/5Kg1",
	"2XXycg7lFORwqs3GHiGqPELsH/bj7fJ6eS1W3jfoIYGGKsigZHY+k3KruWzWaW5r6VnlXxsvN7y1Wyt1",
	"V/Bi4m2J4iiWM4+00zmKy88rob1wiljYjeASRT+3wNoazxSvLSYUBPFoDq84MC5uxnaQwA7JdzPP00W6",
	"CFCsQ6Ochgwu4lESVRxlIXs3eXkYfNXKXgC60ryPUrI+khMEFWGsCsjgBnku1C4/mTn+bnjX2VuNKXLc",
	"CO19p0P0/M4W+z9S7cclNaVI9KBf7IDnvv5vcT4R+oyQIU+eNH/EslFNxS/ckqbSl4qwONta+YTrEw4Z",
	"h3wgsjR2bhPwTV0r2nfvOr/expIw+CQsCfymPWuzFUfdxFrZSzt48zDuxDbg3OIr0lkiP3Ppj6pmsr07",
	"2cxoXrxN84mF9btIXiIPC0H4iEGs9yJcJyYi7Y5MzHfPobSe2+zgLHErg60T2CnSQSrx9mVv1uEaUNlc",
	"VfG4TSDUzcP/LxYXYep9+z0AAP//8s8MArsHAAA=",
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

// Code generated by thriftrw-plugin-yarpc
// @generated

// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package echoserver

import (
	context "context"
	json "encoding/json"
	wire "go.uber.org/thriftrw/wire"
	transport "go.uber.org/yarpc/api/transport"
	thrift "go.uber.org/yarpc/encoding/thrift"
	echo "go.uber.org/yarpc/internal/crossdock/thrift/echo"
	yarpcerrors "go.uber.org/yarpc/yarpcerrors"
)

// Interface is the server-side interface for the Echo service.
type Interface interface {
	Echo(
		ctx context.Context,
		Ping *echo.Ping,
	) (*echo.Pong, error)
}

// New prepares an implementation of the Echo service for
// registration.
//
// 	handler := EchoHandler{}
// 	dispatcher.Register(echoserver.New(handler))
func New(impl Interface, opts ...thrift.RegisterOption) []transport.Procedure {
	h := handler{impl}
	service := thrift.Service{
		Name: "Echo",
		Methods: []thrift.Method{

			thrift.Method{
				Name: "echo",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.Echo),
				},
				Signature:    "Echo(Ping *echo.Ping) (*echo.Pong)",
				ThriftModule: echo.ThriftModule,
			},
		},
	}

	procedures := make([]transport.Procedure, 0, 1)
	procedures = append(procedures, thrift.BuildProcedures(service, opts...)...)
	return procedures
}

type handler struct{ impl Interface }

type errorNamer interface{ ErrorName() string }

type yarpcErrorCodeExtractor interface{ YARPCCode() *yarpcerrors.Code }

func (h handler) Echo(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args echo.Echo_Echo_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode Thrift request for service 'Echo' procedure 'Echo': %w", err)
	}

	success, appErr := h.impl.Echo(ctx, args.Ping)

	hadError := appErr != nil
	result, err := echo.Echo_Echo_Helper.WrapResponse(success, appErr)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
		if namer, ok := appErr.(errorNamer); ok {
			response.ApplicationErrorName = namer.ErrorName()
		}
		if extractor, ok := appErr.(yarpcErrorCodeExtractor); ok {
			response.ApplicationErrorCode = extractor.YARPCCode()
		}
		response.ApplicationError = appErr
	}

	return response, err
}

type stringifier struct{}

// Stringifier returns a thrift.Stringifier capable of stringifying requests
// and responses for the Echo service.
func Stringifier() thrift.Stringifier {
	return &stringifier{}
}

// GetService gets the name of the service for which this stringifier can stringify.
func (s *stringifier) GetService() string {
	return "Echo"
}

// StringifyRequest returns a json string representing the request.
func (s *stringifier) StringifyRequest(procedure string, requestBody wire.Value) (string, error) {
	switch procedure {

	case "Echo":
		var args echo.Echo_Echo_Args
		if err := args.FromWire(requestBody); err != nil {
			return "", err
		}
		b, err := json.Marshal(args)
		if err != nil {
			return "", err
		}
		return string(b), nil

	default:
		return "", yarpcerrors.InvalidArgumentErrorf(
			"could not stringify Thrift request for service 'Echo' procedure '%s'", procedure)
	}
}

// StringifyResponse returns a json string representing the response.
func (s *stringifier) StringifyResponse(procedure string, responseBody wire.Value) (string, error) {
	switch procedure {

	case "Echo":
		var args echo.Echo_Echo_Result
		if err := args.FromWire(responseBody); err != nil {
			return "", err
		}
		b, err := json.Marshal(args)
		if err != nil {
			return "", err
		}
		return string(b), nil

	default:
		return "", yarpcerrors.InvalidArgumentErrorf(
			"could not stringify Thrift request for service 'Echo' procedure '%s'", procedure)
	}
}

// Code generated by thriftrw-plugin-yarpc
// @generated

package nameserver

import (
	context "context"
	json "encoding/json"
	wire "go.uber.org/thriftrw/wire"
	transport "go.uber.org/yarpc/api/transport"
	thrift "go.uber.org/yarpc/encoding/thrift"
	extends "go.uber.org/yarpc/encoding/thrift/thriftrw-plugin-yarpc/internal/tests/extends"
	yarpcerrors "go.uber.org/yarpc/yarpcerrors"
)

// Interface is the server-side interface for the Name service.
type Interface interface {
	Name(
		ctx context.Context,
	) (string, error)
}

// New prepares an implementation of the Name service for
// registration.
//
// 	handler := NameHandler{}
// 	dispatcher.Register(nameserver.New(handler))
func New(impl Interface, opts ...thrift.RegisterOption) []transport.Procedure {
	h := handler{impl}
	service := thrift.Service{
		Name: "Name",
		Methods: []thrift.Method{

			thrift.Method{
				Name: "name",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.Name),
				},
				Signature:    "Name() (string)",
				ThriftModule: extends.ThriftModule,
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

func (h handler) Name(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args extends.Name_Name_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode Thrift request for service 'Name' procedure 'Name': %w", err)
	}

	success, appErr := h.impl.Name(ctx)

	hadError := appErr != nil
	result, err := extends.Name_Name_Helper.WrapResponse(success, appErr)

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

type jsonifier struct{}

// JSONifier returns a thrift.JSONifier capable of producing JSON
// representations of requests and responses for the Name service.
func JSONifier() thrift.JSONifier {
	return &jsonifier{}
}

// GetService gets the name of the service for which this JSONifier can produce
// JSON representations of requests and responses.
func (s *jsonifier) GetService() string {
	return "Name"
}

// RequestToJSON returns a json representation of the request.
func (s *jsonifier) RequestToJSON(procedure string, requestBody wire.Value) ([]byte, error) {
	switch procedure {

	case "Name":
		var args extends.Name_Name_Args
		if err := args.FromWire(requestBody); err != nil {
			return nil, err
		}
		return json.Marshal(args)

	default:
		return nil, yarpcerrors.InvalidArgumentErrorf(
			"could not produce JSON representation of Thrift request for service 'Name' procedure '%s'", procedure)
	}
}

// ResponseToJSON returns a json representation of the response.
func (s *jsonifier) ResponseToJSON(procedure string, responseBody wire.Value) ([]byte, error) {
	switch procedure {

	case "Name":
		var args extends.Name_Name_Result
		if err := args.FromWire(responseBody); err != nil {
			return nil, err
		}
		return json.Marshal(args)

	default:
		return nil, yarpcerrors.InvalidArgumentErrorf(
			"could not produce JSON representation of Thrift response for service 'Name' procedure '%s'", procedure)
	}
}

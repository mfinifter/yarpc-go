// Code generated by thriftrw-plugin-yarpc
// @generated

package storeserver

import (
	context "context"
	json "encoding/json"
	wire "go.uber.org/thriftrw/wire"
	transport "go.uber.org/yarpc/api/transport"
	thrift "go.uber.org/yarpc/encoding/thrift"
	atomic "go.uber.org/yarpc/encoding/thrift/thriftrw-plugin-yarpc/internal/tests/atomic"
	readonlystoreserver "go.uber.org/yarpc/encoding/thrift/thriftrw-plugin-yarpc/internal/tests/atomic/readonlystoreserver"
	yarpcerrors "go.uber.org/yarpc/yarpcerrors"
)

// Interface is the server-side interface for the Store service.
type Interface interface {
	readonlystoreserver.Interface

	CompareAndSwap(
		ctx context.Context,
		Request *atomic.CompareAndSwap,
	) error

	Forget(
		ctx context.Context,
		Key *string,
	) error

	Increment(
		ctx context.Context,
		Key *string,
		Value *int64,
	) error
}

// New prepares an implementation of the Store service for
// registration.
//
// 	handler := StoreHandler{}
// 	dispatcher.Register(storeserver.New(handler))
func New(impl Interface, opts ...thrift.RegisterOption) []transport.Procedure {
	h := handler{impl}
	service := thrift.Service{
		Name: "Store",
		Methods: []thrift.Method{

			thrift.Method{
				Name: "compareAndSwap",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.CompareAndSwap),
				},
				Signature:    "CompareAndSwap(Request *atomic.CompareAndSwap)",
				ThriftModule: atomic.ThriftModule,
			},

			thrift.Method{
				Name: "forget",
				HandlerSpec: thrift.HandlerSpec{

					Type:   transport.Oneway,
					Oneway: thrift.OnewayHandler(h.Forget),
				},
				Signature:    "Forget(Key *string)",
				ThriftModule: atomic.ThriftModule,
			},

			thrift.Method{
				Name: "increment",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.Increment),
				},
				Signature:    "Increment(Key *string, Value *int64)",
				ThriftModule: atomic.ThriftModule,
			},
		},
	}

	procedures := make([]transport.Procedure, 0, 3)

	procedures = append(
		procedures,
		readonlystoreserver.New(
			impl,
			append(
				opts,
				thrift.Named("Store"),
			)...,
		)...,
	)
	procedures = append(procedures, thrift.BuildProcedures(service, opts...)...)
	return procedures
}

type handler struct{ impl Interface }

type errorNamer interface{ ErrorName() string }

type yarpcErrorCodeExtractor interface{ YARPCCode() *yarpcerrors.Code }

func (h handler) CompareAndSwap(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args atomic.Store_CompareAndSwap_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode Thrift request for service 'Store' procedure 'CompareAndSwap': %w", err)
	}

	appErr := h.impl.CompareAndSwap(ctx, args.Request)

	hadError := appErr != nil
	result, err := atomic.Store_CompareAndSwap_Helper.WrapResponse(appErr)

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

func (h handler) Forget(ctx context.Context, body wire.Value) error {
	var args atomic.Store_Forget_Args
	if err := args.FromWire(body); err != nil {
		return err
	}

	return h.impl.Forget(ctx, args.Key)
}

func (h handler) Increment(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args atomic.Store_Increment_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode Thrift request for service 'Store' procedure 'Increment': %w", err)
	}

	appErr := h.impl.Increment(ctx, args.Key, args.Value)

	hadError := appErr != nil
	result, err := atomic.Store_Increment_Helper.WrapResponse(appErr)

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
// representations of requests and responses for the Store service.
func JSONifier() thrift.JSONifier {
	return &jsonifier{}
}

// GetService gets the name of the service for which this JSONifier can produce
// JSON representations of requests and responses.
func (s *jsonifier) GetService() string {
	return "Store"
}

// RequestToJSON returns a json representation of the request.
func (s *jsonifier) RequestToJSON(procedure string, requestBody wire.Value) ([]byte, error) {
	switch procedure {

	case "CompareAndSwap":
		var args atomic.Store_CompareAndSwap_Args
		if err := args.FromWire(requestBody); err != nil {
			return nil, err
		}
		return json.Marshal(args)

	case "Forget":
		var args atomic.Store_Forget_Args
		if err := args.FromWire(requestBody); err != nil {
			return nil, err
		}
		return json.Marshal(args)

	case "Increment":
		var args atomic.Store_Increment_Args
		if err := args.FromWire(requestBody); err != nil {
			return nil, err
		}
		return json.Marshal(args)

	default:
		return nil, yarpcerrors.InvalidArgumentErrorf(
			"could not produce JSON representation of Thrift request for service 'Store' procedure '%s'", procedure)
	}
}

// ResponseToJSON returns a json representation of the response.
func (s *jsonifier) ResponseToJSON(procedure string, responseBody wire.Value) ([]byte, error) {
	switch procedure {

	case "CompareAndSwap":
		var result atomic.Store_CompareAndSwap_Result
		if err := result.FromWire(responseBody); err != nil {
			return nil, err
		}
		return json.Marshal(result)

	case "Increment":
		var result atomic.Store_Increment_Result
		if err := result.FromWire(responseBody); err != nil {
			return nil, err
		}
		return json.Marshal(result)

	default:
		return nil, yarpcerrors.InvalidArgumentErrorf(
			"could not produce JSON representation of Thrift response for service 'Store' procedure '%s'", procedure)
	}
}

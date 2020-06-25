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

package thrifttestserver

import (
	context "context"
	json "encoding/json"
	wire "go.uber.org/thriftrw/wire"
	transport "go.uber.org/yarpc/api/transport"
	thrift "go.uber.org/yarpc/encoding/thrift"
	gauntlet "go.uber.org/yarpc/internal/crossdock/thrift/gauntlet"
	yarpcerrors "go.uber.org/yarpc/yarpcerrors"
)

// Interface is the server-side interface for the ThriftTest service.
type Interface interface {
	TestBinary(
		ctx context.Context,
		Thing []byte,
	) ([]byte, error)

	TestByte(
		ctx context.Context,
		Thing *int8,
	) (int8, error)

	TestDouble(
		ctx context.Context,
		Thing *float64,
	) (float64, error)

	TestEnum(
		ctx context.Context,
		Thing *gauntlet.Numberz,
	) (gauntlet.Numberz, error)

	TestException(
		ctx context.Context,
		Arg *string,
	) error

	TestI32(
		ctx context.Context,
		Thing *int32,
	) (int32, error)

	TestI64(
		ctx context.Context,
		Thing *int64,
	) (int64, error)

	TestInsanity(
		ctx context.Context,
		Argument *gauntlet.Insanity,
	) (map[gauntlet.UserId]map[gauntlet.Numberz]*gauntlet.Insanity, error)

	TestList(
		ctx context.Context,
		Thing []int32,
	) ([]int32, error)

	TestMap(
		ctx context.Context,
		Thing map[int32]int32,
	) (map[int32]int32, error)

	TestMapMap(
		ctx context.Context,
		Hello *int32,
	) (map[int32]map[int32]int32, error)

	TestMulti(
		ctx context.Context,
		Arg0 *int8,
		Arg1 *int32,
		Arg2 *int64,
		Arg3 map[int16]string,
		Arg4 *gauntlet.Numberz,
		Arg5 *gauntlet.UserId,
	) (*gauntlet.Xtruct, error)

	TestMultiException(
		ctx context.Context,
		Arg0 *string,
		Arg1 *string,
	) (*gauntlet.Xtruct, error)

	TestNest(
		ctx context.Context,
		Thing *gauntlet.Xtruct2,
	) (*gauntlet.Xtruct2, error)

	TestOneway(
		ctx context.Context,
		SecondsToSleep *int32,
	) error

	TestSet(
		ctx context.Context,
		Thing map[int32]struct{},
	) (map[int32]struct{}, error)

	TestString(
		ctx context.Context,
		Thing *string,
	) (string, error)

	TestStringMap(
		ctx context.Context,
		Thing map[string]string,
	) (map[string]string, error)

	TestStruct(
		ctx context.Context,
		Thing *gauntlet.Xtruct,
	) (*gauntlet.Xtruct, error)

	TestTypedef(
		ctx context.Context,
		Thing *gauntlet.UserId,
	) (gauntlet.UserId, error)

	TestVoid(
		ctx context.Context,
	) error
}

// New prepares an implementation of the ThriftTest service for
// registration.
//
// 	handler := ThriftTestHandler{}
// 	dispatcher.Register(thrifttestserver.New(handler))
func New(impl Interface, opts ...thrift.RegisterOption) []transport.Procedure {
	h := handler{impl}
	service := thrift.Service{
		Name: "ThriftTest",
		Methods: []thrift.Method{

			thrift.Method{
				Name: "testBinary",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.TestBinary),
				},
				Signature:    "TestBinary(Thing []byte) ([]byte)",
				ThriftModule: gauntlet.ThriftModule,
			},

			thrift.Method{
				Name: "testByte",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.TestByte),
				},
				Signature:    "TestByte(Thing *int8) (int8)",
				ThriftModule: gauntlet.ThriftModule,
			},

			thrift.Method{
				Name: "testDouble",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.TestDouble),
				},
				Signature:    "TestDouble(Thing *float64) (float64)",
				ThriftModule: gauntlet.ThriftModule,
			},

			thrift.Method{
				Name: "testEnum",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.TestEnum),
				},
				Signature:    "TestEnum(Thing *gauntlet.Numberz) (gauntlet.Numberz)",
				ThriftModule: gauntlet.ThriftModule,
			},

			thrift.Method{
				Name: "testException",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.TestException),
				},
				Signature:    "TestException(Arg *string)",
				ThriftModule: gauntlet.ThriftModule,
			},

			thrift.Method{
				Name: "testI32",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.TestI32),
				},
				Signature:    "TestI32(Thing *int32) (int32)",
				ThriftModule: gauntlet.ThriftModule,
			},

			thrift.Method{
				Name: "testI64",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.TestI64),
				},
				Signature:    "TestI64(Thing *int64) (int64)",
				ThriftModule: gauntlet.ThriftModule,
			},

			thrift.Method{
				Name: "testInsanity",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.TestInsanity),
				},
				Signature:    "TestInsanity(Argument *gauntlet.Insanity) (map[gauntlet.UserId]map[gauntlet.Numberz]*gauntlet.Insanity)",
				ThriftModule: gauntlet.ThriftModule,
			},

			thrift.Method{
				Name: "testList",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.TestList),
				},
				Signature:    "TestList(Thing []int32) ([]int32)",
				ThriftModule: gauntlet.ThriftModule,
			},

			thrift.Method{
				Name: "testMap",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.TestMap),
				},
				Signature:    "TestMap(Thing map[int32]int32) (map[int32]int32)",
				ThriftModule: gauntlet.ThriftModule,
			},

			thrift.Method{
				Name: "testMapMap",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.TestMapMap),
				},
				Signature:    "TestMapMap(Hello *int32) (map[int32]map[int32]int32)",
				ThriftModule: gauntlet.ThriftModule,
			},

			thrift.Method{
				Name: "testMulti",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.TestMulti),
				},
				Signature:    "TestMulti(Arg0 *int8, Arg1 *int32, Arg2 *int64, Arg3 map[int16]string, Arg4 *gauntlet.Numberz, Arg5 *gauntlet.UserId) (*gauntlet.Xtruct)",
				ThriftModule: gauntlet.ThriftModule,
			},

			thrift.Method{
				Name: "testMultiException",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.TestMultiException),
				},
				Signature:    "TestMultiException(Arg0 *string, Arg1 *string) (*gauntlet.Xtruct)",
				ThriftModule: gauntlet.ThriftModule,
			},

			thrift.Method{
				Name: "testNest",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.TestNest),
				},
				Signature:    "TestNest(Thing *gauntlet.Xtruct2) (*gauntlet.Xtruct2)",
				ThriftModule: gauntlet.ThriftModule,
			},

			thrift.Method{
				Name: "testOneway",
				HandlerSpec: thrift.HandlerSpec{

					Type:   transport.Oneway,
					Oneway: thrift.OnewayHandler(h.TestOneway),
				},
				Signature:    "TestOneway(SecondsToSleep *int32)",
				ThriftModule: gauntlet.ThriftModule,
			},

			thrift.Method{
				Name: "testSet",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.TestSet),
				},
				Signature:    "TestSet(Thing map[int32]struct{}) (map[int32]struct{})",
				ThriftModule: gauntlet.ThriftModule,
			},

			thrift.Method{
				Name: "testString",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.TestString),
				},
				Signature:    "TestString(Thing *string) (string)",
				ThriftModule: gauntlet.ThriftModule,
			},

			thrift.Method{
				Name: "testStringMap",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.TestStringMap),
				},
				Signature:    "TestStringMap(Thing map[string]string) (map[string]string)",
				ThriftModule: gauntlet.ThriftModule,
			},

			thrift.Method{
				Name: "testStruct",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.TestStruct),
				},
				Signature:    "TestStruct(Thing *gauntlet.Xtruct) (*gauntlet.Xtruct)",
				ThriftModule: gauntlet.ThriftModule,
			},

			thrift.Method{
				Name: "testTypedef",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.TestTypedef),
				},
				Signature:    "TestTypedef(Thing *gauntlet.UserId) (gauntlet.UserId)",
				ThriftModule: gauntlet.ThriftModule,
			},

			thrift.Method{
				Name: "testVoid",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.TestVoid),
				},
				Signature:    "TestVoid()",
				ThriftModule: gauntlet.ThriftModule,
			},
		},
	}

	procedures := make([]transport.Procedure, 0, 21)
	procedures = append(procedures, thrift.BuildProcedures(service, opts...)...)
	return procedures
}

type handler struct{ impl Interface }

type errorNamer interface{ ErrorName() string }

type yarpcErrorCodeExtractor interface{ YARPCCode() *yarpcerrors.Code }

func (h handler) TestBinary(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args gauntlet.ThriftTest_TestBinary_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode Thrift request for service 'ThriftTest' procedure 'TestBinary': %w", err)
	}

	success, appErr := h.impl.TestBinary(ctx, args.Thing)

	hadError := appErr != nil
	result, err := gauntlet.ThriftTest_TestBinary_Helper.WrapResponse(success, appErr)

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

func (h handler) TestByte(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args gauntlet.ThriftTest_TestByte_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode Thrift request for service 'ThriftTest' procedure 'TestByte': %w", err)
	}

	success, appErr := h.impl.TestByte(ctx, args.Thing)

	hadError := appErr != nil
	result, err := gauntlet.ThriftTest_TestByte_Helper.WrapResponse(success, appErr)

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

func (h handler) TestDouble(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args gauntlet.ThriftTest_TestDouble_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode Thrift request for service 'ThriftTest' procedure 'TestDouble': %w", err)
	}

	success, appErr := h.impl.TestDouble(ctx, args.Thing)

	hadError := appErr != nil
	result, err := gauntlet.ThriftTest_TestDouble_Helper.WrapResponse(success, appErr)

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

func (h handler) TestEnum(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args gauntlet.ThriftTest_TestEnum_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode Thrift request for service 'ThriftTest' procedure 'TestEnum': %w", err)
	}

	success, appErr := h.impl.TestEnum(ctx, args.Thing)

	hadError := appErr != nil
	result, err := gauntlet.ThriftTest_TestEnum_Helper.WrapResponse(success, appErr)

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

func (h handler) TestException(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args gauntlet.ThriftTest_TestException_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode Thrift request for service 'ThriftTest' procedure 'TestException': %w", err)
	}

	appErr := h.impl.TestException(ctx, args.Arg)

	hadError := appErr != nil
	result, err := gauntlet.ThriftTest_TestException_Helper.WrapResponse(appErr)

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

func (h handler) TestI32(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args gauntlet.ThriftTest_TestI32_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode Thrift request for service 'ThriftTest' procedure 'TestI32': %w", err)
	}

	success, appErr := h.impl.TestI32(ctx, args.Thing)

	hadError := appErr != nil
	result, err := gauntlet.ThriftTest_TestI32_Helper.WrapResponse(success, appErr)

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

func (h handler) TestI64(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args gauntlet.ThriftTest_TestI64_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode Thrift request for service 'ThriftTest' procedure 'TestI64': %w", err)
	}

	success, appErr := h.impl.TestI64(ctx, args.Thing)

	hadError := appErr != nil
	result, err := gauntlet.ThriftTest_TestI64_Helper.WrapResponse(success, appErr)

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

func (h handler) TestInsanity(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args gauntlet.ThriftTest_TestInsanity_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode Thrift request for service 'ThriftTest' procedure 'TestInsanity': %w", err)
	}

	success, appErr := h.impl.TestInsanity(ctx, args.Argument)

	hadError := appErr != nil
	result, err := gauntlet.ThriftTest_TestInsanity_Helper.WrapResponse(success, appErr)

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

func (h handler) TestList(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args gauntlet.ThriftTest_TestList_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode Thrift request for service 'ThriftTest' procedure 'TestList': %w", err)
	}

	success, appErr := h.impl.TestList(ctx, args.Thing)

	hadError := appErr != nil
	result, err := gauntlet.ThriftTest_TestList_Helper.WrapResponse(success, appErr)

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

func (h handler) TestMap(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args gauntlet.ThriftTest_TestMap_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode Thrift request for service 'ThriftTest' procedure 'TestMap': %w", err)
	}

	success, appErr := h.impl.TestMap(ctx, args.Thing)

	hadError := appErr != nil
	result, err := gauntlet.ThriftTest_TestMap_Helper.WrapResponse(success, appErr)

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

func (h handler) TestMapMap(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args gauntlet.ThriftTest_TestMapMap_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode Thrift request for service 'ThriftTest' procedure 'TestMapMap': %w", err)
	}

	success, appErr := h.impl.TestMapMap(ctx, args.Hello)

	hadError := appErr != nil
	result, err := gauntlet.ThriftTest_TestMapMap_Helper.WrapResponse(success, appErr)

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

func (h handler) TestMulti(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args gauntlet.ThriftTest_TestMulti_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode Thrift request for service 'ThriftTest' procedure 'TestMulti': %w", err)
	}

	success, appErr := h.impl.TestMulti(ctx, args.Arg0, args.Arg1, args.Arg2, args.Arg3, args.Arg4, args.Arg5)

	hadError := appErr != nil
	result, err := gauntlet.ThriftTest_TestMulti_Helper.WrapResponse(success, appErr)

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

func (h handler) TestMultiException(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args gauntlet.ThriftTest_TestMultiException_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode Thrift request for service 'ThriftTest' procedure 'TestMultiException': %w", err)
	}

	success, appErr := h.impl.TestMultiException(ctx, args.Arg0, args.Arg1)

	hadError := appErr != nil
	result, err := gauntlet.ThriftTest_TestMultiException_Helper.WrapResponse(success, appErr)

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

func (h handler) TestNest(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args gauntlet.ThriftTest_TestNest_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode Thrift request for service 'ThriftTest' procedure 'TestNest': %w", err)
	}

	success, appErr := h.impl.TestNest(ctx, args.Thing)

	hadError := appErr != nil
	result, err := gauntlet.ThriftTest_TestNest_Helper.WrapResponse(success, appErr)

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

func (h handler) TestOneway(ctx context.Context, body wire.Value) error {
	var args gauntlet.ThriftTest_TestOneway_Args
	if err := args.FromWire(body); err != nil {
		return err
	}

	return h.impl.TestOneway(ctx, args.SecondsToSleep)
}

func (h handler) TestSet(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args gauntlet.ThriftTest_TestSet_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode Thrift request for service 'ThriftTest' procedure 'TestSet': %w", err)
	}

	success, appErr := h.impl.TestSet(ctx, args.Thing)

	hadError := appErr != nil
	result, err := gauntlet.ThriftTest_TestSet_Helper.WrapResponse(success, appErr)

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

func (h handler) TestString(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args gauntlet.ThriftTest_TestString_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode Thrift request for service 'ThriftTest' procedure 'TestString': %w", err)
	}

	success, appErr := h.impl.TestString(ctx, args.Thing)

	hadError := appErr != nil
	result, err := gauntlet.ThriftTest_TestString_Helper.WrapResponse(success, appErr)

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

func (h handler) TestStringMap(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args gauntlet.ThriftTest_TestStringMap_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode Thrift request for service 'ThriftTest' procedure 'TestStringMap': %w", err)
	}

	success, appErr := h.impl.TestStringMap(ctx, args.Thing)

	hadError := appErr != nil
	result, err := gauntlet.ThriftTest_TestStringMap_Helper.WrapResponse(success, appErr)

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

func (h handler) TestStruct(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args gauntlet.ThriftTest_TestStruct_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode Thrift request for service 'ThriftTest' procedure 'TestStruct': %w", err)
	}

	success, appErr := h.impl.TestStruct(ctx, args.Thing)

	hadError := appErr != nil
	result, err := gauntlet.ThriftTest_TestStruct_Helper.WrapResponse(success, appErr)

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

func (h handler) TestTypedef(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args gauntlet.ThriftTest_TestTypedef_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode Thrift request for service 'ThriftTest' procedure 'TestTypedef': %w", err)
	}

	success, appErr := h.impl.TestTypedef(ctx, args.Thing)

	hadError := appErr != nil
	result, err := gauntlet.ThriftTest_TestTypedef_Helper.WrapResponse(success, appErr)

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

func (h handler) TestVoid(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args gauntlet.ThriftTest_TestVoid_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode Thrift request for service 'ThriftTest' procedure 'TestVoid': %w", err)
	}

	appErr := h.impl.TestVoid(ctx)

	hadError := appErr != nil
	result, err := gauntlet.ThriftTest_TestVoid_Helper.WrapResponse(appErr)

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
// representations of requests and responses for the ThriftTest service.
func JSONifier() thrift.JSONifier {
	return &jsonifier{}
}

// GetService gets the name of the service for which this JSONifier can produce
// JSON representations of requests and responses.
func (s *jsonifier) GetService() string {
	return "ThriftTest"
}

// RequestToJSON returns a json representation of the request.
func (s *jsonifier) RequestToJSON(procedure string, requestBody wire.Value) ([]byte, error) {
	switch procedure {

	case "TestBinary":
		var args gauntlet.ThriftTest_TestBinary_Args
		if err := args.FromWire(requestBody); err != nil {
			return nil, err
		}
		return json.Marshal(args)

	case "TestByte":
		var args gauntlet.ThriftTest_TestByte_Args
		if err := args.FromWire(requestBody); err != nil {
			return nil, err
		}
		return json.Marshal(args)

	case "TestDouble":
		var args gauntlet.ThriftTest_TestDouble_Args
		if err := args.FromWire(requestBody); err != nil {
			return nil, err
		}
		return json.Marshal(args)

	case "TestEnum":
		var args gauntlet.ThriftTest_TestEnum_Args
		if err := args.FromWire(requestBody); err != nil {
			return nil, err
		}
		return json.Marshal(args)

	case "TestException":
		var args gauntlet.ThriftTest_TestException_Args
		if err := args.FromWire(requestBody); err != nil {
			return nil, err
		}
		return json.Marshal(args)

	case "TestI32":
		var args gauntlet.ThriftTest_TestI32_Args
		if err := args.FromWire(requestBody); err != nil {
			return nil, err
		}
		return json.Marshal(args)

	case "TestI64":
		var args gauntlet.ThriftTest_TestI64_Args
		if err := args.FromWire(requestBody); err != nil {
			return nil, err
		}
		return json.Marshal(args)

	case "TestInsanity":
		var args gauntlet.ThriftTest_TestInsanity_Args
		if err := args.FromWire(requestBody); err != nil {
			return nil, err
		}
		return json.Marshal(args)

	case "TestList":
		var args gauntlet.ThriftTest_TestList_Args
		if err := args.FromWire(requestBody); err != nil {
			return nil, err
		}
		return json.Marshal(args)

	case "TestMap":
		var args gauntlet.ThriftTest_TestMap_Args
		if err := args.FromWire(requestBody); err != nil {
			return nil, err
		}
		return json.Marshal(args)

	case "TestMapMap":
		var args gauntlet.ThriftTest_TestMapMap_Args
		if err := args.FromWire(requestBody); err != nil {
			return nil, err
		}
		return json.Marshal(args)

	case "TestMulti":
		var args gauntlet.ThriftTest_TestMulti_Args
		if err := args.FromWire(requestBody); err != nil {
			return nil, err
		}
		return json.Marshal(args)

	case "TestMultiException":
		var args gauntlet.ThriftTest_TestMultiException_Args
		if err := args.FromWire(requestBody); err != nil {
			return nil, err
		}
		return json.Marshal(args)

	case "TestNest":
		var args gauntlet.ThriftTest_TestNest_Args
		if err := args.FromWire(requestBody); err != nil {
			return nil, err
		}
		return json.Marshal(args)

	case "TestOneway":
		var args gauntlet.ThriftTest_TestOneway_Args
		if err := args.FromWire(requestBody); err != nil {
			return nil, err
		}
		return json.Marshal(args)

	case "TestSet":
		var args gauntlet.ThriftTest_TestSet_Args
		if err := args.FromWire(requestBody); err != nil {
			return nil, err
		}
		return json.Marshal(args)

	case "TestString":
		var args gauntlet.ThriftTest_TestString_Args
		if err := args.FromWire(requestBody); err != nil {
			return nil, err
		}
		return json.Marshal(args)

	case "TestStringMap":
		var args gauntlet.ThriftTest_TestStringMap_Args
		if err := args.FromWire(requestBody); err != nil {
			return nil, err
		}
		return json.Marshal(args)

	case "TestStruct":
		var args gauntlet.ThriftTest_TestStruct_Args
		if err := args.FromWire(requestBody); err != nil {
			return nil, err
		}
		return json.Marshal(args)

	case "TestTypedef":
		var args gauntlet.ThriftTest_TestTypedef_Args
		if err := args.FromWire(requestBody); err != nil {
			return nil, err
		}
		return json.Marshal(args)

	case "TestVoid":
		var args gauntlet.ThriftTest_TestVoid_Args
		if err := args.FromWire(requestBody); err != nil {
			return nil, err
		}
		return json.Marshal(args)

	default:
		return nil, yarpcerrors.InvalidArgumentErrorf(
			"could not produce JSON representation of Thrift request for service 'ThriftTest' procedure '%s'", procedure)
	}
}

// ResponseToJSON returns a json representation of the response.
func (s *jsonifier) ResponseToJSON(procedure string, responseBody wire.Value) ([]byte, error) {
	switch procedure {

	case "TestBinary":
		var result gauntlet.ThriftTest_TestBinary_Result
		if err := result.FromWire(responseBody); err != nil {
			return nil, err
		}
		return json.Marshal(result)

	case "TestByte":
		var result gauntlet.ThriftTest_TestByte_Result
		if err := result.FromWire(responseBody); err != nil {
			return nil, err
		}
		return json.Marshal(result)

	case "TestDouble":
		var result gauntlet.ThriftTest_TestDouble_Result
		if err := result.FromWire(responseBody); err != nil {
			return nil, err
		}
		return json.Marshal(result)

	case "TestEnum":
		var result gauntlet.ThriftTest_TestEnum_Result
		if err := result.FromWire(responseBody); err != nil {
			return nil, err
		}
		return json.Marshal(result)

	case "TestException":
		var result gauntlet.ThriftTest_TestException_Result
		if err := result.FromWire(responseBody); err != nil {
			return nil, err
		}
		return json.Marshal(result)

	case "TestI32":
		var result gauntlet.ThriftTest_TestI32_Result
		if err := result.FromWire(responseBody); err != nil {
			return nil, err
		}
		return json.Marshal(result)

	case "TestI64":
		var result gauntlet.ThriftTest_TestI64_Result
		if err := result.FromWire(responseBody); err != nil {
			return nil, err
		}
		return json.Marshal(result)

	case "TestInsanity":
		var result gauntlet.ThriftTest_TestInsanity_Result
		if err := result.FromWire(responseBody); err != nil {
			return nil, err
		}
		return json.Marshal(result)

	case "TestList":
		var result gauntlet.ThriftTest_TestList_Result
		if err := result.FromWire(responseBody); err != nil {
			return nil, err
		}
		return json.Marshal(result)

	case "TestMap":
		var result gauntlet.ThriftTest_TestMap_Result
		if err := result.FromWire(responseBody); err != nil {
			return nil, err
		}
		return json.Marshal(result)

	case "TestMapMap":
		var result gauntlet.ThriftTest_TestMapMap_Result
		if err := result.FromWire(responseBody); err != nil {
			return nil, err
		}
		return json.Marshal(result)

	case "TestMulti":
		var result gauntlet.ThriftTest_TestMulti_Result
		if err := result.FromWire(responseBody); err != nil {
			return nil, err
		}
		return json.Marshal(result)

	case "TestMultiException":
		var result gauntlet.ThriftTest_TestMultiException_Result
		if err := result.FromWire(responseBody); err != nil {
			return nil, err
		}
		return json.Marshal(result)

	case "TestNest":
		var result gauntlet.ThriftTest_TestNest_Result
		if err := result.FromWire(responseBody); err != nil {
			return nil, err
		}
		return json.Marshal(result)

	case "TestSet":
		var result gauntlet.ThriftTest_TestSet_Result
		if err := result.FromWire(responseBody); err != nil {
			return nil, err
		}
		return json.Marshal(result)

	case "TestString":
		var result gauntlet.ThriftTest_TestString_Result
		if err := result.FromWire(responseBody); err != nil {
			return nil, err
		}
		return json.Marshal(result)

	case "TestStringMap":
		var result gauntlet.ThriftTest_TestStringMap_Result
		if err := result.FromWire(responseBody); err != nil {
			return nil, err
		}
		return json.Marshal(result)

	case "TestStruct":
		var result gauntlet.ThriftTest_TestStruct_Result
		if err := result.FromWire(responseBody); err != nil {
			return nil, err
		}
		return json.Marshal(result)

	case "TestTypedef":
		var result gauntlet.ThriftTest_TestTypedef_Result
		if err := result.FromWire(responseBody); err != nil {
			return nil, err
		}
		return json.Marshal(result)

	case "TestVoid":
		var result gauntlet.ThriftTest_TestVoid_Result
		if err := result.FromWire(responseBody); err != nil {
			return nil, err
		}
		return json.Marshal(result)

	default:
		return nil, yarpcerrors.InvalidArgumentErrorf(
			"could not produce JSON representation of Thrift response for service 'ThriftTest' procedure '%s'", procedure)
	}
}

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

package secondservicefx

import (
	fx "go.uber.org/fx"
	transport "go.uber.org/yarpc/api/transport"
	thrift "go.uber.org/yarpc/encoding/thrift"
	secondserviceserver "go.uber.org/yarpc/internal/crossdock/thrift/gauntlet/secondserviceserver"
)

// ServerParams defines the dependencies for the SecondService server.
type ServerParams struct {
	fx.In

	Handler secondserviceserver.Interface
}

// ServerResult defines the output of SecondService server module. It provides the
// procedures of a SecondService handler to an Fx application.
//
// The procedures are provided to the "yarpcfx" value group. Dig 1.2 or newer
// must be used for this feature to work.
type ServerResult struct {
	fx.Out

	Procedures []transport.Procedure `group:"yarpcfx"`
}

// Server provides procedures for SecondService to an Fx application. It expects a
// secondservicefx.Interface to be present in the container.
//
// 	fx.Provide(
// 		func(h *MySecondServiceHandler) secondserviceserver.Interface {
// 			return h
// 		},
// 		secondservicefx.Server(),
// 	)
func Server(opts ...thrift.RegisterOption) interface{} {
	return func(p ServerParams) ServerResult {
		procedures := secondserviceserver.New(p.Handler, opts...)
		return ServerResult{Procedures: procedures}
	}
}

// Decoder provides a thrift.Decoder for SecondService to an Fx application. It is
// annotated with the value group "thriftdecoders".
//
// 	fx.Provide(
// 		func(h *MySecondServiceHandler) secondserviceserver.Interface {
// 			return h
// 		},
// 		secondservicefx.Server(),
// 		secondservicefx.Decoder(),
// 	)
func Decoder() interface{} {
	return fx.Annotated{
		Group:  "thriftdecoders",
		Target: secondserviceserver.Decoder,
	}
}

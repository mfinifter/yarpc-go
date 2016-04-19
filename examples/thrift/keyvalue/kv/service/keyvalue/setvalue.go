// Code generated by thriftrw

// Copyright (c) 2016 Uber Technologies, Inc.
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

package keyvalue

import (
	"fmt"
	"github.com/thriftrw/thriftrw-go/wire"
	"strings"
)

type SetValueArgs struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

func (v *SetValueArgs) ToWire() wire.Value {
	var fields [2]wire.Field
	i := 0
	if v.Key != nil {
		fields[i] = wire.Field{ID: 1, Value: wire.NewValueString(*(v.Key))}
		i++
	}
	if v.Value != nil {
		fields[i] = wire.Field{ID: 2, Value: wire.NewValueString(*(v.Value))}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]})
}

func (v *SetValueArgs) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Key = &x
				if err != nil {
					return err
				}
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Value = &x
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *SetValueArgs) String() string {
	var fields [2]string
	i := 0
	if v.Key != nil {
		fields[i] = fmt.Sprintf("Key: %v", *(v.Key))
		i++
	}
	if v.Value != nil {
		fields[i] = fmt.Sprintf("Value: %v", *(v.Value))
		i++
	}
	return fmt.Sprintf("SetValueArgs{%v}", strings.Join(fields[:i], ", "))
}

type SetValueResult struct{}

func (v *SetValueResult) ToWire() wire.Value {
	var fields [0]wire.Field
	i := 0
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]})
}

func (v *SetValueResult) FromWire(w wire.Value) error {
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}
	return nil
}

func (v *SetValueResult) String() string {
	var fields [0]string
	i := 0
	return fmt.Sprintf("SetValueResult{%v}", strings.Join(fields[:i], ", "))
}

var SetValueHelper = struct {
	IsException    func(error) bool
	Args           func(key *string, value *string) *SetValueArgs
	WrapResponse   func(error) (*SetValueResult, error)
	UnwrapResponse func(*SetValueResult) error
}{}

func init() {
	SetValueHelper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}
	SetValueHelper.Args = func(key *string, value *string) *SetValueArgs {
		return &SetValueArgs{Key: key, Value: value}
	}
	SetValueHelper.WrapResponse = func(err error) (*SetValueResult, error) {
		if err == nil {
			return &SetValueResult{}, nil
		}
		return nil, err
	}
	SetValueHelper.UnwrapResponse = func(result *SetValueResult) (err error) {
		return
	}
}
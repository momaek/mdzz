package mdzz

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
)

var (
	NotFound = errors.New("not found")
)

type Mux struct {
	m *Safetymap
}

type Method struct {
	method reflect.Method
	args   []reflect.Type
}

func NewMux() *Mux {
	return &Mux{
		m: NewSafetyMap(),
	}
}

func (mux *Mux) Register(rcvr interface{}) {
	typ := reflect.TypeOf(rcvr)
	// rcvr1 := reflect.ValueOf(rcvr)

	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		mt := method.Type

		nOut := mt.NumOut()
		if nOut < 1 {
			panic(fmt.Sprintf("%s return final output param must be error interface", method.Name))
		}
		_, ok := mt.Out(nOut - 1).MethodByName("Error") // 返回值的最后一个参数必须为 error
		if !ok {
			panic(fmt.Sprintf("%s return final output param must be error interface", method.Name))
		}

		m := Method{}
		m.method = method
		args := []reflect.Type{}
		for p := 1; p < mt.NumIn(); p++ {
			args = append(args, mt.In(p))
		}
		m.args = args

		mux.m.Set(method.Name, m)
	}
}

func (mux *Mux) Call(key string, rcvr interface{}, req *http.Request) (interface{}, error) {
	mi, ok := mux.m.Get(key)
	if !ok {
		return nil, NotFound
	}

	method, ok := mi.(Method)
	if !ok {
		return nil, NotFound
	}

	in := []reflect.Value{}
	in = append(in, reflect.ValueOf(rcvr))

	for _, v := range method.args {
		var (
			result interface{}
			val    reflect.Value
		)
		if v.Kind() == reflect.Ptr {
			result = reflect.New(v.Elem()).Elem().Interface()
			// TODO 这个地方没有把 BindValuesToStruct 放出来
			// 如果有需要的话给我 issue
			val = params.BindValuesToStruct(result, req, true).Addr()
		} else {
			result = reflect.New(v).Elem().Interface()
			val = params.BindValuesToStruct(result, req, true)
		}

		in = append(in, val)
	}

	ret := method.method.Func.Call(in)

	retLength := len(ret)
	var (
		err error
		res interface{}
	)
	err, _ = ret[retLength-1].Interface().(error)
	if retLength > 1 {
		res = ret[0].Interface()
	}

	return res, err
}

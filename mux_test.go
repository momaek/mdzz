package mdzz

import (
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewMux(t *testing.T) {
	am := NewMux()
	require.NotNil(t, am)
}

type TestSt struct {
	str string
	in  int
	flo float64
}

type Args struct {
	Helloworld string `json:"helloworld"`
	Nice       []int  `json:"nice"`
}

func (t *TestSt) ThisIsATest(args *Args) (err error) {
	return
}

func (t *TestSt) ThisIsATest2(args Args) (int, error) {
	return 2, errors.New("text")
}

func TestRegister(t *testing.T) {
	am := NewMux()
	require.NotNil(t, am)

	ts := new(TestSt)

	am.Register(ts)
}

func TestCall(t *testing.T) {
	am := NewMux()
	require.NotNil(t, am)
	ts := new(TestSt)
	am.Register(ts)

	param := url.Values{}
	param.Set("helloworld", "value")
	param.Set("nice.0", "1")
	param.Set("nice.1", "2")

	req, _ := http.NewRequest(http.MethodGet, "http://xxx.com?"+param.Encode(), nil)

	rcvr := TestSt{
		str: "test222222",
	}

	ret, err := am.Call("ThisIsATest", &rcvr, req)
	require.NoError(t, err)
	require.Nil(t, ret)

	ret2, err2 := am.Call("ThisIsATest2", &rcvr, req)
	require.Error(t, err2)
	d, ok := ret2.(int)
	require.Equal(t, true, ok)
	require.Equal(t, 2, d)
}

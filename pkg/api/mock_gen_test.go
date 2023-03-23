// Code generated by mocktail; DO NOT EDIT.

package api

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/traefik/hub-agent-kubernetes/pkg/edgeingress"
)

// platformClientMock mock of PlatformClient.
type platformClientMock struct{ mock.Mock }

// newPlatformClientMock creates a new platformClientMock.
func newPlatformClientMock(tb testing.TB) *platformClientMock {
	tb.Helper()

	m := &platformClientMock{}
	m.Mock.Test(tb)

	tb.Cleanup(func() { m.AssertExpectations(tb) })

	return m
}

func (_m *platformClientMock) GetAPIs(_ context.Context) ([]API, error) {
	_ret := _m.Called()

	if _rf, ok := _ret.Get(0).(func() ([]API, error)); ok {
		return _rf()
	}

	_ra0, _ := _ret.Get(0).([]API)
	_rb1 := _ret.Error(1)

	return _ra0, _rb1
}

func (_m *platformClientMock) OnGetAPIs() *platformClientGetAPIsCall {
	return &platformClientGetAPIsCall{Call: _m.Mock.On("GetAPIs"), Parent: _m}
}

func (_m *platformClientMock) OnGetAPIsRaw() *platformClientGetAPIsCall {
	return &platformClientGetAPIsCall{Call: _m.Mock.On("GetAPIs"), Parent: _m}
}

type platformClientGetAPIsCall struct {
	*mock.Call
	Parent *platformClientMock
}

func (_c *platformClientGetAPIsCall) Panic(msg string) *platformClientGetAPIsCall {
	_c.Call = _c.Call.Panic(msg)
	return _c
}

func (_c *platformClientGetAPIsCall) Once() *platformClientGetAPIsCall {
	_c.Call = _c.Call.Once()
	return _c
}

func (_c *platformClientGetAPIsCall) Twice() *platformClientGetAPIsCall {
	_c.Call = _c.Call.Twice()
	return _c
}

func (_c *platformClientGetAPIsCall) Times(i int) *platformClientGetAPIsCall {
	_c.Call = _c.Call.Times(i)
	return _c
}

func (_c *platformClientGetAPIsCall) WaitUntil(w <-chan time.Time) *platformClientGetAPIsCall {
	_c.Call = _c.Call.WaitUntil(w)
	return _c
}

func (_c *platformClientGetAPIsCall) After(d time.Duration) *platformClientGetAPIsCall {
	_c.Call = _c.Call.After(d)
	return _c
}

func (_c *platformClientGetAPIsCall) Run(fn func(args mock.Arguments)) *platformClientGetAPIsCall {
	_c.Call = _c.Call.Run(fn)
	return _c
}

func (_c *platformClientGetAPIsCall) Maybe() *platformClientGetAPIsCall {
	_c.Call = _c.Call.Maybe()
	return _c
}

func (_c *platformClientGetAPIsCall) TypedReturns(a []API, b error) *platformClientGetAPIsCall {
	_c.Call = _c.Return(a, b)
	return _c
}

func (_c *platformClientGetAPIsCall) ReturnsFn(fn func() ([]API, error)) *platformClientGetAPIsCall {
	_c.Call = _c.Return(fn)
	return _c
}

func (_c *platformClientGetAPIsCall) TypedRun(fn func()) *platformClientGetAPIsCall {
	_c.Call = _c.Call.Run(func(args mock.Arguments) {
		fn()
	})
	return _c
}

func (_c *platformClientGetAPIsCall) OnGetAPIs() *platformClientGetAPIsCall {
	return _c.Parent.OnGetAPIs()
}

func (_c *platformClientGetAPIsCall) OnGetAccesses() *platformClientGetAccessesCall {
	return _c.Parent.OnGetAccesses()
}

func (_c *platformClientGetAPIsCall) OnGetCertificateByDomains(domains []string) *platformClientGetCertificateByDomainsCall {
	return _c.Parent.OnGetCertificateByDomains(domains)
}

func (_c *platformClientGetAPIsCall) OnGetCollections() *platformClientGetCollectionsCall {
	return _c.Parent.OnGetCollections()
}

func (_c *platformClientGetAPIsCall) OnGetGateways() *platformClientGetGatewaysCall {
	return _c.Parent.OnGetGateways()
}

func (_c *platformClientGetAPIsCall) OnGetPortals() *platformClientGetPortalsCall {
	return _c.Parent.OnGetPortals()
}

func (_c *platformClientGetAPIsCall) OnGetWildcardCertificate() *platformClientGetWildcardCertificateCall {
	return _c.Parent.OnGetWildcardCertificate()
}

func (_c *platformClientGetAPIsCall) OnGetAPIsRaw() *platformClientGetAPIsCall {
	return _c.Parent.OnGetAPIsRaw()
}

func (_c *platformClientGetAPIsCall) OnGetAccessesRaw() *platformClientGetAccessesCall {
	return _c.Parent.OnGetAccessesRaw()
}

func (_c *platformClientGetAPIsCall) OnGetCertificateByDomainsRaw(domains interface{}) *platformClientGetCertificateByDomainsCall {
	return _c.Parent.OnGetCertificateByDomainsRaw(domains)
}

func (_c *platformClientGetAPIsCall) OnGetCollectionsRaw() *platformClientGetCollectionsCall {
	return _c.Parent.OnGetCollectionsRaw()
}

func (_c *platformClientGetAPIsCall) OnGetGatewaysRaw() *platformClientGetGatewaysCall {
	return _c.Parent.OnGetGatewaysRaw()
}

func (_c *platformClientGetAPIsCall) OnGetPortalsRaw() *platformClientGetPortalsCall {
	return _c.Parent.OnGetPortalsRaw()
}

func (_c *platformClientGetAPIsCall) OnGetWildcardCertificateRaw() *platformClientGetWildcardCertificateCall {
	return _c.Parent.OnGetWildcardCertificateRaw()
}

func (_m *platformClientMock) GetAccesses(_ context.Context) ([]Access, error) {
	_ret := _m.Called()

	if _rf, ok := _ret.Get(0).(func() ([]Access, error)); ok {
		return _rf()
	}

	_ra0, _ := _ret.Get(0).([]Access)
	_rb1 := _ret.Error(1)

	return _ra0, _rb1
}

func (_m *platformClientMock) OnGetAccesses() *platformClientGetAccessesCall {
	return &platformClientGetAccessesCall{Call: _m.Mock.On("GetAccesses"), Parent: _m}
}

func (_m *platformClientMock) OnGetAccessesRaw() *platformClientGetAccessesCall {
	return &platformClientGetAccessesCall{Call: _m.Mock.On("GetAccesses"), Parent: _m}
}

type platformClientGetAccessesCall struct {
	*mock.Call
	Parent *platformClientMock
}

func (_c *platformClientGetAccessesCall) Panic(msg string) *platformClientGetAccessesCall {
	_c.Call = _c.Call.Panic(msg)
	return _c
}

func (_c *platformClientGetAccessesCall) Once() *platformClientGetAccessesCall {
	_c.Call = _c.Call.Once()
	return _c
}

func (_c *platformClientGetAccessesCall) Twice() *platformClientGetAccessesCall {
	_c.Call = _c.Call.Twice()
	return _c
}

func (_c *platformClientGetAccessesCall) Times(i int) *platformClientGetAccessesCall {
	_c.Call = _c.Call.Times(i)
	return _c
}

func (_c *platformClientGetAccessesCall) WaitUntil(w <-chan time.Time) *platformClientGetAccessesCall {
	_c.Call = _c.Call.WaitUntil(w)
	return _c
}

func (_c *platformClientGetAccessesCall) After(d time.Duration) *platformClientGetAccessesCall {
	_c.Call = _c.Call.After(d)
	return _c
}

func (_c *platformClientGetAccessesCall) Run(fn func(args mock.Arguments)) *platformClientGetAccessesCall {
	_c.Call = _c.Call.Run(fn)
	return _c
}

func (_c *platformClientGetAccessesCall) Maybe() *platformClientGetAccessesCall {
	_c.Call = _c.Call.Maybe()
	return _c
}

func (_c *platformClientGetAccessesCall) TypedReturns(a []Access, b error) *platformClientGetAccessesCall {
	_c.Call = _c.Return(a, b)
	return _c
}

func (_c *platformClientGetAccessesCall) ReturnsFn(fn func() ([]Access, error)) *platformClientGetAccessesCall {
	_c.Call = _c.Return(fn)
	return _c
}

func (_c *platformClientGetAccessesCall) TypedRun(fn func()) *platformClientGetAccessesCall {
	_c.Call = _c.Call.Run(func(args mock.Arguments) {
		fn()
	})
	return _c
}

func (_c *platformClientGetAccessesCall) OnGetAPIs() *platformClientGetAPIsCall {
	return _c.Parent.OnGetAPIs()
}

func (_c *platformClientGetAccessesCall) OnGetAccesses() *platformClientGetAccessesCall {
	return _c.Parent.OnGetAccesses()
}

func (_c *platformClientGetAccessesCall) OnGetCertificateByDomains(domains []string) *platformClientGetCertificateByDomainsCall {
	return _c.Parent.OnGetCertificateByDomains(domains)
}

func (_c *platformClientGetAccessesCall) OnGetCollections() *platformClientGetCollectionsCall {
	return _c.Parent.OnGetCollections()
}

func (_c *platformClientGetAccessesCall) OnGetGateways() *platformClientGetGatewaysCall {
	return _c.Parent.OnGetGateways()
}

func (_c *platformClientGetAccessesCall) OnGetPortals() *platformClientGetPortalsCall {
	return _c.Parent.OnGetPortals()
}

func (_c *platformClientGetAccessesCall) OnGetWildcardCertificate() *platformClientGetWildcardCertificateCall {
	return _c.Parent.OnGetWildcardCertificate()
}

func (_c *platformClientGetAccessesCall) OnGetAPIsRaw() *platformClientGetAPIsCall {
	return _c.Parent.OnGetAPIsRaw()
}

func (_c *platformClientGetAccessesCall) OnGetAccessesRaw() *platformClientGetAccessesCall {
	return _c.Parent.OnGetAccessesRaw()
}

func (_c *platformClientGetAccessesCall) OnGetCertificateByDomainsRaw(domains interface{}) *platformClientGetCertificateByDomainsCall {
	return _c.Parent.OnGetCertificateByDomainsRaw(domains)
}

func (_c *platformClientGetAccessesCall) OnGetCollectionsRaw() *platformClientGetCollectionsCall {
	return _c.Parent.OnGetCollectionsRaw()
}

func (_c *platformClientGetAccessesCall) OnGetGatewaysRaw() *platformClientGetGatewaysCall {
	return _c.Parent.OnGetGatewaysRaw()
}

func (_c *platformClientGetAccessesCall) OnGetPortalsRaw() *platformClientGetPortalsCall {
	return _c.Parent.OnGetPortalsRaw()
}

func (_c *platformClientGetAccessesCall) OnGetWildcardCertificateRaw() *platformClientGetWildcardCertificateCall {
	return _c.Parent.OnGetWildcardCertificateRaw()
}

func (_m *platformClientMock) GetCertificateByDomains(_ context.Context, domains []string) (edgeingress.Certificate, error) {
	_ret := _m.Called(domains)

	if _rf, ok := _ret.Get(0).(func([]string) (edgeingress.Certificate, error)); ok {
		return _rf(domains)
	}

	_ra0, _ := _ret.Get(0).(edgeingress.Certificate)
	_rb1 := _ret.Error(1)

	return _ra0, _rb1
}

func (_m *platformClientMock) OnGetCertificateByDomains(domains []string) *platformClientGetCertificateByDomainsCall {
	return &platformClientGetCertificateByDomainsCall{Call: _m.Mock.On("GetCertificateByDomains", domains), Parent: _m}
}

func (_m *platformClientMock) OnGetCertificateByDomainsRaw(domains interface{}) *platformClientGetCertificateByDomainsCall {
	return &platformClientGetCertificateByDomainsCall{Call: _m.Mock.On("GetCertificateByDomains", domains), Parent: _m}
}

type platformClientGetCertificateByDomainsCall struct {
	*mock.Call
	Parent *platformClientMock
}

func (_c *platformClientGetCertificateByDomainsCall) Panic(msg string) *platformClientGetCertificateByDomainsCall {
	_c.Call = _c.Call.Panic(msg)
	return _c
}

func (_c *platformClientGetCertificateByDomainsCall) Once() *platformClientGetCertificateByDomainsCall {
	_c.Call = _c.Call.Once()
	return _c
}

func (_c *platformClientGetCertificateByDomainsCall) Twice() *platformClientGetCertificateByDomainsCall {
	_c.Call = _c.Call.Twice()
	return _c
}

func (_c *platformClientGetCertificateByDomainsCall) Times(i int) *platformClientGetCertificateByDomainsCall {
	_c.Call = _c.Call.Times(i)
	return _c
}

func (_c *platformClientGetCertificateByDomainsCall) WaitUntil(w <-chan time.Time) *platformClientGetCertificateByDomainsCall {
	_c.Call = _c.Call.WaitUntil(w)
	return _c
}

func (_c *platformClientGetCertificateByDomainsCall) After(d time.Duration) *platformClientGetCertificateByDomainsCall {
	_c.Call = _c.Call.After(d)
	return _c
}

func (_c *platformClientGetCertificateByDomainsCall) Run(fn func(args mock.Arguments)) *platformClientGetCertificateByDomainsCall {
	_c.Call = _c.Call.Run(fn)
	return _c
}

func (_c *platformClientGetCertificateByDomainsCall) Maybe() *platformClientGetCertificateByDomainsCall {
	_c.Call = _c.Call.Maybe()
	return _c
}

func (_c *platformClientGetCertificateByDomainsCall) TypedReturns(a edgeingress.Certificate, b error) *platformClientGetCertificateByDomainsCall {
	_c.Call = _c.Return(a, b)
	return _c
}

func (_c *platformClientGetCertificateByDomainsCall) ReturnsFn(fn func([]string) (edgeingress.Certificate, error)) *platformClientGetCertificateByDomainsCall {
	_c.Call = _c.Return(fn)
	return _c
}

func (_c *platformClientGetCertificateByDomainsCall) TypedRun(fn func([]string)) *platformClientGetCertificateByDomainsCall {
	_c.Call = _c.Call.Run(func(args mock.Arguments) {
		_domains, _ := args.Get(0).([]string)
		fn(_domains)
	})
	return _c
}

func (_c *platformClientGetCertificateByDomainsCall) OnGetAPIs() *platformClientGetAPIsCall {
	return _c.Parent.OnGetAPIs()
}

func (_c *platformClientGetCertificateByDomainsCall) OnGetAccesses() *platformClientGetAccessesCall {
	return _c.Parent.OnGetAccesses()
}

func (_c *platformClientGetCertificateByDomainsCall) OnGetCertificateByDomains(domains []string) *platformClientGetCertificateByDomainsCall {
	return _c.Parent.OnGetCertificateByDomains(domains)
}

func (_c *platformClientGetCertificateByDomainsCall) OnGetCollections() *platformClientGetCollectionsCall {
	return _c.Parent.OnGetCollections()
}

func (_c *platformClientGetCertificateByDomainsCall) OnGetGateways() *platformClientGetGatewaysCall {
	return _c.Parent.OnGetGateways()
}

func (_c *platformClientGetCertificateByDomainsCall) OnGetPortals() *platformClientGetPortalsCall {
	return _c.Parent.OnGetPortals()
}

func (_c *platformClientGetCertificateByDomainsCall) OnGetWildcardCertificate() *platformClientGetWildcardCertificateCall {
	return _c.Parent.OnGetWildcardCertificate()
}

func (_c *platformClientGetCertificateByDomainsCall) OnGetAPIsRaw() *platformClientGetAPIsCall {
	return _c.Parent.OnGetAPIsRaw()
}

func (_c *platformClientGetCertificateByDomainsCall) OnGetAccessesRaw() *platformClientGetAccessesCall {
	return _c.Parent.OnGetAccessesRaw()
}

func (_c *platformClientGetCertificateByDomainsCall) OnGetCertificateByDomainsRaw(domains interface{}) *platformClientGetCertificateByDomainsCall {
	return _c.Parent.OnGetCertificateByDomainsRaw(domains)
}

func (_c *platformClientGetCertificateByDomainsCall) OnGetCollectionsRaw() *platformClientGetCollectionsCall {
	return _c.Parent.OnGetCollectionsRaw()
}

func (_c *platformClientGetCertificateByDomainsCall) OnGetGatewaysRaw() *platformClientGetGatewaysCall {
	return _c.Parent.OnGetGatewaysRaw()
}

func (_c *platformClientGetCertificateByDomainsCall) OnGetPortalsRaw() *platformClientGetPortalsCall {
	return _c.Parent.OnGetPortalsRaw()
}

func (_c *platformClientGetCertificateByDomainsCall) OnGetWildcardCertificateRaw() *platformClientGetWildcardCertificateCall {
	return _c.Parent.OnGetWildcardCertificateRaw()
}

func (_m *platformClientMock) GetCollections(_ context.Context) ([]Collection, error) {
	_ret := _m.Called()

	if _rf, ok := _ret.Get(0).(func() ([]Collection, error)); ok {
		return _rf()
	}

	_ra0, _ := _ret.Get(0).([]Collection)
	_rb1 := _ret.Error(1)

	return _ra0, _rb1
}

func (_m *platformClientMock) OnGetCollections() *platformClientGetCollectionsCall {
	return &platformClientGetCollectionsCall{Call: _m.Mock.On("GetCollections"), Parent: _m}
}

func (_m *platformClientMock) OnGetCollectionsRaw() *platformClientGetCollectionsCall {
	return &platformClientGetCollectionsCall{Call: _m.Mock.On("GetCollections"), Parent: _m}
}

type platformClientGetCollectionsCall struct {
	*mock.Call
	Parent *platformClientMock
}

func (_c *platformClientGetCollectionsCall) Panic(msg string) *platformClientGetCollectionsCall {
	_c.Call = _c.Call.Panic(msg)
	return _c
}

func (_c *platformClientGetCollectionsCall) Once() *platformClientGetCollectionsCall {
	_c.Call = _c.Call.Once()
	return _c
}

func (_c *platformClientGetCollectionsCall) Twice() *platformClientGetCollectionsCall {
	_c.Call = _c.Call.Twice()
	return _c
}

func (_c *platformClientGetCollectionsCall) Times(i int) *platformClientGetCollectionsCall {
	_c.Call = _c.Call.Times(i)
	return _c
}

func (_c *platformClientGetCollectionsCall) WaitUntil(w <-chan time.Time) *platformClientGetCollectionsCall {
	_c.Call = _c.Call.WaitUntil(w)
	return _c
}

func (_c *platformClientGetCollectionsCall) After(d time.Duration) *platformClientGetCollectionsCall {
	_c.Call = _c.Call.After(d)
	return _c
}

func (_c *platformClientGetCollectionsCall) Run(fn func(args mock.Arguments)) *platformClientGetCollectionsCall {
	_c.Call = _c.Call.Run(fn)
	return _c
}

func (_c *platformClientGetCollectionsCall) Maybe() *platformClientGetCollectionsCall {
	_c.Call = _c.Call.Maybe()
	return _c
}

func (_c *platformClientGetCollectionsCall) TypedReturns(a []Collection, b error) *platformClientGetCollectionsCall {
	_c.Call = _c.Return(a, b)
	return _c
}

func (_c *platformClientGetCollectionsCall) ReturnsFn(fn func() ([]Collection, error)) *platformClientGetCollectionsCall {
	_c.Call = _c.Return(fn)
	return _c
}

func (_c *platformClientGetCollectionsCall) TypedRun(fn func()) *platformClientGetCollectionsCall {
	_c.Call = _c.Call.Run(func(args mock.Arguments) {
		fn()
	})
	return _c
}

func (_c *platformClientGetCollectionsCall) OnGetAPIs() *platformClientGetAPIsCall {
	return _c.Parent.OnGetAPIs()
}

func (_c *platformClientGetCollectionsCall) OnGetAccesses() *platformClientGetAccessesCall {
	return _c.Parent.OnGetAccesses()
}

func (_c *platformClientGetCollectionsCall) OnGetCertificateByDomains(domains []string) *platformClientGetCertificateByDomainsCall {
	return _c.Parent.OnGetCertificateByDomains(domains)
}

func (_c *platformClientGetCollectionsCall) OnGetCollections() *platformClientGetCollectionsCall {
	return _c.Parent.OnGetCollections()
}

func (_c *platformClientGetCollectionsCall) OnGetGateways() *platformClientGetGatewaysCall {
	return _c.Parent.OnGetGateways()
}

func (_c *platformClientGetCollectionsCall) OnGetPortals() *platformClientGetPortalsCall {
	return _c.Parent.OnGetPortals()
}

func (_c *platformClientGetCollectionsCall) OnGetWildcardCertificate() *platformClientGetWildcardCertificateCall {
	return _c.Parent.OnGetWildcardCertificate()
}

func (_c *platformClientGetCollectionsCall) OnGetAPIsRaw() *platformClientGetAPIsCall {
	return _c.Parent.OnGetAPIsRaw()
}

func (_c *platformClientGetCollectionsCall) OnGetAccessesRaw() *platformClientGetAccessesCall {
	return _c.Parent.OnGetAccessesRaw()
}

func (_c *platformClientGetCollectionsCall) OnGetCertificateByDomainsRaw(domains interface{}) *platformClientGetCertificateByDomainsCall {
	return _c.Parent.OnGetCertificateByDomainsRaw(domains)
}

func (_c *platformClientGetCollectionsCall) OnGetCollectionsRaw() *platformClientGetCollectionsCall {
	return _c.Parent.OnGetCollectionsRaw()
}

func (_c *platformClientGetCollectionsCall) OnGetGatewaysRaw() *platformClientGetGatewaysCall {
	return _c.Parent.OnGetGatewaysRaw()
}

func (_c *platformClientGetCollectionsCall) OnGetPortalsRaw() *platformClientGetPortalsCall {
	return _c.Parent.OnGetPortalsRaw()
}

func (_c *platformClientGetCollectionsCall) OnGetWildcardCertificateRaw() *platformClientGetWildcardCertificateCall {
	return _c.Parent.OnGetWildcardCertificateRaw()
}

func (_m *platformClientMock) GetGateways(_ context.Context) ([]Gateway, error) {
	_ret := _m.Called()

	if _rf, ok := _ret.Get(0).(func() ([]Gateway, error)); ok {
		return _rf()
	}

	_ra0, _ := _ret.Get(0).([]Gateway)
	_rb1 := _ret.Error(1)

	return _ra0, _rb1
}

func (_m *platformClientMock) OnGetGateways() *platformClientGetGatewaysCall {
	return &platformClientGetGatewaysCall{Call: _m.Mock.On("GetGateways"), Parent: _m}
}

func (_m *platformClientMock) OnGetGatewaysRaw() *platformClientGetGatewaysCall {
	return &platformClientGetGatewaysCall{Call: _m.Mock.On("GetGateways"), Parent: _m}
}

type platformClientGetGatewaysCall struct {
	*mock.Call
	Parent *platformClientMock
}

func (_c *platformClientGetGatewaysCall) Panic(msg string) *platformClientGetGatewaysCall {
	_c.Call = _c.Call.Panic(msg)
	return _c
}

func (_c *platformClientGetGatewaysCall) Once() *platformClientGetGatewaysCall {
	_c.Call = _c.Call.Once()
	return _c
}

func (_c *platformClientGetGatewaysCall) Twice() *platformClientGetGatewaysCall {
	_c.Call = _c.Call.Twice()
	return _c
}

func (_c *platformClientGetGatewaysCall) Times(i int) *platformClientGetGatewaysCall {
	_c.Call = _c.Call.Times(i)
	return _c
}

func (_c *platformClientGetGatewaysCall) WaitUntil(w <-chan time.Time) *platformClientGetGatewaysCall {
	_c.Call = _c.Call.WaitUntil(w)
	return _c
}

func (_c *platformClientGetGatewaysCall) After(d time.Duration) *platformClientGetGatewaysCall {
	_c.Call = _c.Call.After(d)
	return _c
}

func (_c *platformClientGetGatewaysCall) Run(fn func(args mock.Arguments)) *platformClientGetGatewaysCall {
	_c.Call = _c.Call.Run(fn)
	return _c
}

func (_c *platformClientGetGatewaysCall) Maybe() *platformClientGetGatewaysCall {
	_c.Call = _c.Call.Maybe()
	return _c
}

func (_c *platformClientGetGatewaysCall) TypedReturns(a []Gateway, b error) *platformClientGetGatewaysCall {
	_c.Call = _c.Return(a, b)
	return _c
}

func (_c *platformClientGetGatewaysCall) ReturnsFn(fn func() ([]Gateway, error)) *platformClientGetGatewaysCall {
	_c.Call = _c.Return(fn)
	return _c
}

func (_c *platformClientGetGatewaysCall) TypedRun(fn func()) *platformClientGetGatewaysCall {
	_c.Call = _c.Call.Run(func(args mock.Arguments) {
		fn()
	})
	return _c
}

func (_c *platformClientGetGatewaysCall) OnGetAPIs() *platformClientGetAPIsCall {
	return _c.Parent.OnGetAPIs()
}

func (_c *platformClientGetGatewaysCall) OnGetAccesses() *platformClientGetAccessesCall {
	return _c.Parent.OnGetAccesses()
}

func (_c *platformClientGetGatewaysCall) OnGetCertificateByDomains(domains []string) *platformClientGetCertificateByDomainsCall {
	return _c.Parent.OnGetCertificateByDomains(domains)
}

func (_c *platformClientGetGatewaysCall) OnGetCollections() *platformClientGetCollectionsCall {
	return _c.Parent.OnGetCollections()
}

func (_c *platformClientGetGatewaysCall) OnGetGateways() *platformClientGetGatewaysCall {
	return _c.Parent.OnGetGateways()
}

func (_c *platformClientGetGatewaysCall) OnGetPortals() *platformClientGetPortalsCall {
	return _c.Parent.OnGetPortals()
}

func (_c *platformClientGetGatewaysCall) OnGetWildcardCertificate() *platformClientGetWildcardCertificateCall {
	return _c.Parent.OnGetWildcardCertificate()
}

func (_c *platformClientGetGatewaysCall) OnGetAPIsRaw() *platformClientGetAPIsCall {
	return _c.Parent.OnGetAPIsRaw()
}

func (_c *platformClientGetGatewaysCall) OnGetAccessesRaw() *platformClientGetAccessesCall {
	return _c.Parent.OnGetAccessesRaw()
}

func (_c *platformClientGetGatewaysCall) OnGetCertificateByDomainsRaw(domains interface{}) *platformClientGetCertificateByDomainsCall {
	return _c.Parent.OnGetCertificateByDomainsRaw(domains)
}

func (_c *platformClientGetGatewaysCall) OnGetCollectionsRaw() *platformClientGetCollectionsCall {
	return _c.Parent.OnGetCollectionsRaw()
}

func (_c *platformClientGetGatewaysCall) OnGetGatewaysRaw() *platformClientGetGatewaysCall {
	return _c.Parent.OnGetGatewaysRaw()
}

func (_c *platformClientGetGatewaysCall) OnGetPortalsRaw() *platformClientGetPortalsCall {
	return _c.Parent.OnGetPortalsRaw()
}

func (_c *platformClientGetGatewaysCall) OnGetWildcardCertificateRaw() *platformClientGetWildcardCertificateCall {
	return _c.Parent.OnGetWildcardCertificateRaw()
}

func (_m *platformClientMock) GetPortals(_ context.Context) ([]Portal, error) {
	_ret := _m.Called()

	if _rf, ok := _ret.Get(0).(func() ([]Portal, error)); ok {
		return _rf()
	}

	_ra0, _ := _ret.Get(0).([]Portal)
	_rb1 := _ret.Error(1)

	return _ra0, _rb1
}

func (_m *platformClientMock) OnGetPortals() *platformClientGetPortalsCall {
	return &platformClientGetPortalsCall{Call: _m.Mock.On("GetPortals"), Parent: _m}
}

func (_m *platformClientMock) OnGetPortalsRaw() *platformClientGetPortalsCall {
	return &platformClientGetPortalsCall{Call: _m.Mock.On("GetPortals"), Parent: _m}
}

type platformClientGetPortalsCall struct {
	*mock.Call
	Parent *platformClientMock
}

func (_c *platformClientGetPortalsCall) Panic(msg string) *platformClientGetPortalsCall {
	_c.Call = _c.Call.Panic(msg)
	return _c
}

func (_c *platformClientGetPortalsCall) Once() *platformClientGetPortalsCall {
	_c.Call = _c.Call.Once()
	return _c
}

func (_c *platformClientGetPortalsCall) Twice() *platformClientGetPortalsCall {
	_c.Call = _c.Call.Twice()
	return _c
}

func (_c *platformClientGetPortalsCall) Times(i int) *platformClientGetPortalsCall {
	_c.Call = _c.Call.Times(i)
	return _c
}

func (_c *platformClientGetPortalsCall) WaitUntil(w <-chan time.Time) *platformClientGetPortalsCall {
	_c.Call = _c.Call.WaitUntil(w)
	return _c
}

func (_c *platformClientGetPortalsCall) After(d time.Duration) *platformClientGetPortalsCall {
	_c.Call = _c.Call.After(d)
	return _c
}

func (_c *platformClientGetPortalsCall) Run(fn func(args mock.Arguments)) *platformClientGetPortalsCall {
	_c.Call = _c.Call.Run(fn)
	return _c
}

func (_c *platformClientGetPortalsCall) Maybe() *platformClientGetPortalsCall {
	_c.Call = _c.Call.Maybe()
	return _c
}

func (_c *platformClientGetPortalsCall) TypedReturns(a []Portal, b error) *platformClientGetPortalsCall {
	_c.Call = _c.Return(a, b)
	return _c
}

func (_c *platformClientGetPortalsCall) ReturnsFn(fn func() ([]Portal, error)) *platformClientGetPortalsCall {
	_c.Call = _c.Return(fn)
	return _c
}

func (_c *platformClientGetPortalsCall) TypedRun(fn func()) *platformClientGetPortalsCall {
	_c.Call = _c.Call.Run(func(args mock.Arguments) {
		fn()
	})
	return _c
}

func (_c *platformClientGetPortalsCall) OnGetAPIs() *platformClientGetAPIsCall {
	return _c.Parent.OnGetAPIs()
}

func (_c *platformClientGetPortalsCall) OnGetAccesses() *platformClientGetAccessesCall {
	return _c.Parent.OnGetAccesses()
}

func (_c *platformClientGetPortalsCall) OnGetCertificateByDomains(domains []string) *platformClientGetCertificateByDomainsCall {
	return _c.Parent.OnGetCertificateByDomains(domains)
}

func (_c *platformClientGetPortalsCall) OnGetCollections() *platformClientGetCollectionsCall {
	return _c.Parent.OnGetCollections()
}

func (_c *platformClientGetPortalsCall) OnGetGateways() *platformClientGetGatewaysCall {
	return _c.Parent.OnGetGateways()
}

func (_c *platformClientGetPortalsCall) OnGetPortals() *platformClientGetPortalsCall {
	return _c.Parent.OnGetPortals()
}

func (_c *platformClientGetPortalsCall) OnGetWildcardCertificate() *platformClientGetWildcardCertificateCall {
	return _c.Parent.OnGetWildcardCertificate()
}

func (_c *platformClientGetPortalsCall) OnGetAPIsRaw() *platformClientGetAPIsCall {
	return _c.Parent.OnGetAPIsRaw()
}

func (_c *platformClientGetPortalsCall) OnGetAccessesRaw() *platformClientGetAccessesCall {
	return _c.Parent.OnGetAccessesRaw()
}

func (_c *platformClientGetPortalsCall) OnGetCertificateByDomainsRaw(domains interface{}) *platformClientGetCertificateByDomainsCall {
	return _c.Parent.OnGetCertificateByDomainsRaw(domains)
}

func (_c *platformClientGetPortalsCall) OnGetCollectionsRaw() *platformClientGetCollectionsCall {
	return _c.Parent.OnGetCollectionsRaw()
}

func (_c *platformClientGetPortalsCall) OnGetGatewaysRaw() *platformClientGetGatewaysCall {
	return _c.Parent.OnGetGatewaysRaw()
}

func (_c *platformClientGetPortalsCall) OnGetPortalsRaw() *platformClientGetPortalsCall {
	return _c.Parent.OnGetPortalsRaw()
}

func (_c *platformClientGetPortalsCall) OnGetWildcardCertificateRaw() *platformClientGetWildcardCertificateCall {
	return _c.Parent.OnGetWildcardCertificateRaw()
}

func (_m *platformClientMock) GetWildcardCertificate(_ context.Context) (edgeingress.Certificate, error) {
	_ret := _m.Called()

	if _rf, ok := _ret.Get(0).(func() (edgeingress.Certificate, error)); ok {
		return _rf()
	}

	_ra0, _ := _ret.Get(0).(edgeingress.Certificate)
	_rb1 := _ret.Error(1)

	return _ra0, _rb1
}

func (_m *platformClientMock) OnGetWildcardCertificate() *platformClientGetWildcardCertificateCall {
	return &platformClientGetWildcardCertificateCall{Call: _m.Mock.On("GetWildcardCertificate"), Parent: _m}
}

func (_m *platformClientMock) OnGetWildcardCertificateRaw() *platformClientGetWildcardCertificateCall {
	return &platformClientGetWildcardCertificateCall{Call: _m.Mock.On("GetWildcardCertificate"), Parent: _m}
}

type platformClientGetWildcardCertificateCall struct {
	*mock.Call
	Parent *platformClientMock
}

func (_c *platformClientGetWildcardCertificateCall) Panic(msg string) *platformClientGetWildcardCertificateCall {
	_c.Call = _c.Call.Panic(msg)
	return _c
}

func (_c *platformClientGetWildcardCertificateCall) Once() *platformClientGetWildcardCertificateCall {
	_c.Call = _c.Call.Once()
	return _c
}

func (_c *platformClientGetWildcardCertificateCall) Twice() *platformClientGetWildcardCertificateCall {
	_c.Call = _c.Call.Twice()
	return _c
}

func (_c *platformClientGetWildcardCertificateCall) Times(i int) *platformClientGetWildcardCertificateCall {
	_c.Call = _c.Call.Times(i)
	return _c
}

func (_c *platformClientGetWildcardCertificateCall) WaitUntil(w <-chan time.Time) *platformClientGetWildcardCertificateCall {
	_c.Call = _c.Call.WaitUntil(w)
	return _c
}

func (_c *platformClientGetWildcardCertificateCall) After(d time.Duration) *platformClientGetWildcardCertificateCall {
	_c.Call = _c.Call.After(d)
	return _c
}

func (_c *platformClientGetWildcardCertificateCall) Run(fn func(args mock.Arguments)) *platformClientGetWildcardCertificateCall {
	_c.Call = _c.Call.Run(fn)
	return _c
}

func (_c *platformClientGetWildcardCertificateCall) Maybe() *platformClientGetWildcardCertificateCall {
	_c.Call = _c.Call.Maybe()
	return _c
}

func (_c *platformClientGetWildcardCertificateCall) TypedReturns(a edgeingress.Certificate, b error) *platformClientGetWildcardCertificateCall {
	_c.Call = _c.Return(a, b)
	return _c
}

func (_c *platformClientGetWildcardCertificateCall) ReturnsFn(fn func() (edgeingress.Certificate, error)) *platformClientGetWildcardCertificateCall {
	_c.Call = _c.Return(fn)
	return _c
}

func (_c *platformClientGetWildcardCertificateCall) TypedRun(fn func()) *platformClientGetWildcardCertificateCall {
	_c.Call = _c.Call.Run(func(args mock.Arguments) {
		fn()
	})
	return _c
}

func (_c *platformClientGetWildcardCertificateCall) OnGetAPIs() *platformClientGetAPIsCall {
	return _c.Parent.OnGetAPIs()
}

func (_c *platformClientGetWildcardCertificateCall) OnGetAccesses() *platformClientGetAccessesCall {
	return _c.Parent.OnGetAccesses()
}

func (_c *platformClientGetWildcardCertificateCall) OnGetCertificateByDomains(domains []string) *platformClientGetCertificateByDomainsCall {
	return _c.Parent.OnGetCertificateByDomains(domains)
}

func (_c *platformClientGetWildcardCertificateCall) OnGetCollections() *platformClientGetCollectionsCall {
	return _c.Parent.OnGetCollections()
}

func (_c *platformClientGetWildcardCertificateCall) OnGetGateways() *platformClientGetGatewaysCall {
	return _c.Parent.OnGetGateways()
}

func (_c *platformClientGetWildcardCertificateCall) OnGetPortals() *platformClientGetPortalsCall {
	return _c.Parent.OnGetPortals()
}

func (_c *platformClientGetWildcardCertificateCall) OnGetWildcardCertificate() *platformClientGetWildcardCertificateCall {
	return _c.Parent.OnGetWildcardCertificate()
}

func (_c *platformClientGetWildcardCertificateCall) OnGetAPIsRaw() *platformClientGetAPIsCall {
	return _c.Parent.OnGetAPIsRaw()
}

func (_c *platformClientGetWildcardCertificateCall) OnGetAccessesRaw() *platformClientGetAccessesCall {
	return _c.Parent.OnGetAccessesRaw()
}

func (_c *platformClientGetWildcardCertificateCall) OnGetCertificateByDomainsRaw(domains interface{}) *platformClientGetCertificateByDomainsCall {
	return _c.Parent.OnGetCertificateByDomainsRaw(domains)
}

func (_c *platformClientGetWildcardCertificateCall) OnGetCollectionsRaw() *platformClientGetCollectionsCall {
	return _c.Parent.OnGetCollectionsRaw()
}

func (_c *platformClientGetWildcardCertificateCall) OnGetGatewaysRaw() *platformClientGetGatewaysCall {
	return _c.Parent.OnGetGatewaysRaw()
}

func (_c *platformClientGetWildcardCertificateCall) OnGetPortalsRaw() *platformClientGetPortalsCall {
	return _c.Parent.OnGetPortalsRaw()
}

func (_c *platformClientGetWildcardCertificateCall) OnGetWildcardCertificateRaw() *platformClientGetWildcardCertificateCall {
	return _c.Parent.OnGetWildcardCertificateRaw()
}
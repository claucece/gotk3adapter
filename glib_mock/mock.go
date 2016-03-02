package glib_mock

import "github.com/twstrike/gotk3adapter/glibi"

type Mock struct{}

func (*Mock) IdleAdd(f interface{}, args ...interface{}) (glibi.SourceHandle, error) {
	return glibi.SourceHandle(0), nil
}

func (*Mock) InitI18n(domain string, dir string) {
}

func (*Mock) Local(vx string) string {
	return vx
}

func (*Mock) SignalNew(s string) (glibi.Signal, error) {
	return &MockSignal{}, nil
}

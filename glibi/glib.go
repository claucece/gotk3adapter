package glibi

type Glib interface {
	IdleAdd(interface{}, ...interface{}) (SourceHandle, error)
	InitI18n(string, string)
	SignalNew(string) (Signal, error)
} // end of Glib

func AssertGlib(_ Glib) {}
package gtki

import "github.com/twstrike/gotk3adapter/gdki"

type Window interface {
	Bin

	AddAccelGroup(AccelGroup)
	GetTitle() string
	IsActive() bool
	Present()
	Resize(int, int)
	SetApplication(Application)
	SetIcon(gdki.Pixbuf)
	SetTitle(string)
	SetTitlebar(Widget) // Since 3.10
	SetTransientFor(Window)
}

func AssertWindow(_ Window) {}

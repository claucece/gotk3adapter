package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type textIter struct {
	*gtk.TextIter
}

func wrapTextIter(v *gtk.TextIter, e error) (*textIter, error) {
	if v == nil {
		return nil, e
	}
	return &textIter{v}, e
}

func unwrapTextIter(v gtki.TextIter) *gtk.TextIter {
	if v == nil {
		return nil
	}
	return v.(*textIter).TextIter
}

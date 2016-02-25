package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type treePath struct {
	*gtk.TreePath
}

func wrapTreePath(v *gtk.TreePath, e error) (*treePath, error) {
	if v == nil {
		return nil, e
	}
	return &treePath{v}, e
}

func unwrapTreePath(v gtki.TreePath) *gtk.TreePath {
	if v == nil {
		return nil
	}
	return v.(*treePath).TreePath
}

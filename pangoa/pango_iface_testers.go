package pangoa

import "github.com/twstrike/gotk3adapter/pangoi"

func init() {
	pangoi.AssertPango(&RealPango{})
	pangoi.AssertFontDescription(&fontDescription{})
}

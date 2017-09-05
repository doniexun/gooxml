// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import sml "baliance.com/gooxml/schema/schemas.openxmlformats.org/spreadsheetml"

// NumberFormat is a number formatting string that can be applied to a cell
// style.
type NumberFormat struct {
	wb *Workbook
	x  *sml.CT_NumFmt
}

// X returns the inner wrapped XML type.
func (n NumberFormat) X() *sml.CT_NumFmt {
	return n.x
}

// SetCode sets the number format code.
func (n NumberFormat) SetCode(f string) {
	n.x.FormatCodeAttr = f
}

// ID returns the number format ID.  This is not an index as there are some
// predefined number formats which can be used in cell styles and don't need a
// corresponding NumberFormat.
func (n NumberFormat) ID() uint32 {
	return n.x.NumFmtIdAttr
}

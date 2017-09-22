// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml_test

import (
	"encoding/xml"
	"testing"

	"baliance.com/gooxml/schema/soo/wordprocessingml"
)

func TestCT_ColorSchemeMappingConstructor(t *testing.T) {
	v := wordprocessingml.NewCT_ColorSchemeMapping()
	if v == nil {
		t.Errorf("wordprocessingml.NewCT_ColorSchemeMapping must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wordprocessingml.CT_ColorSchemeMapping should validate: %s", err)
	}
}

func TestCT_ColorSchemeMappingMarshalUnmarshal(t *testing.T) {
	v := wordprocessingml.NewCT_ColorSchemeMapping()
	buf, _ := xml.Marshal(v)
	v2 := wordprocessingml.NewCT_ColorSchemeMapping()
	xml.Unmarshal(buf, v2)
}
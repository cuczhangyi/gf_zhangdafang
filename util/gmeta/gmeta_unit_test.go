// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gmeta_test

import (
	"github.com/gogf/gf/internal/json"
	"github.com/gogf/gf/util/gmeta"
	"testing"

	"github.com/gogf/gf/test/gtest"
)

func TestMeta_Basic(t *testing.T) {
	type A struct {
		gmeta.Meta `tag:"123" orm:"456"`
		Id         int
		Name       string
	}

	gtest.C(t, func(t *gtest.T) {
		a := &A{
			Id:   100,
			Name: "john",
		}
		t.Assert(len(gmeta.Data(a)), 2)
		t.AssertEQ(gmeta.Get(a, "tag").String(), "123")
		t.AssertEQ(gmeta.Get(a, "orm").String(), "456")
		t.AssertEQ(gmeta.Get(a, "none"), nil)

		b, err := json.Marshal(a)
		t.AssertNil(err)
		t.Assert(b, `{"Id":100,"Name":"john"}`)
	})
}
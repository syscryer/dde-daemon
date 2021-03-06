/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package accounts

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetLocaleFromFile(t *testing.T) {
	Convey("getLocaleFromFile", t, func() {
		So(getLocaleFromFile("testdata/locale"), ShouldEqual, "zh_CN.UTF-8")
	})
}

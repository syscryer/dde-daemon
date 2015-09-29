package systeminfo

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestCPUInfo(t *testing.T) {
	Convey("Test cpu info", t, func() {
		cpu, err := getCPUInfo("testdata/cpuinfo")
		So(cpu, ShouldEqual,
			"Intel(R) Core(TM) i3 CPU M 330 @ 2.13GHz x 4")
		So(err, ShouldBeNil)

		cpu, err = getCPUInfo("testdata/sw-cpuinfo")
		So(cpu, ShouldEqual, "sw 1.4GHz x 4")
		So(err, ShouldBeNil)
	})
}

func TestMemInfo(t *testing.T) {
	Convey("Test memory info", t, func() {
		mem, err := getMemoryFromFile("testdata/meminfo")
		So(mem, ShouldEqual, 4005441536)
		So(err, ShouldBeNil)
	})
}

func TestVersion(t *testing.T) {
	Convey("Test os version", t, func() {
		lang := os.Getenv("LANGUAGE")
		os.Setenv("LANGUAGE", "en_US")
		defer os.Setenv("LANGUAGE", lang)

		deepin, err := getVersionFromDeepin("testdata/deepin-version")
		So(deepin, ShouldEqual, "2015 Desktop Alpha1")
		So(err, ShouldBeNil)

		lsb, err := getVersionFromLSB("testdata/lsb-release")
		So(lsb, ShouldEqual, "2014.3")
		So(err, ShouldBeNil)
	})
}
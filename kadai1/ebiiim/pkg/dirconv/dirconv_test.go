package dirconv_test

import (
	"reflect"
	"testing"

	"github.com/gopherdojo/dojo5/kadai1/ebiiim/pkg/dirconv"
)

func check(t *testing.T, got interface{}, want interface{}) {
	if got != want {
		t.Errorf("got: %v, want %v", got, want)
	}
}

func TestNewCli1(t *testing.T) {
	// normal: no options
	args := []string{"imgconv", "../../testdata"}
	dc := dirconv.NewDirConv(args)
	check(t, *dc, dirconv.DirConv{Dir: "../../testdata", SrcExt: ".jpg", TgtExt: ".png"})
}

func TestNewCli2(t *testing.T) {
	// normal: with options
	args := []string{"imgconv", "-source_ext=PNG", "-target_ext=.tiff", "../../testdata"}
	dc := dirconv.NewDirConv(args)
	check(t, *dc, dirconv.DirConv{Dir: "../../testdata", SrcExt: ".png", TgtExt: ".tiff"})
}

func TestCli_DirConv(t *testing.T) {
	args := []string{"imgconv", "../testdata"}
	got := dirconv.NewDirConv(args).Convert()
	want := []dirconv.Result{
		{Index: 0, RelPath: "dummy.jpg", IsOk: false},
		{Index: 1, RelPath: "gopherA.jpg", IsOk: true},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want %v", got, want)
	}
}
package md5

import "testing"

func TestMake(t *testing.T) {
	s := Make("test make")
	if s != "a746b9b89249687b95c0796392d258b0" {
		t.Error("md5 not match")
	}
}

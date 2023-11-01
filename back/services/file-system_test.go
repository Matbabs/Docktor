package services

import (
	"reflect"
	"testing"
)

type testCheckLocalPath struct {
	name string
	res  []string
}

func TestCheckLocalPath(t *testing.T) {

	executor = execCommandMock{}

	tests := []testCheckLocalPath{
		{name: "/", res: []string{".", "..", "bin", "boot", ".cache", "dev", "etc", "home", "lib", "lib32", "lib64", "libx32", "lost+found", "media", "mnt", "opt", "proc", "root", "run", "sbin", "snap", "srv", "sys", "tmp", "usr", "var", ""}},
		{name: "/home/user", res: []string{".", "..", "Documents", "Downloads", ""}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if res := CheckLocalPath(test.name); !reflect.DeepEqual(res, test.res) {
				t.Errorf("Expected %v, received %v", test.res, res)
			}
		})
	}

}

package systemdconf

import (
	"io/ioutil"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMarshalFile(t *testing.T) {
	unmarshal, err := ioutil.ReadFile("testdata/unmarshal.conf")
	if err != nil {
		panic("Error reading testdata/unmarshal.conf: " + err.Error())
	}
	marshal, err := ioutil.ReadFile("testdata/marshal.conf")
	if err != nil {
		panic("Error reading testdata/marshal.conf: " + err.Error())
	}
	var s fileTest
	err = Unmarshal(unmarshal, &s)
	if err != nil {
		t.Fatalf("Unmarshal() = %v; want nil", err)
	}
	b, err := Marshal(s)
	if err != nil {
		t.Fatalf("Marshal() = %v; want nil", err)
	}
	if diff := cmp.Diff(marshal, b, nil); diff != "" {
		t.Errorf("Marshal() mismatch (-want +got):\n%s", diff)
	}
}

func TestMarshalShouldFail(t *testing.T) {
	tests := []struct {
		Name string
		V    interface{}
	}{
		{
			Name: "nil",
			V:    nil,
		},
		{
			Name: "non-struct file",
			V:    true,
		},
		{
			Name: "non-struct section",
			V: struct {
				Section1 struct {
					Entry *string
				}
				Section2 bool
			}{},
		},
		{
			Name: "duplicated entry names",
			V: struct {
				Section1 struct {
					Entry       string
					Duplicated1 string `systemd:"Entry"`
					Duplicated2 string `systemd:"Entry"`
				}
			}{
				Section1: struct {
					Entry       string
					Duplicated1 string `systemd:"Entry"`
					Duplicated2 string `systemd:"Entry"`
				}{},
			},
		},
		{
			Name: "unsupported entry type",
			V: struct {
				Section1 struct {
					Entry string
				}
				Section2 struct {
					Entry int
				}
			}{},
		},
		{
			Name: "unsupported entry slice type",
			V: struct {
				Section1 struct {
					Entry string
				}
				Section2 struct {
					Entry []int
				}
			}{
				Section2: struct {
					Entry []int
				}{
					Entry: []int{0},
				},
			},
		},
	}
	for _, td := range tests {
		t.Run("Name="+td.Name, func(t *testing.T) {
			_, err := Marshal(td.V)
			if err == nil {
				t.Error("Marshal() = _, nil; want non-nil")
			}
		})
	}
}

func TestMarshalWithMarshaler(t *testing.T) {
	expected := []byte(`[Section]
Key=success: ok
`)
	s := marshalTest{
		Section: struct {
			Key marshalType
		}{
			Key: "ok",
		},
	}
	b, err := Marshal(s)
	if err != nil {
		t.Fatalf("Marshal() = %v; want nil", err)
	}
	if diff := cmp.Diff(expected, b, nil); diff != "" {
		t.Errorf("Marshal() mismatch (-want +got):\n%s", diff)
	}
}

func TestMarshalShouldFailOnMarshalerError(t *testing.T) {
	s := marshalTest{
		Section: struct {
			Key marshalType
		}{
			Key: "fail",
		},
	}
	_, err := Marshal(s)
	if err == nil {
		t.Errorf("Marshal() = nil; want non-nil")
	}
}

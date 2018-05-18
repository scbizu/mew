package drawer

import "testing"

func TestDraw(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"draw"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := DrawWithSlice("test", []string{"a", "b"})
			if err != nil {
				t.Error(err.Error())
			} else {
				t.Logf("output:%v", res)
			}
		})
	}
}

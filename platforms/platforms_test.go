package platforms

import (
	"reflect"
	"testing"
)

func TestConfigParse(t *testing.T) {
	tests := []struct {
		s string
		c Config
	}{
		{`x=1`, Config{"x": "1"}},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			c, err := parse(tt.s)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(c, tt.c) {
				t.Fatalf("Config error. Got %#v want %#v", c, tt.c)
			}
		})
	}
}

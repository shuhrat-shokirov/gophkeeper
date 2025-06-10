package memorycache

import (
	"reflect"
	"testing"

	"github.com/patrickmn/go-cache"
)

func TestNew(t *testing.T) {
	type args struct {
		p Params
	}
	tests := []struct {
		name string
		args args
		want Cache
	}{
		{
			name: "default cache",
			args: args{
				p: Params{},
			},
			want: &c{
				cache: cache.New(cache.NoExpiration, cache.NoExpiration),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_c_Get(t *testing.T) {
	t.Run("get existing key", func(t *testing.T) {
		c := New(Params{})
		c.Set("testKey", "testValue", cache.NoExpiration)

		value, found := c.Get("testKey")
		if !found || value != "testValue" {
			t.Errorf("Expected to find 'testValue' for 'testKey', got %v", value)
		}
	})
	t.Run("get non-existing key", func(t *testing.T) {
		c := New(Params{})
		value, found := c.Get("nonExistingKey")
		if found || value != nil {
			t.Errorf("Expected not to find 'nonExistingKey', got %v", value)
		}
	})
}

func Test_c_Set(t *testing.T) {
	t.Run("set existing key", func(t *testing.T) {
		c := New(Params{})
		c.Set("testKey", "testValue", cache.NoExpiration)

		value, found := c.Get("testKey")
		if !found || value != "testValue" {
			t.Errorf("Expected to find 'testValue' for 'testKey', got %v", value)
		}
	})

	t.Run("set non-existing key", func(t *testing.T) {
		c := New(Params{})
		c.Set("newKey", "newValue", cache.NoExpiration)

		value, found := c.Get("newKey")
		if !found || value != "newValue" {
			t.Errorf("Expected to find 'newValue' for 'newKey', got %v", value)
		}
	})
}

func Test_c_Delete(t *testing.T) {
	t.Run("delete existing key", func(t *testing.T) {
		c := New(Params{})
		c.Set("testKey", "testValue", cache.NoExpiration)
		c.Delete("testKey")

		value, found := c.Get("testKey")
		if found || value != nil {
			t.Errorf("Expected not to find 'testKey' after deletion, got %v", value)
		}
	})

	t.Run("delete non-existing key", func(t *testing.T) {
		c := New(Params{})
		c.Delete("nonExistingKey") // Should not panic or error
	})
}

package utils

import (
	"testing"

	"github.com/spf13/viper"
)

func TestCheckMutuallyExclusiveSettings(t *testing.T) {
	config := viper.New()
	got := CheckMutuallyExclusiveSettings(config)
	if got != nil {
		t.Errorf("got: %s, want: %s", got, "")
	}
}

func TestCheckMutuallyExclusiveSettingsFail(t *testing.T) {
	config := viper.New()
	config.Set("image", "test")
	config.Set("dockerfile", "test")
	got := CheckMutuallyExclusiveSettings(config)
	if got == nil {
		t.Errorf("got: %s, want: %s", got, "error")
	}
}

func TestRemoveFromSlice(t *testing.T) {
	got := RemoveFromSlice([]string{"a", "b", "c", "c"}, "c")
	want := []string{"a", "b"}
	for i, v := range got {
		if v != want[i] {
			t.Errorf("got: %s, want: %s", got, want)
		}
	}
}

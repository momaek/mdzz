package mdzz

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	sm := NewSafetyMap()
	require.NotNil(t, sm)
}

func TestGet(t *testing.T) {
	sm := NewSafetyMap()
	require.NotNil(t, sm)

	sm.Set("key", "value")

	value, ok := sm.Get("key")
	require.Equal(t, true, ok)
	require.NotNil(t, value)

	s, ok := value.(string)
	require.Equal(t, true, ok)
	require.Equal(t, "value", s)
}

func TestSet(t *testing.T) {
	sm := NewSafetyMap()
	require.NotNil(t, sm)

	sm.Set("key", "value")
	_, ok := sm.Get("key")
	require.Equal(t, true, ok)
}

func TestHas(t *testing.T) {
	sm := NewSafetyMap()
	require.NotNil(t, sm)

	sm.Set("key", "value")
	has := sm.Has("key")
	require.Equal(t, true, has)

	has = sm.Has("a")
	require.Equal(t, false, has)
}

func TestDelete(t *testing.T) {
	sm := NewSafetyMap()
	require.NotNil(t, sm)

	sm.Set("key", "value")
	has := sm.Has("key")
	require.Equal(t, true, has)
	sm.Delete("key")
	has = sm.Has("key")
	require.Equal(t, false, has)
}

package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConvertFieldValueInterfaceToString(t *testing.T) {
	var val string
	var err error

	val, err = ConvertFieldValueInterfaceToString("hello", EventFieldString)
	require.NoError(t, err)
	require.NotEmpty(t, val)

	val, err = ConvertFieldValueInterfaceToString(true, EventFieldBoolean)
	require.NoError(t, err)
	require.NotEmpty(t, val)

	val, err = ConvertFieldValueInterfaceToString("hello", EventFieldBoolean)
	require.Error(t, err)
	require.Empty(t, val)

	val, err = ConvertFieldValueInterfaceToString(float64(1), EventFieldNumber)
	require.NoError(t, err)
	require.NotEmpty(t, val)

}

package template

import (
	"os"
	"testing"

	"github.com/Masterminds/sprig/v3"
	"github.com/stretchr/testify/require"
)

func TestParseFile(t *testing.T) {
	data, err := os.ReadFile("testdata/flowschema.yaml")
	require.NoError(t, err)
	tpl, err := New("flowschema").Funcs(FuncMap(sprig.FuncMap())).Parse(string(data))
	require.NoError(t, err)
	err = tpl.Execute(os.Stdout, nil)
	require.NoError(t, err)
}

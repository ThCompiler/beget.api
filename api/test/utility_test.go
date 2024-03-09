package test

import (
	"net/url"
	"testing"

	"github.com/ThCompiler/go.beget.api/core"
	"github.com/ThCompiler/go.beget.api/pkg/maps"
	"github.com/stretchr/testify/require"
)

func requireEqualValues(t *testing.T, expected, actual url.Values, client core.Client, skipValues ...string) {
	t.Helper()

	expected.Add("login", client.Login)
	expected.Add("passwd", client.Password)
	expected.Add("output_format", string(core.JSON))

	cpExpected := maps.Clone(expected)
	cpActual := maps.Clone(actual)

	for _, values := range skipValues {
		cpExpected.Del(values)
		cpActual.Del(values)
	}

	require.EqualValues(t, cpExpected, cpActual)
}

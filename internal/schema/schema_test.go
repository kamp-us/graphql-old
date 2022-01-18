package schema_test

import (
	"testing"

	"github.com/kamp-us/graphql/internal/schema"
	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	s, err := schema.String()

	require.NoError(t, err)
	require.NotEmpty(t, s)
}

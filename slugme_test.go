package slugme

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test1(t *testing.T) {
	require := require.New(t)

	title, err := New(DefaultOptions)
	require.NoError(err)
	require.NotNil(title)

	slug := title.Slug(" L'oiseau à   deux becs ")
	require.Equal("l-oiseau-a-deux-becs", slug)
}

func Test2(t *testing.T) {
	require := require.New(t)

	opts := Options{
		Allowed:  "_-+*/",
		Replace:  "",
		KeepCase: true,
	}

	ref, err := New(opts)
	require.NoError(err)
	require.NotNil(ref)

	slug := ref.Slug("MF   218 F/A _ LIMF 218 FA")
	require.Equal("MF218F/A_LIMF218FA", slug)
}

func Test3(t *testing.T) {
	require := require.New(t)

	title, err := New(DefaultOptions)
	require.NoError(err)
	require.NotNil(title)

	slug := title.Slug("\u019a")
	require.Equal("l", slug)

	slug = title.Slug("")
	require.Equal("", slug)

}

func Test4(t *testing.T) {
	require := require.New(t)

	title, err := New(DefaultOptions)
	require.NoError(err)
	require.NotNil(title)

	slug := title.Slug("€€€€  €€€€")
	require.Equal("", slug)
}

func Test5(t *testing.T) {
	require := require.New(t)

	title, err := New(DefaultOptions)
	require.NoError(err)
	require.NotNil(title)

	slug := title.Slug("$$$hello$$$")
	require.Equal("hello", slug)

	slug = title.Slug("$$$hello$$$world€€€")
	require.Equal("hello-world", slug)
}

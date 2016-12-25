package opts

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewTagSet_Basic(t *testing.T) {
	set := NewTagSet(`default:"true"`)
	require.Equal(t, 1, len(set))
	require.Equal(t, set["default"], "true")
}

func TestNewTagSet_Multiline(t *testing.T) {
	set := NewTagSet(`
        default:"true"
        description:"Use verbose logging."
        help:"Be very talkative when logging"
        long:"verbose"
        short:"v"`)
	require.Equal(t, 5, len(set))
	require.Equal(t, set["default"], "true")
	require.Equal(t, set["description"], "Use verbose logging.")
	require.Equal(t, set["help"], "Be very talkative when logging")
	require.Equal(t, set["long"], "verbose")
	require.Equal(t, set["short"], "v")
}

func TestNewTagSet_MultilineValue(t *testing.T) {
	set := NewTagSet(`
        default:"true"
        description:"Use verbose
        logging."
        help:"Be very talkative when logging"
        long:"verbose"
        short:"v"`)
	require.Equal(t, 5, len(set))
	require.Equal(t, set["default"], "true")
	require.Equal(t, set["description"], "Use verbose\n        logging.")
	require.Equal(t, set["help"], "Be very talkative when logging")
	require.Equal(t, set["long"], "verbose")
	require.Equal(t, set["short"], "v")
}

func TestNewTagSet_ResourceUrls(t *testing.T) {
	set := NewTagSet(`
        long:"db"
        default:"mongodb://localhost:27017/db"
        description:"The db resource to connect to."`)
	require.Equal(t, 3, len(set))
	require.Equal(t, set["default"], "mongodb://localhost:27017/db")
	require.Equal(t, set["description"], "The db resource to connect to.")
	require.Equal(t, set["long"], "db")
}

func TestTagSetGet(t *testing.T) {
	set := NewTagSet(`default:"true"`)
	require.Equal(t, set.Get("default"), "true")
}

func TestTagSetHas(t *testing.T) {
	set := NewTagSet(`default:"true"`)
	require.True(t, set.Has("default"))
}

package opts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTagSet_Basic(t *testing.T) {
	set := NewTagSet(`default:"true"`)
	assert.Equal(t, 1, len(set))
	assert.Equal(t, set["default"], "true")
}

func TestNewTagSet_Multiline(t *testing.T) {
	set := NewTagSet(`
        default:"true"
        description:"Use verbose logging."
        help:"Be very talkative when logging"
        long:"verbose"
        short:"v"`)
	assert.Equal(t, 5, len(set))
	assert.Equal(t, set["default"], "true")
	assert.Equal(t, set["description"], "Use verbose logging.")
	assert.Equal(t, set["help"], "Be very talkative when logging")
	assert.Equal(t, set["long"], "verbose")
	assert.Equal(t, set["short"], "v")
}

func TestNewTagSet_MultilineValue(t *testing.T) {
	set := NewTagSet(`
        default:"true"
        description:"Use verbose
        logging."
        help:"Be very talkative when logging"
        long:"verbose"
        short:"v"`)
	assert.Equal(t, 5, len(set))
	assert.Equal(t, set["default"], "true")
	assert.Equal(t, set["description"], "Use verbose\n        logging.")
	assert.Equal(t, set["help"], "Be very talkative when logging")
	assert.Equal(t, set["long"], "verbose")
	assert.Equal(t, set["short"], "v")
}

func TestNewTagSet_ResourceUrls(t *testing.T) {
	set := NewTagSet(`
        long:"db"
        default:"mongodb://localhost:27017/db"
        description:"The db resource to connect to."`)
	assert.Equal(t, 3, len(set))
	assert.Equal(t, set["default"], "mongodb://localhost:27017/db")
	assert.Equal(t, set["description"], "The db resource to connect to.")
	assert.Equal(t, set["long"], "db")
}

func TestTagSetGet(t *testing.T) {
	set := NewTagSet(`default:"true"`)
	assert.Equal(t, set.Get("default"), "true")
}

func TestTagSetHas(t *testing.T) {
	set := NewTagSet(`default:"true"`)
	assert.True(t, set.Has("default"))
}

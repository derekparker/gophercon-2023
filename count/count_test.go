package count

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseURLsFromFile(t *testing.T) {
	jobs, err := ParseURLsFromFile("./testdata/urls.txt")
	assert.Nil(t, err)

	// Assert the URLs are parsed in the correct order as we expect.
	assert.Equal(t, "http://bing.com", jobs[2].Url)
	assert.Equal(t, "http://yahoo.com", jobs[0].Url)
	assert.Equal(t, "http://go.dev", jobs[1].Url)
}

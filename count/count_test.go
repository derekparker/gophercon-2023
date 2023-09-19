package count

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseURLsFromFile(t *testing.T) {
	jobs, err := ParseURLsFromFile("./testdata/urls.txt")
	assert.Nil(t, err)

	fmt.Printf("%#v\n", jobs)

	// Assert the URLs are parsed in the correct order as we expect.
	assert.Equal(t, "http://google.com", jobs[0].url)
	assert.Equal(t, "http://yahoo.com", jobs[1].url)
	assert.Equal(t, "http://go.dev", jobs[2].url)
}

package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const UserId = "e0dba740-fc4b-4977-872c-d360239e6b1a"

// TestShortLinkGenerator runs the generator on three different URLs
// and asserts that the generated short links match the expected outputs.
func TestShortLinkGenerator(t *testing.T) {
	initialLink1 := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	shortLink1 := GenerateShortLink(initialLink1, UserId)

	initialLink2 := "https://www.eddywm.com/lets-build-a-url-shortener-in-go-with-redis-part-2-storage-layer/"
	shortLink2 := GenerateShortLink(initialLink2, UserId)

	initialLink3 := "https://spectrum.ieee.org/automaton/robotics/home-robots/hello-robots-stretch-mobile-manipulator"
	shortLink3 := GenerateShortLink(initialLink3, UserId)

	// Assert that the generated short links match the expected outputs.
	assert.Equal(t, "jTa4L57P", shortLink1)
	assert.Equal(t, "d66yfx7N", shortLink2)
	assert.Equal(t, "dhZTayYQ", shortLink3)
}

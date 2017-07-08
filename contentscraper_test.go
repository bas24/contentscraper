package contentscraper

import (
	"github.com/stretchr/testify/assert"
	"testing"
	// "golang.org/x/net/html"
)

func TestScrape(t *testing.T) {
	const txt string = `
		<div itemprop="articleBody" class="article">
			<p>Some cool text.</p>
			<p>Which we are going to scrape.</p>
		</div>
		`
	assert.Equal(t, txt, txt, "Are equal")
	// TODO rewrite code for testing
	// without need to fetch real url

}

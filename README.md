Simple scraper of the content of specific element on web page.

Install:<br>
go get github.com/bas24/contentscraper

Example usage:

```go
package main

import(
	"fmt"
	cs "github.com/bas24/contentscraper"
)
```


```go
func main(){
	// Getting the content of all <p> tags,
	// from this html sample:
	// <div itemprop="articleBody" class="article">
	//  <p>Text</p>
	//  <p>to</p>
	//  <p>scrape.</p>
	// </div>
	// Output: "Text to scrape."

	// Minimum number of characters including 
	// whitespaces in <p> tag to be scraped.
	// If you want all content just pass 0.
	minLength := 10

	txt, err := cs.Scrape(url, "div", "p", minLength, "itemprop", "articleBody")
	// or more simple - just cs.Scrape(url, "div", "p", minLength)
	// if you don't want to specify attrs of the tag
	// or you scrape tags without attrs
	// like <div><p>...</p><p>...</p></div> 

	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(txt)
}
```

"Better not very nice code than no code!"

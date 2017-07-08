package contentscraper

import (
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"strings"
	"unicode"
)

func Scrape(url, elem, tag string, length int, attrs ...string) (string, error) {
	var articleContent []string
	resp, err := http.Get(url)
	if err != nil {
		return "err", errors.New(fmt.Sprintf("Error fetching url: %v", url))
	} else {
		htmlD, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "err", errors.New("Error while reading responce body.")
		}

		txt, err := html.Parse(strings.NewReader(replTxt(string(htmlD))))
		if err != nil {
			return "err", errors.New("Error while parsing html.")
		}

		var f func(*html.Node, bool)
		f = func(n *html.Node, printText bool) {
			if printText && n.Type == html.TextNode && len(n.Data) > length {
				articleContent = append(articleContent, n.Data)
			}
			hasProp := func() bool {
				if n.Parent.Type == html.ElementNode && n.Parent.Data == elem {
					if len(attrs) == 2 {
						for _, a := range n.Parent.Attr {
							if a.Key == attrs[0] && a.Val == attrs[1] {
								return true
							}
						}
					}
					if len(attrs) == 0 {
						return true
					}
				}
				return false
			}

			isPtag := func() bool {
				if n.Type == html.ElementNode && n.Data == tag {
					return true
				}
				return false
			}

			printText = printText || isPtag() && hasProp()
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c, printText)
			}
		}
		f(txt, false)
	}
	// join to get full text with
	// article id first in one string
	art := strings.Join(articleContent, " ")

	if len(art) > 0 {
		if unicode.IsLower(rune(art[0])) || string(art[0]) == " " {
			for k, v := range art {
				// Loop utill first sentence end
				if v == '.' {
					// Trim string left
					art = art[k+2:]
					break
				}
			}
		}
	}

	// cut additional whitespaces if any
	art = sliceFormatter(articleContent)

	return art, nil
}

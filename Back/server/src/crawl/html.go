package crawl

import (
	"golang.org/x/net/html"

	"regexp"
	"strconv"
	"strings"
)

func getElementById(node *html.Node, id string) *html.Node {
	return traverse(node, id)
}

func nextNode(node *html.Node, step int) *html.Node {
	e := step * 2
	next := node
	for i := 0; i < e; i++ {
		next = next.NextSibling
	}
	return next
}

func extractCountryName(node *html.Node) (countryName string, ok bool) {
	if countryName, ok = extractContent(node); !ok {
		return countryName, false
	}
	countryName = cleanText(countryName)
	if countryName == "" {
		return countryName, false
	}
	return countryName, true
}

func extractPercentage(node *html.Node) (percentage float64, ok bool) {
	var content string
	if content, ok = extractContent(node); !ok {
		return percentage, false
	}
	content = cleanNumber(content)
	content = strings.ReplaceAll(content, "%", "")
	if content == "" {
		return percentage, false
	}
	if per, err := strconv.ParseFloat(content, 32); err != nil {
		return percentage, false
	} else {
		return per, true
	}
}

func extractInteger(node *html.Node, zeroReplace bool) (number int64, ok bool) {
	var content string
	if content, ok = extractContent(node); !ok {
		return replaceIntWithZeroIf(number, zeroReplace)
	}
	content = cleanNumber(content)
	if content == "" {
		return replaceIntWithZeroIf(number, zeroReplace)
	}
	if num, err := strconv.Atoi(content); err != nil {
		return number, false
	} else {
		return int64(num), true
	}
}

// ----- helpers

func traverse(node *html.Node, id string) *html.Node {
	if node.Type == html.ElementNode {
		if val, ok := getAttribute(node, "id"); ok && val == id {
			return node
		}
	}
	for n := node.FirstChild; n != nil; n = n.NextSibling {
		if result := traverse(n, id); result != nil {
			return result
		}
	}
	return nil
}

func extractContent(node *html.Node) (content string, ok bool) {
	a := node.FirstChild
	if a == nil {
		return content, false
	}
	if inner := a.FirstChild; inner == nil {
		content = a.Data
	} else {
		content = inner.Data
	}
	return content, true
}

func cleanText(str string) string {
	pattern1 := regexp.MustCompile("[\\d.,?!']+")
	pattern2 := regexp.MustCompile("\\s+")
	str = pattern1.ReplaceAllString(str, "")
	str = pattern2.ReplaceAllString(str, " ")
	str = strings.TrimSpace(str)
	return str
}

func cleanNumber(str string) string {
	str = strings.ReplaceAll(str, ",", "")
	str = strings.ReplaceAll(str, " ", "")
	str = strings.TrimSpace(str)
	return str
}

func getAttribute(node *html.Node, key string) (string, bool) {
	for _, attr := range node.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}

func replaceIntWithZeroIf(number int64, doReplace bool) (int64, bool) {
	if doReplace {
		return 0, true
	} else {
		return number, false
	}
}

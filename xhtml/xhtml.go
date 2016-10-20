package xhtml

import (
	"errors"
	"math"

	"golang.org/x/net/html"
)

// NextSibling Fetch the next Tag sibling
func NextSibling(node *html.Node) (*html.Node, error) {
	return NthSibling(node, 1)
}

// NthSibling Fetch the next Tag sibling
func NthSibling(node *html.Node, amount int) (*html.Node, error) {
	for amount != 0 {
		if amount < 0 {
			node = node.PrevSibling
		} else {
			node = node.NextSibling
		}

		if node == nil {
			return nil, errors.New("Node not found!")
		}

		if node.Type == html.ElementNode {
			amount = amount - int(math.Copysign(1, float64(amount)))
		}
	}

	return node, nil
}

// FirstChild Fetch the first Tag child
func FirstChild(node *html.Node) (*html.Node, error) {
	node = node.FirstChild
	if node.Type == html.ElementNode {
		return node, nil
	}

	return NextSibling(node)
}

// MustNextSibling Fetch the next Tag sibling
func MustNextSibling(node *html.Node) *html.Node {
	n, err := NextSibling(node)
	if err != nil {
		panic(err)
	}

	return n
}

// MustNthSibling Fetch the nth Tag sibling
func MustNthSibling(node *html.Node, amount int) *html.Node {
	n, err := NthSibling(node, amount)
	if err != nil {
		panic(err)
	}

	return n
}

// MustFirstChild Fetch the first Tag child
func MustFirstChild(node *html.Node) *html.Node {
	n, err := FirstChild(node)
	if err != nil {
		panic(err)
	}

	return n
}

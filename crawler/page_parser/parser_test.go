package page_parser

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse1(t *testing.T) {
	r := bytes.NewReader([]byte(html1))
	if result, err := Parse(r); assert.NoError(t, err) {
		assert.Equal(t, "http://example.com/foo/bar", result.CanonicalURL)
		assert.Equal(t, "http://example.com/foo/bar/", result.baseURL)
		assert.Equal(t, []string{
			"http://example.com/",
			"http://example.com/foo/bar/page.html",
			"http://example.com/foo/",
			"http://some.other.com",
			"javascript:void(0)",
		}, result.Links)
	}
}

func TestParse2(t *testing.T) {
	r := bytes.NewReader([]byte(html2))
	if result, err := Parse(r); assert.NoError(t, err) {
		assert.Equal(t, "", result.CanonicalURL)
		assert.Equal(t, "", result.baseURL)
		assert.Equal(t, []string{
			"/",
			"page.html",
			"..",
			"http://some.other.com",
			"javascript:void(0)",
		}, result.Links)
	}
}

// language=HTML
const html1 = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
	<base href="http://example.com/foo/bar/">
	<link rel="canonical" href="http://example.com/foo/bar ">
    <title>Foo Bar</title>
</head>
<body>
	<a href="/">Home</a>
	<a href="page.html">Some page</a>
	<a href="..">Up</a>
	<a href="http://some.other.com">External link</a>
	<a href=" javascript:void(0)">Click here!</a>
</body>
</html>
`

// language=HTML
const html2 = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <title>Foo Bar</title>
	<base href="">
</head>
<body>
	<a href="/">Home</a>
	<a href="page.html">Some page</a>
	<a href="..">Up</a>
	<a href="http://some.other.com">External link</a>
	<a href=" javascript:void(0)">Click here!</a>
	<a href="">Empty?</a>
	<a href="#local">Local Reference</a>
</body>
</html>
`

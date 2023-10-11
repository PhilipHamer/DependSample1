package main

import (
	"fmt"

	"github.com/cespare/xxhash"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

var mds = `# header

Sample text.

[link](http://example.com)
`

func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

func main() {
	fmt.Println("hello world")

	hash := xxhash.New()
	hash.Write([]byte{1, 2, 3, 4})
	foo := hash.Sum64()

	md := []byte(mds)
	html := mdToHTML(md)

	fmt.Printf("--- Markdown:\n%s\n\n--- HTML:\n%s\n", md, html)

	fmt.Println(foo)
}

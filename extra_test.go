package goldmark_test

import (
	"bytes"
	"testing"

	. "github.com/anytypeio/goldmark"
	"github.com/anytypeio/goldmark/ast"
	"github.com/anytypeio/goldmark/parser"
	"github.com/anytypeio/goldmark/renderer/html"
	"github.com/anytypeio/goldmark/testutil"
)

func TestExtras(t *testing.T) {
	markdown := New(WithRendererOptions(
		html.WithXHTML(),
		html.WithUnsafe(),
	))
	testutil.DoTestCaseFile(markdown, "_test/extra.txt", t)
}

func TestEndsWithNonSpaceCharacters(t *testing.T) {
	markdown := New(WithRendererOptions(
		html.WithXHTML(),
		html.WithUnsafe(),
	))
	source := []byte("```\na\n```")
	var b bytes.Buffer
	err := markdown.Convert(source, &b)
	if err != nil {
		t.Error(err.Error())
	}
	if b.String() != "<pre><code>a\n</code></pre>\n" {
		t.Errorf("%s \n---------\n %s", source, b.String())
	}
}

type BlocksWriter interface {

}

func TestWindowsNewLine(t *testing.T) {
	markdown := New(WithRendererOptions(
		html.WithXHTML(),
	))
	source := []byte("a  \r\nb\n")
	var b bytes.Buffer


	err := markdown.Convert(source, &b)
	if err != nil {
		t.Error(err.Error())
	}
	if b.String() != "<p>a<br />\nb</p>\n" {
		t.Errorf("%s\n---------\n%s", source, b.String())
	}

	source = []byte("a\\\r\nb\r\n")
	var b2 bytes.Buffer
	err = markdown.Convert(source, &b2)
	if err != nil {
		t.Error(err.Error())
	}
	if b2.String() != "<p>a<br />\nb</p>\n" {
		t.Errorf("\n%s\n---------\n%s", source, b2.String())
	}
}

type myIDs struct {
}

func (s *myIDs) Generate(value []byte, kind ast.NodeKind) []byte {
	return []byte("my-id")
}

func (s *myIDs) Put(value []byte) {
}

func TestAutogeneratedIDs(t *testing.T) {
	ctx := parser.NewContext(parser.WithIDs(&myIDs{}))
	markdown := New(WithParserOptions(parser.WithAutoHeadingID()))
	source := []byte("# Title1\n## Title2")
	var b bytes.Buffer
	err := markdown.Convert(source, &b, parser.WithContext(ctx))
	if err != nil {
		t.Error(err.Error())
	}
	if b.String() != `<h1 id="my-id">Title1</h1>
<h2 id="my-id">Title2</h2>
` {
		t.Errorf("%s\n---------\n%s", source, b.String())
	}
}

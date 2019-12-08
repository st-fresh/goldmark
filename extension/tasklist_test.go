package extension

import (
	"testing"

	"github.com/enkogu/goldmark"
	"github.com/enkogu/goldmark/renderer/html"
	"github.com/enkogu/goldmark/testutil"
)

func TestTaskList(t *testing.T) {
	markdown := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			TaskList,
		),
	)
	testutil.DoTestCaseFile(markdown, "_test/tasklist.txt", t)
}

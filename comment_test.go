package kss

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCommentParser(t *testing.T) {
	comments := CommentParser("fixtures/comments.txt")

	// finds single line comment styles
	assert.Equal(t, comments[0], strings.TrimSpace(`
This comment block has comment identifiers on every line.

Fun fact: this is Kyle's favorite comment syntax!
    `))

	// finds block style comment styles
	assert.Equal(t, comments[1], strings.TrimSpace(`
This comment block is a block-style comment syntax.

There's only two identifier across multiple lines.
    `))

	assert.Equal(t, comments[2], strings.TrimSpace(`
This is another common multi-line comment style.

It has stars at the begining of every line.
    `))

	// handles mixed styles
	assert.Equal(t, comments[3], strings.TrimSpace(`
This comment has a /* comment */ identifier inside of it!
    `))

	assert.Equal(t, comments[4], strings.TrimSpace(`
Look at my //cool// comment art!
    `))

	// handles indented comments
	assert.Equal(t, comments[5], strings.TrimSpace(`
Indented single-line comment.
    `))

	assert.Equal(t, comments[6], strings.TrimSpace(`
Indented block comment.
    `))

	// handles indented example
	assert.Equal(t, comments[7], strings.TrimSpace(`
This comment has a indented example
<div>
    <div></div>
</div>
    `))

	// handles unicode in comments
	assert.Equal(t, comments[8], strings.TrimSpace(`
Hello, 世界
    `))
}

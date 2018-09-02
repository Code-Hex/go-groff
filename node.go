package groff

type Node interface {
	Token() string
}

type CommentNode struct {
	token   string
	comment string
}

func NewCommentNode(comment string) *CommentNode {
	return &CommentNode{
		token:   Comment,
		comment: comment,
	}
}
func (c *CommentNode) Token() string   { return c.token }
func (c *CommentNode) Comment() string { return c.comment }

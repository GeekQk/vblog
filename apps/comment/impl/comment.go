package impl

import (
	"context"

	"github.com/GeekQk/vblog/apps/comment"
)

func (i *commentServiceImpl) CreateComment(
	ctx context.Context,
	in *comment.CreateCommentRequest) (
	*comment.Comment, error) {
	return &comment.Comment{
		Spec: in,
	}, nil
}

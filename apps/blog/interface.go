package blog

import (
	"context"

	"github.com/go-playground/validator"
)

const (
	AppName = "blog"
)

var (
	v = validator.New()
)

// Blog Service定义
type Service interface {
	// 创建博客
	CreateBlog(context.Context, *CreateBlogRequest) (*Blog, error)
	// 博客列表
	QueryBlog(context.Context, *QueryBlogRequest) (*BlogSet, error)
	// 博客详情
	DescribeBlog(context.Context, *DescribeBlogRequest) (*Blog, error)
	// 更新博客
	UpdateBlog(context.Context, *UpdateBlogRequest) (*Blog, error)
	// 删除博客
	DeleteBlog(context.Context, *DeleteBlogRequest) (*Blog, error)
	//文章发布
	ChangedBlogStatus(context.Context, *ChangedBlogStatusRequest) (*Blog, error)
	//审核接口
	AuditBlog(context.Context, *AuditInfo) (*Blog, error)
}

func NewQueryBlogRequest() *QueryBlogRequest {
	return &QueryBlogRequest{
		PageSize:   20,
		PageNumber: 1,
	}
}

type QueryBlogRequest struct {
	//分页大小，一页多少个
	PageSize int
	//当前页
	PageNumber int
}

func (req *QueryBlogRequest) Limit() int {
	return req.PageSize
}
func (req *QueryBlogRequest) Offset() int {
	return req.PageSize * (req.PageNumber - 1)
}

func NewDescribeUserRequest(id string) *DescribeBlogRequest {
	return &DescribeBlogRequest{
		Id: id,
	}
}

type DescribeBlogRequest struct {
	Id string
}

type UpdateBlogRequest struct {
}

func NewDeleteBlogRequest(id string) *DeleteBlogRequest {
	return &DeleteBlogRequest{
		Id: id,
	}
}

type DeleteBlogRequest struct {
	Id string
}

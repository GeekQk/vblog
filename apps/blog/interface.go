package blog

import (
	"context"
)

var (
	AppName = "blog"
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

type QueryBlogRequest struct {
}

type DescribeBlogRequest struct {
}

type UpdateBlogRequest struct {
}

type DeleteBlogRequest struct {
}

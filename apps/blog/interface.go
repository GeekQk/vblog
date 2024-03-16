package blog

import (
	"context"
	"strconv"

	"github.com/GeekQk/vblog/common"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

const (
	AppName = "blogs"
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
	//创建人
	CreateBy string
	//关键字
	Keywords string
}

func (req *QueryBlogRequest) Limit() int {
	return req.PageSize
}
func (req *QueryBlogRequest) Offset() int {
	return req.PageSize * (req.PageNumber - 1)
}

func NewDescribeBlogRequest(id string) *DescribeBlogRequest {
	return &DescribeBlogRequest{
		Id: id,
	}
}

type DescribeBlogRequest struct {
	Id string
}

func NewUpdateBlogRequest(id string) *UpdateBlogRequest {
	return &UpdateBlogRequest{
		Id:                id,
		UpdateMode:        common.UPDATE_MODE_PUT,
		CreateBlogRequest: NewCreateBlogRequest(),
	}
}

type UpdateBlogRequest struct {
	//更新ID
	Id string `json:"id"`
	//更新模式
	UpdateMode common.UpdateMode `json:"update_mode"`
	//更新的数据
	*CreateBlogRequest
}

func NewDeleteBlogRequest(id string) *DeleteBlogRequest {
	return &DeleteBlogRequest{
		Id: id,
	}
}

type DeleteBlogRequest struct {
	Id string
}

func NewQueryBlogRequestFromGin(c *gin.Context) *QueryBlogRequest {
	req := NewQueryBlogRequest()
	ps := c.Query("page_size") //返回字符串
	//增加创建人信息
	req.CreateBy = c.Query("create_by")
	//增加关键字
	req.Keywords = c.Query("keywords")
	//字符串统一转int
	if ps != "" {
		req.PageSize, _ = strconv.Atoi(ps)
	}
	pn := c.Query("page_number")
	//字符串统一转int
	if pn != "" {
		req.PageNumber, _ = strconv.Atoi(pn)
	}
	return req
}

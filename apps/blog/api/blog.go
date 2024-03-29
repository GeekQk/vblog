package api

import (
	"github.com/GeekQk/vblog/apps/blog"
	"github.com/GeekQk/vblog/apps/token"
	"github.com/GeekQk/vblog/common"
	"github.com/GeekQk/vblog/response"
	"github.com/gin-gonic/gin"
)

// + 创建博客: POST /vblogs/api/v1/blogs
func (h *blogApiHandler) CreateBlog(c *gin.Context) {
	// h.tk.Validate()
	req := blog.NewCreateBlogRequest()

	// 获取登录信息
	if v, ok := c.Get(token.TOKEN_MIDDLEWARE_KEY); ok {
		req.CreateBy = v.(*token.Token).UserName
	}

	if err := c.BindJSON(req); err != nil {
		response.Failed(c, err)
		return
	}
	ins, err := h.svc.CreateBlog(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, ins)
}

// + 修改博客(部分): PATCH /vblogs/api/v1/blogs/:id
// /vblogs/api/v1/blogs/10 --> id = 10
// /vblogs/api/v1/blogs/20 --> id = 20
// c.Param("id") 获取路径变量的值
func (h *blogApiHandler) PatchBlog(c *gin.Context) {
	// h.tk.Validate()

	// 如果解析路径里面的参数
	req := blog.NewUpdateBlogRequest(c.Param("id"))
	req.UpdateMode = common.UPDATE_MODE_PATCH
	// 用户传递的数据
	if err := c.BindJSON(req.CreateBlogRequest); err != nil {
		response.Failed(c, err)
		return
	}
	// 后面请求如何获取 中间信息
	if v, ok := c.Get(token.TOKEN_MIDDLEWARE_KEY); ok {
		req.CreateBy = v.(*token.Token).UserName
	}
	ins, err := h.svc.UpdateBlog(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, ins)
}

// + 修改博客(全量): PUT /vblogs/api/v1/blogs/:id
func (h *blogApiHandler) UpdateBlog(c *gin.Context) {
	// 如果解析路径里面的参数
	req := blog.NewUpdateBlogRequest(c.Param("id"))
	req.UpdateMode = common.UPDATE_MODE_PUT
	// 用户传递的数据
	if err := c.BindJSON(req.CreateBlogRequest); err != nil {
		response.Failed(c, err)
		return
	}
	// 后面请求如何获取 中间信息
	if v, ok := c.Get(token.TOKEN_MIDDLEWARE_KEY); ok {
		req.CreateBy = v.(*token.Token).UserName
	}
	ins, err := h.svc.UpdateBlog(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, ins)
}

// + 删除博客: DELETE /vblogs/api/v1/blogs/:id
func (h *blogApiHandler) DeleteBlog(c *gin.Context) {
	req := blog.NewDeleteBlogRequest(c.Param("id"))
	ins, err := h.svc.DeleteBlog(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, ins)
}

// + 查询列表: GET /vblogs/api/v1/blogs?page_size=10&page_number=2
// 获取url Get方式 c.Query("key")
func (h *blogApiHandler) QueryBlog(c *gin.Context) {
	req := blog.NewQueryBlogRequestFromGin(c)
	set, err := h.svc.QueryBlog(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, set)
}

// + 查询详情: GET /vblogs/api/v1/blogs/:id
func (h *blogApiHandler) DescribeBlog(c *gin.Context) {
	req := blog.NewDescribeBlogRequest(c.Param("id"))
	ins, err := h.svc.DescribeBlog(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, ins)
}

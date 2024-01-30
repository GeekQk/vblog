package impl_test

import (
	"testing"

	"github.com/GeekQk/vblog/apps/blog"
	"github.com/GeekQk/vblog/common"
)

func TestCreateBlog(t *testing.T) {
	req := blog.NewCreateBlogRequest()
	req.Title = "vblog Go语言web开发s"
	req.Author = "qk"
	req.Content = "xxxx"
	req.Summary = "xx"
	req.Tags["目录"] = "go"
	ins, err := impl.CreateBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestQueryBlogList(t *testing.T) {
	req := blog.NewQueryBlogRequest()
	ins, err := impl.QueryBlog(ctx, req)

	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestDetailBlogList(t *testing.T) {
	req := blog.NewDescribeUserRequest("47")
	ins, err := impl.DescribeBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestDeleteBlogList(t *testing.T) {
	req := blog.NewDeleteBlogRequest("47")
	ins, err := impl.DeleteBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestUpdateBlogList(t *testing.T) {
	req := blog.NewUpdateBlogRequest("50")
	req.Title = "牛呀无级别"
	req.Content = "拉拉肥但开发"
	req.UpdateMode = common.UPDATE_MODE_PATCH
	ins, err := impl.UpdateBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestUpdateReplaceBlogList(t *testing.T) {
	req := blog.NewUpdateBlogRequest("50")
	req.Title = "牛呀无级别"
	req.Content = "牛呀无级别"
	req.Author = "qq"
	req.UpdateMode = common.UPDATE_MODE_PUT
	ins, err := impl.UpdateBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

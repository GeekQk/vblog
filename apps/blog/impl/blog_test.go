package impl_test

import (
	"testing"

	"github.com/GeekQk/vblog/apps/blog"
)

func TestCreateBlog(t *testing.T) {
	req := blog.NewCreateBlogRequest()
	req.Title = "vblog Go语言web开发"
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

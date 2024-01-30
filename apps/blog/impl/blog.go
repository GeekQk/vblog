package impl

import (
	"context"

	"github.com/GeekQk/vblog/apps/blog"
	"github.com/GeekQk/vblog/exception"
)

// 创建博客
func (i *blogServiceImpl) CreateBlog(ctx context.Context, req *blog.CreateBlogRequest) (*blog.Blog, error) {
	// 1. 校验请求
	if err := req.Validate(); err != nil {
		return nil, exception.ErrBadRequest.WithMessagef("创建博客失败, %s", err)
	}

	// 2. 构造对象
	ins := blog.NewBlog(req)

	// 3. 对象入库
	// INSERT INTO `blogs` (`created_at`,`updated_at`,`title`,`author`,`content`,`summary`,`create_by`,`tags`,`published_at`,`status`,`audit_at`,`is_audit_pass`) VALUES (1706340774,1706340774,'go语言全栈开发','oldyu','xxx','xx','','{"目录":"Go语言"}',0,'0',0,false)
	err := i.db.WithContext(ctx).Create(ins).Error
	if err != nil {
		return nil, err
	}

	// 4. 返回对象
	return ins, nil
}

// 博客列表
func (i *blogServiceImpl) QueryBlog(ctx context.Context, req *blog.QueryBlogRequest) (*blog.BlogSet, error) {
	set := blog.NewBlogSet()

	// 1. 初始化查询对象
	query := i.db.WithContext(ctx).Model(blog.Blog{})

	// 查询总算
	err := query.Count(&set.Total).Error
	if err != nil {
		return nil, err
	}

	// 查询具体的数据
	err = query.
		Limit(req.Limit()).
		Offset(req.Offset()).
		Find(&set.Items).
		Error
	if err != nil {
		return nil, err
	}

	return set, nil
}

// 博客详情
func (i *blogServiceImpl) DescribeBlog(ctx context.Context, req *blog.DescribeBlogRequest) (*blog.Blog, error) {
	// 构造一个mysql 条件查询语句  select * from users where ....
	query := i.db.WithContext(ctx).Model(&blog.Blog{}).Where("id = ?", req.Id)

	// 准备一个对象 接收数据库的返回
	ins := blog.NewBlog(blog.NewCreateBlogRequest())
	if err := query.First(ins).Error; err != nil {
		return nil, err
	}
	return ins, nil

}

// 更新博客
func (i *blogServiceImpl) UpdateBlog(ctx context.Context, req *blog.UpdateBlogRequest) (*blog.Blog, error) {
	return nil, nil
}

// 删除博客
func (i *blogServiceImpl) DeleteBlog(ctx context.Context, req *blog.DeleteBlogRequest) (*blog.Blog, error) {
	ins, err := i.DescribeBlog(ctx, blog.NewDescribeUserRequest(req.Id))
	if err != nil {
		return nil, err
	}

	err = i.db.
		WithContext(ctx).
		Model(&blog.Blog{}).
		Where("id = ?", req.Id).
		Delete(ins).
		Error
	if err != nil {
		return nil, err
	}

	return ins, nil

}

// 文章发布
func (i *blogServiceImpl) ChangedBlogStatus(ctx context.Context, req *blog.ChangedBlogStatusRequest) (*blog.Blog, error) {
	return nil, nil
}

// 审核接口
func (i *blogServiceImpl) AuditBlog(ctx context.Context, req *blog.AuditInfo) (*blog.Blog, error) {
	return nil, nil
}

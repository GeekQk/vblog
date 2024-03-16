package impl

import (
	"context"

	"dario.cat/mergo"
	"github.com/GeekQk/vblog/apps/blog"
	"github.com/GeekQk/vblog/common"
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

	//补充查询条件
	if req.CreateBy != "" {
		query = query.Where("create_by = ?", req.CreateBy)
	}
	if req.Keywords != "" {
		query = query.Where("title like ?", "%"+req.Keywords+"%")
	}
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
// 1. 全新更新: 对象的替换
// 2. 部分更新: (old obj) --patch--> new obj ---> save
func (i *blogServiceImpl) UpdateBlog(ctx context.Context, req *blog.UpdateBlogRequest) (*blog.Blog, error) {
	// 查询老的对象, 需要被更新的博客对象
	ins, err := i.DescribeBlog(ctx, blog.NewDescribeBlogRequest(req.Id))
	if err != nil {
		return nil, err
	}
	// 对象更新
	switch req.UpdateMode {
	//部分更新
	case common.UPDATE_MODE_PATCH:
		// if req.Author != "" {
		// 	ins.Author = req.Author
		// }
		// if req.Title != "" {
		// 	ins.Title = req.Title
		// }
		//... 有没有其他的办法 帮我们完成2个结构体的合并 merge(patch)
		// https://github.com/darccio/mergo
		// go get dario.cat/mergo
		// WithOverride will make merge override non-empty dst attributes with non-empty src attributes values.
		//后者覆盖前者
		if err := mergo.MapWithOverwrite(ins.CreateBlogRequest, req.CreateBlogRequest); err != nil {
			return nil, err
		}
	default: //全量更新
		ins.CreateBlogRequest = req.CreateBlogRequest
	}
	// 再次校验对象, 校验更新后的数据是否合法
	if err := ins.Validate(); err != nil {
		return nil, exception.ErrBadRequest.WithMessagef("校验更新请求失败: %s", err)
	}

	// 更新数据库
	// UPDATE `blogs` SET `id`=48,`created_at`=1706344163,`updated_at`=1706344423,`title`='go语言全栈开发V2',`author`='oldyu',`content`='xxx',`summary`='xx',`tags`='{"目录":"Go语言"}' WHERE id = 48
	stmt := i.db.WithContext(ctx).Model(&blog.Blog{}).Where("id = ?", ins.Id)
	// 补充更新条件
	if req.CreateBy != "" {
		stmt = stmt.Where("create_by = ?", req.CreateBy)
	}
	err = stmt.Updates(ins).Error
	if err != nil {
		return nil, err
	}

	return ins, nil
}

// 删除博客
func (i *blogServiceImpl) DeleteBlog(ctx context.Context, req *blog.DeleteBlogRequest) (*blog.Blog, error) {
	ins, err := i.DescribeBlog(ctx, blog.NewDescribeBlogRequest(req.Id))
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

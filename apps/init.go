package apps

//业务实现
//只要在main中导入 "github.com/GeekQk/vblog/apps"
//这样就会自动使用引入模块的init方法

import (
	// 通过import方法 完成注册
	_ "github.com/GeekQk/vblog/apps/blog/api"
	_ "github.com/GeekQk/vblog/apps/blog/impl"
	_ "github.com/GeekQk/vblog/apps/comment/impl"
	_ "github.com/GeekQk/vblog/apps/token/api"
	_ "github.com/GeekQk/vblog/apps/token/impl"
	_ "github.com/GeekQk/vblog/apps/user/impl"
)

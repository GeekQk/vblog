# 评论模块

+ 对内提供RPC
+ 对外提供Restful 接口


## 业务定义

Comment
```
BlogId: 那个博客
User: 谁发起的评论
Content: 评论内容
CreateAt: 评论时间
```

```go
CreateComment(CreateCommentRequest) *Comment
```

通过编写Protobuf来定义业务接口
```proto3
service Service {
    rpc CreateComment(CreateCommentRequest) returns (Comment);
}

message Comment {
    // 评论Id
    string id = 1;
    // 10位时间戳
    int64 create_at = 2;
    // 评论具体定义
    CreateCommentRequest spec = 3;
}

message CreateCommentRequest {
    // 被评论文章Id
    string blog_id = 3;
    // 评论的用户
    string user_id = 4;
    // 评论内容
    string content = 5;
}
```

代码生成
```sh
protoc -I=. --go_out=. --go-grpc_out=. --go-grpc_opt=module="gitlab.com/go-course-project/go13/vblog" --go_opt=module="gitlab.com/go-course-project/go13/vblog" apps/*/pb/*.proto
```

```sh
protoc -I=. --go_out=. --go_opt=module="github.com/GeekQk/vblog" --go-grpc_out="."  --go-grpc_opt=module="github.com/GeekQk/vblog"  apps/comment/pb/*

```
## 无法嵌套, 导致数据多层

数据结构多层
```json
{id: '', create_at: 0, spec: {}}
```

1. 扁平化存储: 输入入库需要扁平化存储: GORM: 
+ embedded	嵌套字段
+ embeddedPrefix	嵌入字段的列名前缀

2. 扁平化展示: 使用匿名结构图, 重新调整数据结构, 数据在展示的时候重新组装:
```go
// {"id":"","create_at":0,"content":"test"}
jd2, _ := json.Marshal(struct {
    Id       string `json:"id"`
    CreateAt int64  `json:"create_at"`
    *comment.CreateCommnetRequest
}{
    Id:                   ins.Id,
    CreateAt:             ins.CreateAt,
    CreateCommnetRequest: ins.Spec,
})
```

## 无法添加自定义Tag

github.com/favadi/protoc-go-inject-tag

+ gorm
+ mongo
+ validate

安装插件:
```sh
go install github.com/favadi/protoc-go-inject-tag@latest
```

1. 补充Tag

```sh
protoc-go-inject-tag -input="*.pb.go"
```

2. 执行命令
```sh
protoc-go-inject-tag -input="apps/*/*.pb.go"
```

3. 为了方便make gen
```
gen: ## Gen code
	@protoc -I=. --go_out=. --go-grpc_out=. --go-grpc_opt=module="gitlab.com/go-course-project/go13/vblog" --go_opt=module="gitlab.com/go-course-project/go13/vblog" apps/*/pb/*.proto
	@protoc-go-inject-tag -input="apps/*/*.pb.go"

gen: ## make protobuf
	@protoc -I=. --go_out=. --go_opt=module="github.com/GeekQk/vblog" --go-grpc_out="."  --go-grpc_opt=module="github.com/GeekQk/vblog"  apps/*/pb/*.proto 
	@protoc-go-inject-tag -input="apps/*/*.pb.go"
```


## 接口最小权限设计

+ 进程内 调用: 权限最大, API Handler --> ServiceImpl
+ RPC 内部调用: 安全级别? 基本的安全, 数据修改类接口 适不适合 走RPC, 通过RPC删除文章? 是否允许通过RPC查询文章基本信息?
+ Restful API 调用: 调用的安全级别?  给 Web <-- 代表用户自己的行为

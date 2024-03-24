package comment_test

import (
	context "context"
	"encoding/json"
	"testing"

	"github.com/GeekQk/vblog/apps/comment"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestComment(t *testing.T) {
	ins := &comment.Comment{
		Spec: &comment.CreateCommentRequest{
			Content: "test",
		},
	}

	// {"spec":{"content":"test"}}
	jd, _ := json.Marshal(ins)
	t.Log(string(jd))

	//临时结构体组装数据
	// {"id":"","create_at":0,"content":"test"}
	jd2, _ := json.Marshal(struct {
		Id       string `json:"id"`
		CreateAt int64  `json:"create_at"`
		*comment.CreateCommentRequest
	}{
		Id:                   ins.Id,
		CreateAt:             ins.CreateAt,
		CreateCommentRequest: ins.Spec,
	})
	t.Log(string(jd2))
}

func TestRpcCreateComment(t *testing.T) {
	WithInsecure := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial("localhost:8301", WithInsecure)
	if err != nil {
		panic(err)
	}
	client := comment.NewServiceClient(conn)
	resp, err := client.CreateComment(context.Background(), &comment.CreateCommentRequest{
		Content: "test1",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)

}

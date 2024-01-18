package impl_test

import (
	"context"
	"fmt"
	"math/rand"
	"testing"

	"github.com/GeekQk/vblog/apps/user"
	"github.com/GeekQk/vblog/apps/user/impl"
)

var (
	i          *impl.UserServiceImpl
	ctx        = context.Background()
	firstNames = []string{"Alic", "Bobs", "Charlies", "Davids", "Emmas", "Franks", "Graces", "Henrys", "Ivys", "Jacks"}
	lastNames  = []string{"Smith", "Johnson", "Williams", "Jones", "Brown", "Davis", "Miller", "Wilson", "Moore", "Taylor"}
)

func generateRandomName() string {
	firstName := firstNames[rand.Intn(len(firstNames))]
	lastName := lastNames[rand.Intn(len(lastNames))]
	return fmt.Sprintf("%s %s", firstName, lastName)
}

// 测试创建用户
func TestCreateUser(t *testing.T) {
	req := user.NewCreateUserRequest()
	userList := []*user.User{}
	for index := 0; index < 3; index++ {
		req.Username = generateRandomName()
		req.Password = "11111"
		req.Role = user.ROLE_MEMBER
		u, err := i.CreateUser(ctx, req)
		//直接报错-单元测试失败
		if err != nil {
			t.Fatal(err)
		}
		userList = append(userList, u)
	}
	//正常打印对象
	t.Log(userList)

}

// 测试查询用户
func TestQueryUser(t *testing.T) {
	req := user.NewQueryUserRequest()
	req.PageNumber = 1
	req.PageSize = 10
	userSet, err := i.QueryUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(userSet)
}

// 根据用户Id查询用户
func TestDescribeUser(t *testing.T) {
	req := user.NewDescribeUserRequest(20)
	userSet, err := i.DescribeUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(userSet)
}

func init() {
	//加载测试对象 i就是UserServiceImpl
	i = impl.NewUserServiceImpl()
}

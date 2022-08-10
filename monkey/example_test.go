package monkey

import (
	"testing"

	"github.com/go-kiss/monkey"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestIsExistsUser(t *testing.T) {
	defer monkey.UnpatchAll()
	convey.Convey("测试用户存在", t, func() {
		// 这儿对应这 InstanceMethod  参考 https://github.com/go-kiss/monkey/blob/master/examples/instance_example.go
		monkey.PatchRaw((*Database).GetUserByID, func(_ *Database, id string) *User {
			return &User{
				ID:   id,
				Name: "example",
			}
		}, false, false)
		assert.Equal(t, true, IsExistsUser(new(Database), "1234"))
	})

	convey.Convey("测试用户不存在存在", t, func() {

		// 这儿对应这 InstanceMethod  参考 https://github.com/go-kiss/monkey/blob/master/examples/instance_example.go
		monkey.PatchRaw((*Database).GetUserByID, func(_ *Database, id string) *User {
			return nil
		}, false, false)
		assert.Equal(t, false, IsExistsUser(new(Database), "1234"))
	})

}

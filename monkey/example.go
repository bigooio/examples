package monkey

type Database struct {
}

type User struct {
	ID   string
	Name string
}

// GetUserByID 通过ID获取用户信息
func (d *Database) GetUserByID(id string) *User {
	// 查询数据库
	return nil
}

// IsExistsUser 判断用户是否存在
func IsExistsUser(db *Database, id string) bool {
	return db.GetUserByID(id) != nil
}

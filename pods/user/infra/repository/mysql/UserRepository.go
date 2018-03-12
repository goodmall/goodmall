package mysql

import (
	"fmt"
	// "fmt"
	"log"

	"github.com/go-ozzo/ozzo-dbx"
	_ "github.com/go-sql-driver/mysql"

	"github.com/goodmall/goodmall/pods/user/domain"
)

type userRepository struct {
	db *dbx.DB
}

// NewUserRepository 创建用户仓库的实现类
// TODO 对底层的数据库连接依赖 后期加入 先快速setup起来
func NewUserRepository( /* db *dbx.DB */ ) domain.UserRepository {

	db, _ := dbx.Open("mysql", "root:@/aheadmall")

	return &userRepository{
		db: db,
	}
}

/**
    FindByUsername(username string) (*User, error)
	// FindByEmail(email string) (*User, error)
	// FindByChangePasswordHash(hash string) (*User, error)
	// FindByValidationHash(hash string) (*User, error)
	// FindAll() ([]*User, error)
	Update(user *User) error
*/

func (userRepo *userRepository) FindByUsername(username string) (*domain.User,
	error) {

	log.Printf("username %s ", username)

	db := userRepo.db

	var user domain.User
	// create a new query
	db.Select("username").
		From("tbl_user").
		// Where(dbx.Like("name", "Charles")).
		Where(dbx.HashExp{"username": username}).
		One(&user)

	fmt.Print(user)

	return &user, nil
}

func (userRepo *userRepository) Update(user *domain.User) error {
	return nil
}

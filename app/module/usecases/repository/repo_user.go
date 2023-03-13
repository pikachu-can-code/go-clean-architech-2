package repository

import (
	"context"

	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/app/common"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/app/components"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/app/module/entities"
	"gorm.io/gorm"
)

type userRepo struct {
	db     *gorm.DB
	appCtx components.AppContext
}

func NewUserRepo(db *gorm.DB, appCtx components.AppContext) *userRepo {
	return &userRepo{db: db, appCtx: appCtx}
}

func (repo *userRepo) Create(
	ctx context.Context,
	user *entities.UserCreate,
	userRole *entities.UserRoleCreate,
) (*entities.UserCreate, error) {
	var (
		db = repo.db
		userDB = db.Table(user.TableName())
		roleDB = db.Table(userRole.TableName())
	)
	if err := userDB.Create(user).Error; err != nil {
		return nil, common.ErrCannotCreateEntity("User", err)
	}
	userRole.UserId = user.ID
	if err := roleDB.Create(userRole).Error; err != nil {
		return nil, common.ErrCannotCreateEntity("UserRole", err)
	}
	roleFakeId := common.NewUID(userRole.RoleId, common.DbTypeRole, 1)
	user.Role = []string{roleFakeId.String()}
	return user, nil
}

func (repo *userRepo) FindUser(
	ctx context.Context,
	conditions map[string]interface{},
	moreInfo ...string,
) (*entities.User, error) {
	db := repo.db.Table(entities.User{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user entities.User

	if err := db.Where(conditions).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrCannotGetEntity(user.TableName(), err)
	}

	return &user, nil
}

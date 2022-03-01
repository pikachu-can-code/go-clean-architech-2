package repository

import (
	"context"

	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/common"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/components"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/internal/entities"
	"gorm.io/gorm"
)

type userRepo struct {
	db     *gorm.DB
	appCtx components.AppContext
}

func NewUserRepo(db *gorm.DB, appCtx components.AppContext) *userRepo {
	return &userRepo{db: db, appCtx: appCtx}
}

func (repo *userRepo) Create(ctx context.Context, acc *entities.User) error {
	if err := repo.db.Table(acc.TableName()).Create(acc).Error; err != nil {
		return common.ErrCannotCreateEntity(acc.TableName(), err)
	}
	return nil
}

func (repo *userRepo) FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*entities.User, error) {
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

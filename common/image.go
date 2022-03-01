package common

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Image struct {
	SQLModel  `json:",inline"`
	Url       string `json:"url" gorm:"column:url;"`
	Width     int    `json:"width" gorm:"column:width;"`
	Height    int    `json:"height" gorm:"column:height;"`
	Provider  string `json:"provider,omitempty" gorm:"-"`
	Extension string `json:"extension,omitempty" gorm:"-"`
}

func (Image) TableName() string {
	return "images"
}

func (i *Image) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var img Image
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*i = img
	return nil
}

func (i *Image) Value() (driver.Value, error) {
	if i == nil {
		return nil, nil
	}

	return json.Marshal(i)
}

type Images []Image

func (i *Images) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value: ", value))
	}

	var imgs []Image
	if err := json.Unmarshal(bytes, &imgs); err != nil {
		return err
	}

	*i = imgs
	return nil
}

func (i *Images) Value() (driver.Value, error) {
	if i == nil {
		return nil, nil
	}

	return json.Marshal(i)
}

type ImageStore interface {
	FindImageByCondition(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*Image, error)
}

type ImagesStore interface {
	ListImages(ctx context.Context, ids []uint64, moreKeys ...string) ([]Image, error)
}

func (i *Image) Validate(ctx context.Context, imageStore ImageStore) error {
	if i != nil {
		image, _ := imageStore.FindImageByCondition(ctx, map[string]interface{}{"id": i.ID})
		if image == nil {
			return ErrImgNotExisted
		}
		return i.ValidateData(image)
	}

	return nil
}

func (i *Image) ValidateData(j *Image) error {
	if j.Url != i.Url {
		return ErrImgNotExisted
	}

	return nil
}

func (i *Images) Validate(ctx context.Context, imagesStore ImagesStore) error {
	images := []Image(*i)
	ids := make([]uint64, len(images))

	for index := range images {
		ids[index] = images[index].ID
	}

	if i != nil {
		listImage, _ := imagesStore.ListImages(ctx, ids)
		if len(listImage) != len(ids) {
			return ErrInvalidRequest(errors.New("image list is not enough"))
		}

		for index := range listImage {
			if err := listImage[index].ValidateData(&images[index]); err != nil {
				return err
			}
		}
	}

	return nil
}

var (
	ErrImgNotExisted = NewCustomError(
		errors.New("image is not existed"),
		"image is not existed",
		"ErrImgNotExisted",
	)
)

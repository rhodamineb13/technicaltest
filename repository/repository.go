package repository

import (
	"context"
	"strconv"
	awsService "technicaltest/aws"
	"technicaltest/config"
	"technicaltest/database/entities"
	redisService "technicaltest/redis"
	"time"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"gorm.io/gorm"
)

type IRepository interface {
	Create(context.Context, *entities.MyClient) error
	Select(context.Context) ([]entities.MyClient, error)
	Get(context.Context, int) (*entities.MyClient, error)
	Update(context.Context, *entities.MyClient) error
	Delete(context.Context, int) error
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) IRepository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, mc *entities.MyClient) error {
	uploader := s3manager.NewUploader(awsService.AWSSession)
	uploader.Upload(&s3manager.UploadInput{
		Bucket: &config.Conf.AWSBucketName,
		Key:    &mc.ClientLogo,
	})

	id := strconv.Itoa(mc.Id)

	redisService.RC.Set(ctx, id, mc, time.Hour*24)
	return r.db.WithContext(ctx).Create(&mc).Error
}

func (r *Repository) Select(ctx context.Context) ([]entities.MyClient, error) {
	var clients []entities.MyClient

	if err := r.db.Where("deleted_at IS NULL").Find(&clients).Error; err != nil {
		return nil, err
	}

	return clients, nil
}

func (r *Repository) Get(ctx context.Context, id int) (*entities.MyClient, error) {
	client := &entities.MyClient{}

	if err := r.db.Where("id = ?", id).First(&client).Error; err != nil {
		return nil, err
	}

	return client, nil
}

func (r *Repository) Update(ctx context.Context, mc *entities.MyClient) error {
	if err := r.db.Save(&mc).Error; err != nil {
		return err
	}

	id := strconv.Itoa(mc.Id)

	redisService.RC.Set(ctx, id, mc, time.Hour*24)

	return nil
}

func (r *Repository) Delete(ctx context.Context, id int) error {
	var mc entities.MyClient
	if err := r.db.Where("id = ?", id).Delete(&mc).Error; err != nil {
		return err
	}

	idStr := strconv.Itoa(id)

	redisService.RC.Del(ctx, idStr)

	return nil
}

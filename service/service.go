package service

import (
	"context"
	"technicaltest/database/entities"
	"technicaltest/model"
	"technicaltest/repository"
)

type IService interface {
	Create(context.Context, model.CreateRequest) error
	Select(context.Context) ([]model.QueryResponse, error)
	Get(context.Context, int) (*model.QueryResponse, error)
	Update(context.Context, int, model.UpdateRequest) error
	Delete(context.Context, int) error
}

type Service struct {
	repository repository.IRepository
}

func NewService(r repository.IRepository) IService {
	return &Service{
		repository: r,
	}
}

func (s *Service) Create(ctx context.Context, req model.CreateRequest) error {
	mcEntity := &entities.MyClient{
		Name:         req.Name,
		Slug:         req.Slug,
		SelfCapture:  req.SelfCapture,
		IsProject:    req.IsProject,
		ClientPrefix: req.ClientPrefix,
		ClientLogo:   req.ClientLogo,
		Address:      req.Address,
		PhoneNumber:  req.PhoneNumber,
		City:         req.City,
	}

	if err := s.repository.Create(ctx, mcEntity); err != nil {
		return err
	}

	return nil
}
func (s *Service) Select(ctx context.Context) ([]model.QueryResponse, error) {
	res := make([]model.QueryResponse, 0)

	resEntity, err := s.repository.Select(ctx)
	if err != nil {
		return nil, err
	}

	for _, r := range resEntity {
		res = append(res, model.QueryResponse{
			Id:           r.Id,
			Name:         r.Name,
			IsProject:    r.IsProject,
			Slug:         r.Slug,
			SelfCapture:  r.SelfCapture,
			ClientPrefix: r.ClientPrefix,
			ClientLogo:   r.ClientLogo,
			Address:      r.Address,
			PhoneNumber:  r.PhoneNumber,
			City:         r.City,
			CreatedAt:    r.CreatedAt.String(),
			UpdatedAt:    r.UpdatedAt.String(),
			DeletedAt:    r.DeletedAt.String(),
		})
	}

	return res, nil
}
func (s *Service) Get(ctx context.Context, id int) (*model.QueryResponse, error) {
	res, err := s.repository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &model.QueryResponse{
		Id:           res.Id,
		Name:         res.Name,
		Slug:         res.Slug,
		IsProject:    res.IsProject,
		SelfCapture:  res.SelfCapture,
		ClientPrefix: res.ClientPrefix,
		ClientLogo:   res.ClientLogo,
		PhoneNumber:  res.PhoneNumber,
		Address:      res.Address,
		City:         res.City,
		CreatedAt:    res.CreatedAt.String(),
		UpdatedAt:    res.UpdatedAt.String(),
		DeletedAt:    res.DeletedAt.String(),
	}, nil
}
func (s *Service) Update(ctx context.Context, id int, req model.UpdateRequest) error {
	mcEntity := &entities.MyClient{
		Id:           id,
		Name:         req.Name,
		Slug:         req.Slug,
		SelfCapture:  req.SelfCapture,
		IsProject:    req.IsProject,
		ClientPrefix: req.ClientPrefix,
		ClientLogo:   req.ClientLogo,
		Address:      req.Address,
		PhoneNumber:  req.PhoneNumber,
		City:         req.City,
	}

	if err := s.repository.Update(ctx, mcEntity); err != nil {
		return err
	}

	return nil
}
func (s *Service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}

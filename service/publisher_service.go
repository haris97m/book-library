package service

import (
	"github.com/haris97m/go-fiber/helper"
	"github.com/haris97m/go-fiber/model/domain"
	"github.com/haris97m/go-fiber/model/web"
	"github.com/haris97m/go-fiber/repository"
)

// CONTRACT
type PublisherService interface {
	FindAll() (responses []web.PublisherResponse, err error)
	FindById(publisherId int) (response web.PublisherResponse, err error)
	Insert(request web.CreatePublisherRequest) (response web.PublisherResponse, err error)
	Save(request web.UpdatePublisherRequest) (response web.PublisherResponse, err error)
	Delete(publisherId int) error
}

type PublisherServiceImpl struct {
	PublisherRepository repository.PublisherRepository
}

// CONSTRUCTOR
func NewPublisherServiceImpl(publisherRepository repository.PublisherRepository) *PublisherServiceImpl {
	return &PublisherServiceImpl{
		PublisherRepository: publisherRepository,
	}
}

// GET ALL
func (service *PublisherServiceImpl) FindAll() (responses []web.PublisherResponse, err error) {
	publishers, err := service.PublisherRepository.FindAll() //call query find all

	responses = helper.ToPublisherResponses(publishers) //generate response

	return responses, err
}

// GET BY ID
func (service *PublisherServiceImpl) FindById(publisherId int) (response web.PublisherResponse, err error) {
	publisher, err := service.PublisherRepository.FindById(publisherId) //call query find by id

	response = helper.ToPublisherResponse(publisher) //generate response

	return response, err
}

// CREATE
func (service *PublisherServiceImpl) Insert(request web.CreatePublisherRequest) (response web.PublisherResponse, err error) {
	// assign data from request body
	publisher := domain.Publisher{
		Name:     request.Name,
		Industry: request.Industry,
	}

	newPublisher, err := service.PublisherRepository.Insert(publisher) //call query create new

	response = helper.ToPublisherResponse(newPublisher) //generate response

	return response, err
}

// UPDATE
func (service *PublisherServiceImpl) Save(request web.UpdatePublisherRequest) (response web.PublisherResponse, err error) {
	publisher, err := service.PublisherRepository.FindById(request.ID) //call query find by id
	if err != nil {
		return response, err
	}

	// assign data from request body
	publisher.Name = request.Name
	publisher.Industry = request.Industry

	newPublisher, err := service.PublisherRepository.Save(publisher) //call query update data

	response = helper.ToPublisherResponse(newPublisher) //generate response

	return response, err
}

// DELETE
func (service *PublisherServiceImpl) Delete(publisherId int) error {
	publisher, err := service.PublisherRepository.FindById(publisherId) //call query find by id

	// call query delete
	if err := service.PublisherRepository.Delete(publisher); err != nil {
		return err
	}

	return err
}

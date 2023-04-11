package service

import (
	"github.com/haris97m/go-fiber/helper"
	"github.com/haris97m/go-fiber/model/domain"
	"github.com/haris97m/go-fiber/model/web"
	"github.com/haris97m/go-fiber/repository"
)

type AuthorService interface {
	FindAll() (responses []web.AuthorResponse, err error)
	FindById(authorId int) (response web.AuthorResponse, err error)
	Insert(request web.AuthorCreateRequest) (response web.AuthorResponse, err error)
	Save(request web.AuthorUpdateRequest) (response web.AuthorResponse, err error)
	Delete(authorId int) error
}

type AuthorServiceImpl struct {
	Repository repository.AuthorRepository
}

func NewAuthorServiceImpl(authorRepository repository.AuthorRepository) *AuthorServiceImpl {
	return &AuthorServiceImpl{
		Repository: authorRepository,
	}
}

func (service *AuthorServiceImpl) FindAll() (responses []web.AuthorResponse, err error) {
	authors, err := service.Repository.FindAll()
	responses = helper.ToAuthorResponses(authors)

	return responses, err
}

func (service *AuthorServiceImpl) FindById(id int) (response web.AuthorResponse, err error) {
	author, err := service.Repository.FindById(id)
	response = helper.ToAuthorResponse(author)

	return response, err
}

func (service *AuthorServiceImpl) Insert(request web.AuthorCreateRequest) (response web.AuthorResponse, err error) {

	author := domain.Author{
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		DateOfBirth: request.DateOfBirth,
		DateOfDeath: request.DateOfDeath,
	}

	newAuthor, err := service.Repository.Insert(author)
	response = helper.ToAuthorResponse(newAuthor)

	return response, err
}

func (service *AuthorServiceImpl) Save(request web.AuthorUpdateRequest) (response web.AuthorResponse, err error) {
	author, err := service.Repository.FindById(request.ID)
	if err != nil {
		return response, err
	}

	author.FirstName = request.FirstName
	author.LastName = request.LastName
	author.DateOfBirth = request.DateOfBirth
	author.DateOfDeath = request.DateOfDeath

	newAuthor, err := service.Repository.Save(author)
	response = helper.ToAuthorResponse(newAuthor)

	return response, err
}

func (service *AuthorServiceImpl) Delete(id int) error {
	author, err := service.Repository.FindById(id)
	if err := service.Repository.Delete(author); err != nil {
		return err
	}

	return err
}

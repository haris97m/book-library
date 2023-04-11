package repository

import (
	"github.com/haris97m/go-fiber/model/domain"
	"gorm.io/gorm"
)

type AuthorRepository interface {
	FindAll() (authors []domain.Author, err error)
	FindById(authorId int) (author domain.Author, err error)
	Insert(author domain.Author) (domain.Author, error)
	Save(author domain.Author) (domain.Author, error)
	Delete(author domain.Author) error
}

type AuthorRepositoryImpl struct {
	DB *gorm.DB
}

func NewAuthorRepositoryImpl(db *gorm.DB) *AuthorRepositoryImpl {
	return &AuthorRepositoryImpl{DB: db}
}

func (repository *AuthorRepositoryImpl) FindAll() (authors []domain.Author, err error) {
	err = repository.DB.Find(&authors).Error
	return authors, err
}

func (repository *AuthorRepositoryImpl) FindById(id int) (author domain.Author, er error) {
	err := repository.DB.First(&author, id).Error
	return author, err
}

func (repository *AuthorRepositoryImpl) Insert(author domain.Author) (domain.Author, error) {
	err := repository.DB.Create(&author).Error
	return author, err
}

func (repository *AuthorRepositoryImpl) Save(author domain.Author) (domain.Author, error) {
	err := repository.DB.Save(&author).Error
	return author, err
}

func (repository *AuthorRepositoryImpl) Delete(author domain.Author) error {
	return repository.DB.Delete(&author).Error
}

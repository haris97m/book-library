package repository

import (
	"github.com/haris97m/go-fiber/model/domain"
	"gorm.io/gorm"
)

// CONTRACT
type PublisherRepository interface {
	FindAll() (publishers []domain.Publisher, err error)
	FindById(publisherId int) (publisher domain.Publisher, err error)
	Insert(publisher domain.Publisher) (domain.Publisher, error)
	Save(publisher domain.Publisher) (domain.Publisher, error)
	Delete(publisher domain.Publisher) error
}

type PublisherRepositoryImpl struct {
	DB *gorm.DB
}

// CONSTRUCTOR
func NewPublisherRepositoryImpl(db *gorm.DB) *PublisherRepositoryImpl {
	return &PublisherRepositoryImpl{DB: db}
}

// GET ALL
func (repository PublisherRepositoryImpl) FindAll() (publishers []domain.Publisher, err error) {
	err = repository.DB.Find(&publishers).Error //query find all
	return publishers, err
}

// GET BY ID
func (repository PublisherRepositoryImpl) FindById(id int) (publisher domain.Publisher, err error) {
	err = repository.DB.First(&publisher, id).Error //query find by id
	return publisher, err
}

// CREATE
func (repository PublisherRepositoryImpl) Insert(publisher domain.Publisher) (domain.Publisher, error) {
	err := repository.DB.Create(&publisher).Error //query create new
	return publisher, err
}

// UPDATE
func (repository PublisherRepositoryImpl) Save(publisher domain.Publisher) (domain.Publisher, error) {
	err := repository.DB.Save(&publisher).Error //query update data
	return publisher, err
}

// DELETE
func (repository PublisherRepositoryImpl) Delete(publisher domain.Publisher) error {
	return repository.DB.Delete(&publisher).Error //query delete data
}

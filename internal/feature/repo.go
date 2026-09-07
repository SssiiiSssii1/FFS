package feature

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(feature *Feature) error {
	return r.db.Create(feature).Error
}

func (r *Repository) FindAll(features *[]Feature) error {
	return r.db.Find(features).Error
}

func (r *Repository) FindByID(id uint, feature *Feature) error {
	return r.db.First(feature, id).Error
}

func (r *Repository) Update(feature *Feature) error {
	return r.db.Save(feature).Error
}

func (r *Repository) Delete(feature *Feature) error {
	return r.db.Delete(feature).Error
}

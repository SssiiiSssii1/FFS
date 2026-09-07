package feature

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(feature *Feature) error {
	return s.repo.Create(feature)
}

func (s *Service) FindAll(features *[]Feature) error {
	return s.repo.FindAll(features)
}

func (s *Service) FindByID(id uint, feature *Feature) error {
	return s.repo.FindByID(id, feature)
}

func (s *Service) Update(feature *Feature) error {
	return s.repo.Update(feature)
}

func (s *Service) Delete(feature *Feature) error {
	return s.repo.Delete(feature)
}

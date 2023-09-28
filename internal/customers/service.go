package customers

import (
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
)

type Service interface {
	Create(customers *domain.Customers) error
	ReadAll() ([]*domain.Customers, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Create(customers *domain.Customers) error {
	_, err := s.r.Create(customers)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ReadAll() ([]*domain.Customers, error) {
	return s.r.ReadAll()
}

func (s *service) InsertAll() ([]*domain.Customers, error) {

	return s.r.ReadAll()
}

/*func (s *service) LoadAllCustomers() ([]*domain.Customers, error) {
	var products []domain.Customers
	file, err := os.ReadFile(s.pathToFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(file), &products)
	if err != nil {
		return nil, err
	}
	return products, nil
	return s.r.ReadAll()
}

// read json file and insert into database
func ReadFile(path string) ([]any, error) {

}*/

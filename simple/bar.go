package simple

type BarRepository struct {
}

// Constructor for BarRepository
func NewBarRepository() *BarRepository {
	o := new(BarRepository)
	return o
}

type BarService struct {
	*BarRepository
}

// Constructor for BarService
func NewBarService(barRepository *BarRepository) *BarService {
	return &BarService{
		BarRepository: barRepository,
	}
}

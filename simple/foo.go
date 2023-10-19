package simple

type FooRepository struct {
}

// Constructor for FooRepository
func NewFooRepository() *FooRepository {
	o := new(FooRepository)
	return o
}

type FooService struct {
	*FooRepository
}

// Constructor for FooService
func NewFooService(fooRepository *FooRepository) *FooService {
	return &FooService{
		FooRepository: fooRepository,
	}
}

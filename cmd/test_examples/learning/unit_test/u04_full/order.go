package u04_full

type Order struct {
	Amount float64
}

type Repository interface {
	Save(order *Order) error
}

type PaymentGateway interface {
	Charge(amount float64) error
}

type Service struct {
	repo    Repository
	payment PaymentGateway
}

func NewService(repo Repository, payment PaymentGateway) *Service {
	return &Service{repo: repo, payment: payment}
}

func (s *Service) CreateOrder(amount float64) error {
	order := &Order{Amount: amount}
	if err := s.payment.Charge(amount); err != nil {
		return err
	}
	return s.repo.Save(order)
}

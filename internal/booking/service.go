package booking

type Service struct {
	store BookingStore
}

func NewService(store BookingStore) *Service {
	return &Service{store}
}

func (s *Service) Book(book Booking) error {
	return s.store.Book(book)
}
func (s *Service) ListBookings(movieID string) []Booking {
	return s.ListBookings(movieID)
}

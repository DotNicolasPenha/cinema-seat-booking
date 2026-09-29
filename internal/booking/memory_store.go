package booking

import "sync"

type MemoryStore struct {
	bookings map[string]Booking
	mu       sync.RWMutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		bookings: map[string]Booking{},
	}
}
func (s *MemoryStore) Book(b Booking) error {
	defer s.mu.Unlock()
	s.mu.Lock()

	if _, exists := s.bookings[b.SeatID]; exists {
		return ErrSeatAlreadyBooked
	}
	s.bookings[b.SeatID] = b
	return nil
}

func (s *MemoryStore) ListBookings(movieID string) []Booking {
	defer s.mu.RUnlock()
	s.mu.RLock()

	var result []Booking
	for _, b := range s.bookings {
		if movieID == b.MovieID {
			result = append(result, b)
		}
	}
	return result
}

package main

import "sync"

// Store mantiene los datos en memoria.
// Los mapas son privados para demostrar encapsulación.
type Store struct {
	mu         sync.RWMutex
	products   map[int]Product
	customers  map[int]Customer
	orders     map[int]Order
	payments   map[int]Payment
	nextProduct int
	nextCustomer int
	nextOrder   int
	nextPayment int
}

func NewStore() *Store {
	return &Store{
		products:  make(map[int]Product),
		customers: make(map[int]Customer),
		orders:    make(map[int]Order),
		payments:  make(map[int]Payment),
		nextProduct: 1,
		nextCustomer: 1,
		nextOrder: 1,
		nextPayment: 1,
	}
}

func (s *Store) AddProduct(p Product) Product {
	s.mu.Lock()
	defer s.mu.Unlock()

	p.ID = s.nextProduct
	s.nextProduct++
	s.products[p.ID] = p
	return p
}

func (s *Store) GetProducts() []Product {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]Product, 0, len(s.products))
	for _, p := range s.products {
		result = append(result, p)
	}
	return result
}

func (s *Store) GetProduct(id int) (Product, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	p, ok := s.products[id]
	return p, ok
}

func (s *Store) AddCustomer(c Customer) Customer {
	s.mu.Lock()
	defer s.mu.Unlock()

	c.ID = s.nextCustomer
	s.nextCustomer++
	s.customers[c.ID] = c
	return c
}

func (s *Store) GetCustomer(id int) (Customer, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	c, ok := s.customers[id]
	return c, ok
}

func (s *Store) AddOrder(o Order) Order {
	s.mu.Lock()
	defer s.mu.Unlock()

	o.ID = s.nextOrder
	s.nextOrder++
	s.orders[o.ID] = o
	return o
}

func (s *Store) GetOrder(id int) (Order, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	o, ok := s.orders[id]
	return o, ok
}

func (s *Store) AddPayment(p Payment) Payment {
	s.mu.Lock()
	defer s.mu.Unlock()

	p.ID = s.nextPayment
	s.nextPayment++
	s.payments[p.ID] = p
	return p
}

func (s *Store) GetPayment(id int) (Payment, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	p, ok := s.payments[id]
	return p, ok
}

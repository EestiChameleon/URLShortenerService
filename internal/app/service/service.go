package service

import "fmt"

// Service отвечает за обработку данных в app.
type Service struct {
}

func (s *Service) FooBar() {
	fmt.Println("Готовим пиццу.")
}

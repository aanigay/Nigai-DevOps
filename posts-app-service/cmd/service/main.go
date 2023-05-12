package main

import "posts-app-service/internal/app/service"

func main() {
	s := service.NewService()
	err := s.Run()
	if err != nil {
		panic(err)
	}
}

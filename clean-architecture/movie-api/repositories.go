package main

import "fmt"

type IMovieRepository interface {
	UpdateMovie(id int, movie movieModel) error
	GetAllMovies() []movieModel
}

type inmemoryMovieRepository struct {
	Movies []movieModel
}

func NewInmemoyMovieRepository() IMovieRepository {
	var movies = []movieModel{
		{ID: 1, Title: "The Shawshank Redemption", ReleaseYear: 1994, Score: 9.3},
		{ID: 2, Title: "The Godfather", ReleaseYear: 1972, Score: 9.2},
		{ID: 3, Title: "The Dark Knight", ReleaseYear: 2008, Score: 9.0},
	}

	return &inmemoryMovieRepository{
		Movies: movies,
	}
}

func (immRepo *inmemoryMovieRepository) UpdateMovie(id int, updateMovie movieModel) error {
	for i, movie := range immRepo.Movies {
		if movie.ID == id {
			immRepo.Movies[i].ReleaseYear = updateMovie.ReleaseYear
			immRepo.Movies[i].Score = updateMovie.Score
			immRepo.Movies[i].Title = updateMovie.Title
			return nil
		}
	}
	return fmt.Errorf("ID not found")
}

func (imRepo *inmemoryMovieRepository) GetAllMovies() []movieModel {
	return imRepo.Movies
}

package main

type IMovieService interface {
	UpdateMovie(id int, movie movieModel) error
	GetMovies() []movieModel
}

type DefaultMovieService struct {
	movieRepo IMovieRepository
}

func NewDefaultMovieSerives(mRepo IMovieRepository) IMovieService {
	return &DefaultMovieService{
		movieRepo: mRepo,
	}
}

func (dfmService *DefaultMovieService) UpdateMovie(id int, updateMovie movieModel) error {
	return dfmService.movieRepo.UpdateMovie(id, updateMovie)
}

func (dfmService *DefaultMovieService) GetMovies() []movieModel {
	return dfmService.movieRepo.GetAllMovies()
}

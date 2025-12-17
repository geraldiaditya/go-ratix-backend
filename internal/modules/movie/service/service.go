package service

import (
	"github.com/geraldiaditya/ratix-backend/internal/modules/movie/domain"
	"github.com/geraldiaditya/ratix-backend/internal/modules/movie/dto"
)

type MovieService struct {
	Repo domain.MovieRepository
}

func NewMovieService(repo domain.MovieRepository) *MovieService {
	return &MovieService{Repo: repo}
}

func (s *MovieService) GetCategories() ([]dto.GenreResponse, error) {
	genres, err := s.Repo.GetAllGenres()
	if err != nil {
		return nil, err
	}
	var resp []dto.GenreResponse
	for _, g := range genres {
		resp = append(resp, dto.GenreResponse{ID: g.ID, Name: g.Name})
	}
	return resp, nil
}

func (s *MovieService) GetBanner() (*dto.BannerResponse, error) {
	// Logic: Get 'now_showing' AND standard picking logic (e.g. highest rated or first)
	// For banner, we might just want 1, so limit=1, offset=0
	movies, _, err := s.Repo.GetByStatus("now_showing", 1, 0)
	if err != nil {
		return nil, err
	}
	if len(movies) == 0 {
		return nil, nil // Or default banner
	}
	m := movies[0]

	genres := make([]string, len(m.Genres))
	for i, g := range m.Genres {
		genres[i] = g.Name
	}

	return &dto.BannerResponse{
		MovieID:   m.ID,
		Title:     m.Title,
		PosterURL: m.PosterURL,
		Rating:    m.Rating,
		Genres:    genres,
	}, nil
}

func (s *MovieService) GetMovies(category string, page, limit int) (*dto.MovieListResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	var movies []domain.Movie
	var total int64
	var err error

	if category == "" || category == "now_showing" || category == "coming_soon" {
		status := category
		if status == "" {
			status = "now_showing" // Default
		}
		movies, total, err = s.Repo.GetByStatus(status, limit, offset)
	} else {
		// Assume it's a genre
		movies, total, err = s.Repo.GetByGenre(category, limit, offset)
	}

	if err != nil {
		return nil, err
	}

	resp := make([]dto.MovieResponse, len(movies))
	for i, m := range movies {
		resp[i] = dto.ToMovieResponse(m)
	}

	totalPages := int(total) / limit
	if int(total)%limit != 0 {
		totalPages++
	}

	return &dto.MovieListResponse{
		Movies: resp,
		Meta: dto.PaginationMeta{
			CurrentPage: page,
			TotalPages:  totalPages,
			TotalItems:  total,
			Limit:       limit,
		},
	}, nil
}

func (s *MovieService) GetDetail(id int64) (*dto.MovieDetailResponse, error) {
	movie, err := s.Repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return dto.ToMovieDetailResponse(movie), nil
}

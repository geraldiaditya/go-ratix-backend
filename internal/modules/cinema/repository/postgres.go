package repository

import (
	"github.com/geraldiaditya/ratix-backend/internal/modules/cinema/domain"
	"gorm.io/gorm"
)

type PostgresCinemaRepository struct {
	DB *gorm.DB
}

func NewPostgresCinemaRepository(db *gorm.DB) *PostgresCinemaRepository {
	return &PostgresCinemaRepository{DB: db}
}

func (r *PostgresCinemaRepository) GetAllCities() ([]string, error) {
	var cities []string
	if err := r.DB.Model(&domain.Cinema{}).Distinct("city").Pluck("city", &cities).Error; err != nil {
		return nil, err
	}
	return cities, nil
}

func (r *PostgresCinemaRepository) GetAllBrands() ([]string, error) {
	var brands []string
	if err := r.DB.Model(&domain.Cinema{}).
		Distinct("brand").
		Where("brand IS NOT NULL AND brand != ''").
		Pluck("brand", &brands).Error; err != nil {
		return nil, err
	}
	return brands, nil
}

func (r *PostgresCinemaRepository) GetCinemas(filter domain.CinemaFilter) ([]domain.Cinema, error) {
	cinemas := []domain.Cinema{}
	query := r.DB.Model(&domain.Cinema{})

	if filter.Name != "" {
		query = query.Where("name ILIKE ?", "%"+filter.Name+"%")
	}
	if filter.City != "" {
		query = query.Where("city = ?", filter.City)
	}
	if filter.Brand != "" {
		query = query.Where("brand = ?", filter.Brand)
	}

	if filter.Lat != 0 && filter.Lon != 0 {
		// Haversine formula
		haversine := `(
			6371 * acos(
				cos(radians(?)) * cos(radians(lat)) * cos(radians(lon) - radians(?)) +
				sin(radians(?)) * sin(radians(lat))
			)
		)`
		query = query.Select("*, "+haversine+" AS hav_dist", filter.Lat, filter.Lon, filter.Lat)
		query = query.Order("hav_dist ASC")

		if filter.Radius > 0 {
			query = query.Where(haversine+" < ?", filter.Lat, filter.Lon, filter.Lat, filter.Radius)
		}

		// Use a temporary struct to scan the result including the distance
		var results []struct {
			domain.Cinema
			Distance float64 `gorm:"column:hav_dist"`
		}

		if err := query.Scan(&results).Error; err != nil {
			return nil, err
		}

		// Map back to domain.Cinema
		for _, r := range results {
			c := r.Cinema
			c.Distance = r.Distance
			cinemas = append(cinemas, c)
		}
	} else {
		if err := query.Find(&cinemas).Error; err != nil {
			return nil, err
		}
	}
	return cinemas, nil
}

func (r *PostgresCinemaRepository) GetByID(id int64) (*domain.Cinema, error) {
	var cinema domain.Cinema
	if err := r.DB.First(&cinema, id).Error; err != nil {
		return nil, err
	}
	return &cinema, nil
}

func (r *PostgresCinemaRepository) Create(cinema *domain.Cinema) error {
	return r.DB.Create(cinema).Error
}

func (r *PostgresCinemaRepository) GetCinemaByShowtimeID(showtimeID int64) (*domain.Cinema, error) {
	// Join Showtime and Cinema tables
	// Assuming tables are "showtimes" and "cinemas"
	// and showtimes has cinema_id
	var cinema domain.Cinema

	// We need to access Showtime table which is likely "showtimes".
	// Since we don't import movieDomain here, we can use raw query or map to struct if we had it.
	// But actually, we can just do a join if we know the schema.
	// "SELECT cinemas.* FROM cinemas JOIN showtimes ON showtimes.cinema_id = cinemas.id WHERE showtimes.id = ?"

	if err := r.DB.Joins("JOIN showtimes ON showtimes.cinema_id = cinemas.id").
		Where("showtimes.id = ?", showtimeID).
		First(&cinema).Error; err != nil {
		return nil, err
	}
	return &cinema, nil
}

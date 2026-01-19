package domain

import (
	"time"

	"gorm.io/gorm"
)

type Cinema struct {
	ID         int64          `gorm:"primaryKey" json:"id"`
	Name       string         `gorm:"not null;type:varchar(100)" json:"name"`
	Brand      string         `gorm:"type:varchar(50)" json:"brand"` // e.g. "Cinema XXI", "CGV", "Cinepolis"
	City       string         `gorm:"not null;type:varchar(50)" json:"city"`
	Address    string         `gorm:"type:text" json:"address"`
	BasePrice  float64        `gorm:"not null;type:decimal(10,2);default:50000" json:"base_price"`
	Rating     float64        `gorm:"type:decimal(3,1);default:0" json:"rating"`
	Lat        float64        `gorm:"type:decimal(10,8)" json:"lat"`
	Lon        float64        `gorm:"type:decimal(11,8)" json:"lon"`
	PictureURL string         `gorm:"type:text" json:"picture_url"`
	Distance   float64        `gorm:"-" json:"distance"` // Calculated at runtime, not stored
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Theater struct {
	ID        int64          `gorm:"primaryKey" json:"id"`
	CinemaID  int64          `gorm:"not null" json:"cinema_id"`
	Cinema    Cinema         `gorm:"foreignKey:CinemaID" json:"cinema"`
	Name      string         `gorm:"not null;type:varchar(50)" json:"name"` // e.g. "Studio 1", "IMAX"
	Type      string         `gorm:"type:varchar(20)" json:"type"`          // Regular, IMAX, Premiere
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type CinemaFilter struct {
	City   string
	Brand  string
	Lat    float64
	Lon    float64
	Radius float64
}

type CinemaRepository interface {
	GetAllCities() ([]string, error)
	GetAllBrands() ([]string, error)
	GetCinemas(filter CinemaFilter) ([]Cinema, error)
	GetByID(id int64) (*Cinema, error)
	GetCinemaByShowtimeID(showtimeID int64) (*Cinema, error)
	Create(cinema *Cinema) error
}

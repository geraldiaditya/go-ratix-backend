package repository

import (
	"errors"
	"fmt"

	"github.com/geraldiaditya/ratix-backend/internal/modules/ticket/domain"
	"gorm.io/gorm"
)

type PostgresTicketRepository struct {
	DB *gorm.DB
}

func NewPostgresTicketRepository(db *gorm.DB) *PostgresTicketRepository {
	return &PostgresTicketRepository{DB: db}
}

func (r *PostgresTicketRepository) GetByUserID(userID int64, status string) ([]domain.Ticket, error) {
	var tickets []domain.Ticket
	query := r.DB.Where("user_id = ?", userID).Preload("Movie")

	if status != "" {
		if status == "history" {
			// History means watched or cancelled -> simplified for now, assuming 'history' status exists
			// Or logic: status IN ('completed', 'cancelled')
			query = query.Where("status IN ?", []string{"completed", "cancelled"})
		} else {
			// Active default
			query = query.Where("status = ?", "active")
		}
	}

	// Recent bookings first
	if err := query.Order("created_at desc").Find(&tickets).Error; err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r *PostgresTicketRepository) GetByID(id int64) (*domain.Ticket, error) {
	var ticket domain.Ticket
	if err := r.DB.Preload("Movie").First(&ticket, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("ticket not found")
		}
		return nil, err
	}
	return &ticket, nil
}

func (r *PostgresTicketRepository) Create(ticket *domain.Ticket) error {
	return r.DB.Create(ticket).Error
}

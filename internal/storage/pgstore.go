package storage

import (
	"context"
	"errors"
	"time"

	"github.com/ardaeu/go-contacts-api/config"
	"github.com/ardaeu/go-contacts-api/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrNotFound = errors.New("kayıt bulunamadı")

type PGStore struct {
	db *pgxpool.Pool
}

func NewPGStore() *PGStore {
	return &PGStore{
		db: config.DB,
	}
}

func (p *PGStore) Create(ctx context.Context, c *model.Contact) error {
	c.ID = uuid.New().String()
	now := time.Now()

	_, err := p.db.Exec(ctx,
		`INSERT INTO contacts (id, name, email, phone, created_at, updated_at)
		 VALUES ($1, $2, $3, $4, $5, $6)`,
		c.ID, c.Name, c.Email, c.Phone, now, now,
	)

	return err
}

func (p *PGStore) GetAll(ctx context.Context) ([]model.Contact, error) {
	rows, err := p.db.Query(ctx, `SELECT id, name, email, phone FROM contacts`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []model.Contact

	for rows.Next() {
		var c model.Contact
		if err := rows.Scan(&c.ID, &c.Name, &c.Email, &c.Phone); err != nil {
			return nil, err
		}
		contacts = append(contacts, c)
	}

	return contacts, nil
}

func (p *PGStore) GetByID(ctx context.Context, id string) (*model.Contact, error) {
	row := p.db.QueryRow(ctx, `SELECT id, name, email, phone FROM contacts WHERE id=$1`, id)

	var c model.Contact
	err := row.Scan(&c.ID, &c.Name, &c.Email, &c.Phone)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &c, nil
}

func (p *PGStore) Update(ctx context.Context, c *model.Contact) error {
	now := time.Now()
	cmdTag, err := p.db.Exec(ctx,
		`UPDATE contacts SET name=$1, email=$2, phone=$3, updated_at=$4 WHERE id=$5`,
		c.Name, c.Email, c.Phone, now, c.ID,
	)

	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}

func (p *PGStore) Delete(ctx context.Context, id string) error {
	cmdTag, err := p.db.Exec(ctx, `DELETE FROM contacts WHERE id=$1`, id)
	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}

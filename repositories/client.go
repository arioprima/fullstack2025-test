package repositories

import (
	"database/sql"
	"test-coding/models"
	"time"
)

type ClientRepository struct {
    DB *sql.DB
}

func (r *ClientRepository) Create(client *models.Client) error {
    query := `INSERT INTO my_client (name, slug, is_project, self_capture, client_prefix, client_logo, address, phone_number, city, created_at) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`
    return r.DB.QueryRow(query, client.Name, client.Slug, client.IsProject, client.SelfCapture, client.ClientPrefix, client.ClientLogo, client.Address, client.PhoneNumber, client.City, time.Now()).Scan(&client.ID)
}

func (r *ClientRepository) Update(client *models.Client) error {
    query := `UPDATE my_client SET name=$1, is_project=$2, self_capture=$3, client_prefix=$4, client_logo=$5, address=$6, phone_number=$7, city=$8, updated_at=$9 WHERE slug=$10`
    _, err := r.DB.Exec(query, client.Name, client.IsProject, client.SelfCapture, client.ClientPrefix, client.ClientLogo, client.Address, client.PhoneNumber, client.City, time.Now(), client.Slug)
    return err
}

func (r *ClientRepository) Delete(slug string) error {
    query := `UPDATE my_client SET deleted_at=$1 WHERE slug=$2`
    _, err := r.DB.Exec(query, time.Now(), slug)
    return err
}

func (r *ClientRepository) GetBySlug(slug string) (*models.Client, error) {
    var client models.Client
    query := `SELECT id, name, slug, is_project, self_capture, client_prefix, client_logo, address, phone_number, city, created_at, updated_at, deleted_at 
              FROM my_client WHERE slug=$1`
    err := r.DB.QueryRow(query, slug).Scan(&client.ID, &client.Name, &client.Slug, &client.IsProject, &client.SelfCapture, &client.ClientPrefix, &client.ClientLogo, &client.Address, &client.PhoneNumber, &client.City, &client.CreatedAt, &client.UpdatedAt, &client.DeletedAt)
    if err != nil {
        return nil, err
    }
    return &client, nil
}
package repository

import (
	"context"
	"database/sql"
	"tradingsystembackend/modul/registrationotp/model/domain"
)

type InitRegistartionRepository interface {
	Init(ctx context.Context, tx *sql.Tx, phonenum string, secretquestion string, secretanswer string) int
	FindIsExist(ctx context.Context, tx *sql.Tx, phonenum string, secretquestion string, secretanswer string) domain.InitRegistration
	UpdateSecretAnswer(ctx context.Context, tx *sql.Tx, category domain.InitRegistration) domain.InitRegistration
	//	Save(ctx context.Context, tx *sql.Tx, category domain.InitRegistration) domain.InitRegistration
	// Delete(ctx context.Context, tx *sql.Tx, category domain.InitRegistration)
	// FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.InitRegistration, error)
	// FindAll(ctx context.Context, tx *sql.Tx) []domain.InitRegistration
}

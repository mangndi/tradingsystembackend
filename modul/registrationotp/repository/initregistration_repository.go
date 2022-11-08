package repository

import (
	"context"
	"database/sql"
	"tradingsystembackend/modul/registrationotp/model/domain"
)

type InitRegistartionRepository interface {
	ValidatePhoneNum(ctx context.Context, tx *sql.Tx, phonenum string) string
	// Jika hasil validasi ada, tanya secret Question
	GetDataIfValid(ctx context.Context, tx *sql.Tx, phonenum string, secretanswer string) (domain.Registration, error)
	// jika tidak ada
	Register(ctx context.Context, tx *sql.Tx, phonenum string, secretquestion string, secretanswer string) int

	FindIsExist(ctx context.Context, tx *sql.Tx, phonenum string) (domain.InitRegistration, error)

	UpdateSecretAnswer(ctx context.Context, tx *sql.Tx, category domain.InitRegistration) domain.InitRegistration
	//	Save(ctx context.Context, tx *sql.Tx, category domain.InitRegistration) domain.InitRegistration
	// Delete(ctx context.Context, tx *sql.Tx, category domain.InitRegistration)
	// FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.InitRegistration, error)
	// FindAll(ctx context.Context, tx *sql.Tx) []domain.InitRegistration
}

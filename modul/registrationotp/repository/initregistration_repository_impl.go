package repository

import (
	"context"
	"database/sql"
	"tradingsystembackend/helper"
	"tradingsystembackend/modul/registrationotp/model/domain"
)

type InitRegistrationRepositoryImpl struct {
}

func NewInitRegistrationRepository() InitRegistartionRepository {
	return &InitRegistrationRepositoryImpl{}
}

func (initRegRepo *InitRegistrationRepositoryImpl) Init(ctx context.Context, tx *sql.Tx, phonenum string, secretquestion string, secretanswer string) int {
	//TODO implement me
	SQL := "insert into initregistration(phonenum, secretquestion, secretanswer) values (?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, phonenum, secretquestion, secretanswer)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	Id = int(id)
	return Id

	//panic("implement me")
}

func (initRegRepo *InitRegistrationRepositoryImpl) FindIsExist(ctx context.Context, tx *sql.Tx, phonenum string, secretquestion string, secretanswer string) domain.InitRegistration {
	//TODO implement me
	panic("implement me")
}

func (initRegRepo *InitRegistrationRepositoryImpl) UpdateSecretAnswer(ctx context.Context, tx *sql.Tx, category domain.InitRegistration) domain.InitRegistration {
	//TODO implement me
	panic("implement me")
}

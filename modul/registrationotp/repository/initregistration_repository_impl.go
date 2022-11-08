package repository

import (
	"context"
	"database/sql"
	"errors"
	"tradingsystembackend/helper"
	"tradingsystembackend/modul/registrationotp/model/domain"
)

type InitRegistrationRepositoryImpl struct {
}

func NewInitRegistrationRepository() InitRegistartionRepository {
	return &InitRegistrationRepositoryImpl{}
}

func (i *InitRegistrationRepositoryImpl) ValidatePhoneNum(ctx context.Context, tx *sql.Tx, phonenum string) string {
	//Apakah nomor HP sudah pernah terdaftar
	SQL := "select secretquestion from initregistration where phonenum = ? LIMIT 1"

	rows, err := tx.QueryContext(ctx, SQL, phonenum)
	helper.PanicIfError(err)
	defer rows.Close()

	initRegistration := domain.InitRegistration{}
	if rows.Next() {
		err := rows.Scan(&initRegistration.SecretQuestion)
		helper.PanicIfError(err)
		return initRegistration.SecretQuestion
	} else {
		return ""
	}
}

func (i *InitRegistrationRepositoryImpl) GetDataIfValid(ctx context.Context, tx *sql.Tx, phonenum string, secretanswer string) (domain.Registration, error) {
	//Jika sudah terdaftar  apakah nasabah ini nasabah yg valid dengan menanyakan secret question
	SQL := "select initregid from initregistration where phonenum = ? and secretanswer = ? LIMIT 1"

	rows, err := tx.QueryContext(ctx, SQL, phonenum, secretanswer)
	helper.PanicIfError(err)
	defer rows.Close()

	id := int(-1)

	initRegistration := domain.InitRegistration{}
	if rows.Next() {
		err := rows.Scan(&initRegistration.InitRegId)
		id = int(initRegistration.InitRegId)
		helper.PanicIfError(err)

		//		return initRegistration, nil
	}

	registration := domain.Registration{}
	if id != -1 {
		SQL := "select regid,phonenum,email,idcardnumber,idcardimage,idcardname,idcardaddr1,idcardaddr2,idcardsex,idcardcity,idcardprovince,idcardzip,  idcardcountry,  nationality,    job,    sourceofincome,  taxcardimage,taxcardnumber,photoselfieimage,mothermaindenname,parentname,parentjob,registdate,pagenum  from initregistration where regid = ? and phonenum = ? LIMIT 1"

		rows, err := tx.QueryContext(ctx, SQL, id, phonenum)
		helper.PanicIfError(err)
		defer rows.Close()

		if rows.Next() {
			err := rows.Scan(&registration.Regid, &registration.Phonenum, &registration.Email, &registration.Idcardnumber, &registration.Idcardimage,
				&registration.Idcardname, &registration.Idcardaddr1, &registration.Idcardaddr2, &registration.Idcardsex, &registration.Idcardcity, &registration.Idcardprovince,
				&registration.Idcardzip, &registration.Idcardcountry, &registration.Nationality, &registration.Job, &registration.Sourceofincome, &registration.Taxcardimage,
				&registration.Taxcardnumber, &registration.Photoselfieimage, &registration.Mothermaindenname, &registration.Parentname, &registration.Parentjob, &registration.Registdate, &registration.Pagenum)

			helper.PanicIfError(err)

			return registration, nil
		} else {
			return registration, errors.New("Data is not found")
		}

	} else {
		return registration, errors.New("Data is not found")
	}
}

func (i *InitRegistrationRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, phonenum string, secretquestion string, secretanswer string) int {
	//TODO implement me
	panic("implement me")
}

func (i InitRegistrationRepositoryImpl) FindIsExist(ctx context.Context, tx *sql.Tx, phonenum string) (domain.InitRegistration, error) {
	//TODO implement me
	panic("implement me")
}

func (i *InitRegistrationRepositoryImpl) UpdateSecretAnswer(ctx context.Context, tx *sql.Tx, category domain.InitRegistration) domain.InitRegistration {
	//TODO implement me
	panic("implement me")
}

/*
func (initRegRepo *InitRegistrationRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, phonenum string, secretquestion string, secretanswer string) int {
	//TODO implement me
	SQL := "insert into initregistration(phonenum, secretquestion, secretanswer) values (?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, phonenum, secretquestion, secretanswer)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	resid := int(id)
	return resid

	//panic("implement me")
}

func (initRegRepo *InitRegistrationRepositoryImpl) FindIsExist(ctx context.Context, tx *sql.Tx, phonenum string) (domain.InitRegistration, error) {
	//TODO implement me
	SQL := "select phonenum from initregistration where phonenum = ?"

	rows, err := tx.QueryContext(ctx, SQL, phonenum)
	helper.PanicIfError(err)
	defer rows.Close()

	initRegistration := domain.InitRegistration{}
	if rows.Next() {
		err := rows.Scan(&initRegistration.PhoneNum, &initRegistration.SecretAnswer)
		helper.PanicIfError(err)
		return initRegistration, nil
	} else {
		return initRegistration, errors.New("Phone number is not found")
	}

}

func (initRegRepo *InitRegistrationRepositoryImpl) FindIsExist(ctx context.Context, tx *sql.Tx, phonenum string) (domain.InitRegistration, error) {
	//TODO implement me
	SQL := "select phonenum from initregistration where phonenum = ?"

	rows, err := tx.QueryContext(ctx, SQL, phonenum)
	helper.PanicIfError(err)
	defer rows.Close()

	initRegistration := domain.InitRegistration{}
	if rows.Next() {
		err := rows.Scan(&initRegistration.PhoneNum, &initRegistration.SecretAnswer)
		helper.PanicIfError(err)
		return initRegistration, nil
	} else {
		return initRegistration, errors.New("Phone number is not found")
	}

}

func (initRegRepo *InitRegistrationRepositoryImpl) UpdateSecretAnswer(ctx context.Context, tx *sql.Tx, category domain.InitRegistration) domain.InitRegistration {
	//TODO implement me
	panic("implement me")
}
*/

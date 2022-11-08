package domain

import "time"

type Registration struct {
	Regid             int
	Phonenum          string
	Email             string
	Idcardnumber      string
	Idcardimage       string
	Idcardname        string
	Idcardaddr1       string
	Idcardaddr2       string
	Idcardsex         string
	Idcardcity        string
	Idcardprovince    string
	Idcardzip         string
	Idcardcountry     string
	Nationality       string
	Job               string
	Sourceofincome    float64
	Taxcardimage      string
	Taxcardnumber     string
	Photoselfieimage  string
	Mothermaindenname string
	Parentname        string
	Parentjob         string
	Registdate        time.Time

	Pagenum int // if pageNum -1 -> then it finished

}

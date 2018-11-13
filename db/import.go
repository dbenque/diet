package db

import (
	"database/sql"
	"database/sql/driver"
	"log"
	"strings"
	"time"

	"github.com/dbenque/diet/api"
	"github.com/dbenque/diet/utils"
	"github.com/go-sqlite/sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

type DB interface {
	GetPersons() []api.Person
	Close() error
}

var _ DB = &DBSqliteImport{}

type DBSqliteImport struct {
	dbFile  *sql.DB
	persons *sqlite3.Table
}

func Open(filepath string) (*DBSqliteImport, error) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}

	return &DBSqliteImport{dbFile: db}, nil
}

func (db *DBSqliteImport) Close() error {
	return db.dbFile.Close()
}

func (db *DBSqliteImport) GetPersons() []api.Person {
	result := []api.Person{}

	rows, err := db.dbFile.Query(`select
 "N° Sécu",
 "Nom Patient",
 "Prénom Patient",
 "Date de naissance",
 "Sexe",
 "Téléphone Domicile",
 "Portable",
 "Email",
 "Allergies",
 "Sports 1",
 "Fréquence Sport 1",
 "Sports 2",
 "Fréquence Sport 2",
 "Sports 3",
 "Fréquence Sport 3",
 "Sports 4",
 "Fréquence Sport 4",
 "Médicament1",
 "Posologie1",
 "Médicament2",
 "Posologie2",
 "Médicament3",
 "Posologie3",
 "Poids Naissance",
 "N° Médecin",
 "Antécédents",
 "Particularités"
 from "tbl Adhérent"`)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var p api.Person
		var tel sql.NullString
		var allergies sql.NullString
		var s1, sf1, s2, sf2, s3, sf3, s4, sf4 sql.NullString //sport
		var m1, mf1, m2, mf2, m3, mf3 sql.NullString          //medoc
		var pSocialNumber, pGender, pPhone, pEmail, pDoctor, pName, pFirstName, pAntecedents, pSpecific sql.NullString
		var pBirthWeight sql.NullFloat64
		var pBirthDate NullDate
		err = rows.Scan(
			&pSocialNumber,
			&pName,
			&pFirstName,
			&pBirthDate,
			&pGender,
			&tel,
			&pPhone,
			&pEmail,
			&allergies,
			&s1, &sf1, &s2, &sf2, &s3, &sf3, &s4, &sf4,
			&m1, &mf1, &m2, &mf2, &m3, &mf3,
			&pBirthWeight,
			&pDoctor,
			&pAntecedents,
			&pSpecific,
		)
		if err != nil {
			log.Fatal(err)
		}
		p.Gender = utils.TextToGender(string(pGender.String))
		p.Name = pName.String
		p.FirstName = pFirstName.String
		p.Phone = pPhone.String
		p.SocialNumber = pSocialNumber.String
		p.Email = pEmail.String
		p.Doctor = pDoctor.String
		p.BirthWeight = pBirthWeight.Float64
		p.BirthDate = pBirthDate.Date
		p.Antecedents = pAntecedents.String
		if p.Antecedents == "Néant" {
			p.Antecedents = ""
		}
		p.Specific = pSpecific.String
		if p.Specific == "Néant" {
			p.Specific = ""
		}

		if p.Phone == "" {
			p.Phone = tel.String
		}
		p.Allergies = strings.Split(allergies.String, " ")
		p.Sports = []string{}
		if s1.String != "" {
			p.Sports = append(p.Sports, strings.TrimSpace(s1.String+" "+sf1.String))
		}
		if s2.String != "" {
			p.Sports = append(p.Sports, strings.TrimSpace(s2.String+" "+sf2.String))
		}
		if s3.String != "" {
			p.Sports = append(p.Sports, strings.TrimSpace(s3.String+" "+sf3.String))
		}
		if s4.String != "" {
			p.Sports = append(p.Sports, strings.TrimSpace(s4.String+" "+sf4.String))
		}
		p.Medecines = []string{}
		if m1.String != "" {
			p.Medecines = append(p.Medecines, strings.TrimSpace(m1.String+" "+mf1.String))
		}
		if m2.String != "" {
			p.Medecines = append(p.Medecines, strings.TrimSpace(m2.String+" "+mf2.String))
		}
		if m3.String != "" {
			p.Medecines = append(p.Medecines, strings.TrimSpace(m3.String+" "+mf3.String))
		}
		result = append(result, p)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return result
}

type NullDate struct {
	Date  time.Time
	Valid bool // Valid is true if Float64 is not NULL
}

// Scan implements the Scanner interface.
func (n *NullDate) Scan(value interface{}) error {
	if value == nil {
		n.Date, n.Valid = time.Unix(0, 0), false
		return nil
	}
	n.Valid = true
	switch value.(type) {
	case time.Time:
		n.Date = value.(time.Time)
	default:
		n.Date = time.Unix(0, 0)
	}
	return nil
}

// Value implements the driver Valuer interface.
func (n NullDate) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Date, nil
}

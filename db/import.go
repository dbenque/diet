package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/dbenque/diet/api"
	"github.com/dbenque/diet/utils"
	_ "github.com/mattn/go-sqlite3"
	"github.com/satori/go.uuid"
)

type DBSqliteImport struct {
	dbFile         *sql.DB
	personsBySecu  map[string]*api.Person
	citiesPerCode  map[int64]*api.City
	doctorsPerCode map[int64]*api.Doctor
	consultPerID   map[string]*api.Consult
}

func Open(filepath string) (*DBSqliteImport, error) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}

	return &DBSqliteImport{
		dbFile:         db,
		personsBySecu:  map[string]*api.Person{},
		citiesPerCode:  map[int64]*api.City{},
		doctorsPerCode: map[int64]*api.Doctor{},
		consultPerID:   map[string]*api.Consult{},
	}, nil
}
func (db *DBSqliteImport) Import() error {
	db.GetCities()
	fmt.Printf("%d Cities imported\n", len(db.citiesPerCode))
	db.GetDoctors()
	fmt.Printf("%d Doctors imported\n", len(db.doctorsPerCode))
	db.GetPersons()
	fmt.Printf("%d Patients imported\n", len(db.personsBySecu))
	db.GetConsults()
	fmt.Printf("%d Consult imported\n", len(db.consultPerID))
	return nil
}

func (db *DBSqliteImport) Close() error {
	return db.dbFile.Close()
}
func (db *DBSqliteImport) GetDoctors() []*api.Doctor {
	result := []*api.Doctor{}
	rows, err := db.dbFile.Query(`select
	"Numéro Médecin",
	"Nom Médecin",
	"Prénom Médecin",
	"Titre",
	"Adresse 1",
	"Adresse 2",
	"Code Ville",
	"Téléphone",
	"Email",
	"Profession"
	from "tbl Médecin"`)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var p api.Doctor
		var pCode sql.NullInt64
		var pCodeVille sql.NullInt64
		var pName, pFirstName, pTitle, pAddr1, pAddr2, pPhone, pEmail, pType sql.NullString
		err = rows.Scan(
			&pCode, &pName, &pFirstName, &pTitle, &pAddr1, &pAddr2, &pCodeVille, &pPhone, &pEmail, &pType,
		)
		if err != nil {
			log.Fatal(err)
		}
		p.ID = uuid.Must(uuid.NewV4()).String()
		p.Name = pName.String
		p.FirstName = pFirstName.String
		p.Title = pTitle.String
		p.Address1 = pAddr1.String
		p.Address2 = pAddr2.String
		if city, ok := db.citiesPerCode[pCodeVille.Int64]; ok {
			p.City = city.Name
			p.PostCode = city.PostCode
		}
		p.Phone = pPhone.String
		p.Email = pEmail.String
		p.Type = pType.String

		result = append(result, &p)
		db.doctorsPerCode[pCode.Int64] = &p
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func (db *DBSqliteImport) GetCities() []*api.City {
	result := []*api.City{}
	rows, err := db.dbFile.Query(`select
	"Code Ville",
	"Nom Ville",
	"Code postal",
	"Département",
	"Pays"
	from "tbl Villes"`)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var p api.City
		var pCode sql.NullInt64
		var pCity, pPostal, pDep, pCountry sql.NullString
		err = rows.Scan(
			&pCode, &pCity, &pPostal, &pDep, &pCountry,
		)
		if err != nil {
			log.Fatal(err)
		}
		p.ID = uuid.Must(uuid.NewV4()).String()
		p.Country = pCountry.String
		p.Name = pCity.String
		p.PostCode = pPostal.String
		p.Department = pDep.String

		result = append(result, &p)
		db.citiesPerCode[pCode.Int64] = &p
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func (db *DBSqliteImport) GetConsults() []*api.Consult {
	result := []*api.Consult{}

	rows, err := db.dbFile.Query(`select
	"N° Sécu",
	"ConsultID",
	"Date de la consultation",
	"Nombril (cm)",
	"Taille (cm)",
	"Poids (kg)",
	"PoidsSouh",
	"Contour taille (cm)",
	"Contour hanches (cm)",
	"Aisselles",
	"Poitrine",
	"Genou droit",
	"Genou gauche",
	"Cuisse droite",
	"Cuisse gauche",
	"Mollet droit",
	"Mollet gauche",
	"Cheville droite",
	"Cheville gauche",
	"Masse maigre",
	"Masse grasse",
	"Choléstérol",
	"Créatine",
	"Remarques Info",
	"Glycémie",
	"DCI"
	from "tbl Données Médicales Adhérents"`)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var p api.Consult
		var pConsultID sql.NullInt64
		var pDate NullDate
		var pSocialNumber, pRemarks sql.NullString
		var pBelly, pHeight, pWeight, pWeightTarget, pLeanMass, pFatMass sql.NullFloat64
		var pWaist, pHip, pArmpits, pBreast, pKneeR, pKneeL, pThighR, pThighL, pCalfR, pCalfL, pAnkleR, pAnkleL sql.NullInt64
		var pCholesterol, pGlucose, pCreatine, pDCI sql.NullFloat64
		err = rows.Scan(
			&pSocialNumber,
			&pConsultID,
			&pDate,
			&pBelly,
			&pHeight,
			&pWeight,
			&pWeightTarget,
			&pWaist,
			&pHip,
			&pArmpits,
			&pBreast,
			&pKneeR,
			&pKneeL,
			&pThighR,
			&pThighL,
			&pCalfR,
			&pCalfL,
			&pAnkleR,
			&pAnkleL,
			&pLeanMass,
			&pFatMass,
			&pCholesterol,
			&pCreatine,
			&pRemarks,
			&pGlucose,
			&pDCI,
		)

		if err != nil {
			log.Fatal(err)
		}

		p.Date = pDate.Date
		p.ID = strconv.Itoa(int(pConsultID.Int64))
		p.DCI = pDCI.Float64
		p.Remarks = pRemarks.String

		p.Measures.AnkleLeft = float64(pAnkleL.Int64)
		p.Measures.AnkleRight = float64(pAnkleR.Int64)
		p.Measures.Armpits = float64(pArmpits.Int64)
		p.Measures.BellyButton = pBelly.Float64
		p.Measures.Breast = float64(pBreast.Int64)
		p.Measures.CalfLeft = float64(pCalfL.Int64)
		p.Measures.CalfRight = float64(pAnkleR.Int64)
		p.Measures.FatMass = pFatMass.Float64
		p.Measures.Height = pHeight.Float64
		p.Measures.Hip = float64(pHip.Int64)
		p.Measures.KneeLeft = float64(pKneeL.Int64)
		p.Measures.KneeRight = float64(pKneeR.Int64)
		p.Measures.LeanMass = pLeanMass.Float64
		p.Measures.ThighLeft = float64(pThighL.Int64)
		p.Measures.ThighRight = float64(pThighR.Int64)
		p.Measures.Waist = float64(pWaist.Int64)
		p.Measures.Weight = pWeight.Float64
		p.Measures.WeightTarget = pWeightTarget.Float64

		p.BloodAnalysis.Cholesterol = pCholesterol.Float64
		p.BloodAnalysis.Creatine = pCreatine.Float64
		p.BloodAnalysis.Glucose = pGlucose.Float64

		if person, ok := db.personsBySecu[pSocialNumber.String]; !ok {
			log.Fatalf("Unknow Social Number for consult %s", pSocialNumber.String)
		} else {
			p.PersonID = person.ID
		}
		//p.PersonID = TODO il faut le retrouver
		result = append(result, &p)
		db.consultPerID[p.ID] = &p
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func (db *DBSqliteImport) GetPersons() []*api.Person {
	result := []*api.Person{}

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
 "Particularités",
 "Remarques",
 "Profession",
 "Régimes Traitements passés"
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
		var pDoctor sql.NullInt64
		var pSocialNumber, pGender, pPhone, pEmail, pName, pFirstName, pAntecedents, pSpecific, pRemark, pPastTreatment, pWork sql.NullString
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
			&pRemark,
			&pPastTreatment,
			&pWork,
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
		if doctor, ok := db.doctorsPerCode[pDoctor.Int64]; ok {
			p.Doctor = doctor.ID
		}
		p.BirthWeight = pBirthWeight.Float64
		p.BirthDate = pBirthDate.Date
		p.Antecedents = pAntecedents.String
		p.Remark = pRemark.String
		p.Work = pWork.String
		p.PastTreatment = pPastTreatment.String
		if p.PastTreatment == "Néant" {
			p.PastTreatment = ""
		}

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
		result = append(result, &p)
		db.personsBySecu[p.SocialNumber] = &p
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

//Diagnostic
//Symptomes / Douleurs / Selles
//Prescription

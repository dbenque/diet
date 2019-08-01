package db

import (
	"strings"

	"github.com/dbenque/diet/api"
)

func PatientFrom(p *api.Person, db *DBSqliteImport) *Patient {
	P := &Patient{}

	P.PatientData.Nom_complet = p.Name + " " + p.FirstName
	P.PatientData.Sexe = string(p.Gender) // Todo
	P.PatientData.Date_de_naissance = p.BirthDate.Format("02/01/2006")
	P.PatientData.Adresse_email = p.Email
	P.PatientData.Portable = p.Phone
	P.PatientData.Pays = ""              // Non Collecté ?
	P.PatientData.Adresse = ""           // Non Collecté ?
	P.PatientData.Code_postal = ""       // Non Collecté ?
	P.PatientData.Numéro_du_dossier = "" // Non Collecté ?
	P.PatientData.NSS = p.SocialNumber   // Non Collecté ?
	P.PatientData.INS = ""               // Non Collecté | Todo
	P.PatientData.NIF = ""               // Non Collecté | Todo
	P.PatientData.Lieu_de_consultation = "L'Hay-les-Roses"
	P.PatientData.Occupation = strings.Join(append(p.Sports, p.Work, p.WorkConditions), "\n")
	P.PatientData.Etiquettes = "" //Todo

	P.PatientConsult.Motif_de_la_consultation = "" //Todo
	P.PatientConsult.Attentes = ""                 //Todo
	P.PatientConsult.Autres_informations = p.Specific

	if cs, ok := db.consultPerID[p.ID]; ok {
		P.PatientConsult.Autres_informations += "\n" + cs.Remarks
	}

	return P
}

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
	P.PatientData.Occupation = PurgeToString(append(p.Sports, p.Work, p.WorkConditions))
	P.PatientData.Etiquettes = "" //Todo

	P.PatientConsult.Motif_de_la_consultation = "" //Todo
	P.PatientConsult.Attentes = ""                 //Todo
	P.PatientConsult.Autres_informations = ""

	if cs, ok := db.ConsultForPerson[p.ID]; ok {
		for _, c := range cs {
			if len(c.Remarks) > 0 && c.Remarks != "Néant" {
				P.PatientConsult.Autres_informations += "\n" + c.Remarks
			}
		}
	}

	P.PatientConsult.Autres_informations = PurgeString(P.PatientConsult.Autres_informations)

	P.PatientHistory.Fumeur = p.Tabac
	P.PatientHistory.etat_civil = "" //Todo Célibataire, Marié, Divorcé, Veuf.
	P.PatientHistory.Autres_informations = p.Specific
	
	P.PatientAlimentation.Aversions = p.Degout
	P.PatientAlimentation.Aliments_preferes = p.Preferences
	
	P.PatientAlimentation.Allergies_et_intolerances = PurgeToString(p.Allergies)
	
	P.PatientMedical.Medication = PurgeToString(p.Medecines)
	P.PatientMedical.Antecedents_personnels = PurgeToString([]string{p.Antecedents,p.PastTreatment})
	
	if p.Doctor!="" {
		doctor,ok := db.DoctorsPerID[p.Doctor]
		if ok && doctor.Name!="INCONNU" {
			P.PatientMedical.Autres_informations = "Docteur: "+doctor.Name+" "+doctor.FirstName+" "+doctor.Phone+" "+doctor.City
		}
	}

	P.PatientMesure.Taille = fmt.Sprintf("%.2f",p.Height)
	poids:=""
	var dateConsult time.Time
	if cs, ok := db.ConsultForPerson[p.ID]; ok {
		for _, c := range cs {
			if c.Date.After(dateConsult) {
			   poids = c.Measures.Weight
			   if c.Measures.Weight != 0 {
				poids = fmt.Sprintf("%.2f Kg",c.Measures.Weight)
			   }
			   if consult taille non vide
			}
		}
	}

	P.PatientMesure.Taille = fmt.Sprintf("%.2f",p.Height)
	return P
}

func PurgeString(in string) string {
	return strings.Join(strings.Split(in, "\n"), "\n")
}

func PurgeToString(in []string) string {
	return strings.Join(Purge(in), "\n")
}

func Purge(in []string) []string {
	n := 0
	for _, x := range in {
		if len(x) > 0 && x != "Néant" {
			in[n] = x
			n++
		}
	}
	return in[:n]
}

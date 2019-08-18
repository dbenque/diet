package db

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/dbenque/diet/api"
)

func PatientFrom(p *api.Person, db *DBSqliteImport) *Patient {
	P := &Patient{}

	P.PatientData.Nom_complet = p.Name + " " + p.FirstName
	P.PatientData.Sexe = string(p.Gender) // Todo
	P.PatientData.Date_de_naissance = p.BirthDate.Format("02/01/2006")
	P.PatientData.Adresse_email = p.Email
	P.PatientData.Portable = p.Phone
	P.PatientData.Pays = "France"
	P.PatientData.Adresse = p.Addr1
	if p.Addr2 != "" {
		P.PatientData.Adresse += "\n" + p.Addr2
	}
	P.PatientData.Code_postal = p.CityCode
	P.PatientData.Numéro_du_dossier = "" // Non Collecté ?
	P.PatientData.NSS = p.SocialNumber
	P.PatientData.INS = "" // Non Collecté | Todo
	P.PatientData.NIF = "" // Non Collecté | Todo
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
	P.PatientMedical.Antecedents_personnels = PurgeToString([]string{p.Antecedents, p.PastTreatment})

	if p.Doctor != "" {
		doctor, ok := db.DoctorsPerID[p.Doctor]
		if ok && doctor.Name != "INCONNU" {
			P.PatientMedical.Autres_informations = "Docteur: " + doctor.Name + " " + doctor.FirstName + " " + doctor.Phone + " " + doctor.City
		}
	}

	poids := ""
	taille := ""
	var dateConsult time.Time
	poidsHistory := []string{}
	hipHistory := []string{}
	waistHistory := []string{}
	breastHistory := []string{}
	heightHistory := []string{}

	if cs, ok := db.ConsultForPerson[p.ID]; ok {
		for _, c := range cs {
			if c.Date.After(dateConsult) {
				dateConsult = c.Date
				if c.Measures.Weight != 0 {
					poids = fmt.Sprintf("%.2f kg", c.Measures.Weight)
				}
			}
			if c.Date.After(dateConsult) {
				dateConsult = c.Date
				if c.Measures.Height != 0 {
					taille = fmt.Sprintf("%.0f cm", c.Measures.Height)
				}
			}

			if c.Measures.Hip > 1 {
				hipHistory = append(hipHistory,
					fmt.Sprintf("%s: %.0fcm", c.Date.Format("2006/01/02"), c.Measures.Hip))
			}

			if c.Measures.Waist > 1 {
				waistHistory = append(waistHistory,
					fmt.Sprintf("%s: %.0fcm", c.Date.Format("2006/01/02"), c.Measures.Waist))
			}

			if c.Measures.Breast > 1 {
				breastHistory = append(breastHistory,
					fmt.Sprintf("%s: %.0fcm", c.Date.Format("2006/01/02"), c.Measures.Breast))
			}

			if c.Measures.Weight > 1 {
				poidsHistory = append(poidsHistory,
					fmt.Sprintf("%s: %.2fkg", c.Date.Format("2006/01/02"), c.Measures.Weight))
			}

			if c.Measures.Height > 1 {
				heightHistory = append(heightHistory,
					fmt.Sprintf("%s: %.0fcm", c.Date.Format("2006/01/02"), c.Measures.Height))
			}

		}
	}
	sort.Strings(hipHistory)
	sort.Strings(waistHistory)
	sort.Strings(breastHistory)
	sort.Strings(poidsHistory)
	sort.Strings(heightHistory)

	// fmt.Printf("TailleV: %v\n", waistHistory)
	// fmt.Printf("Poids: %v\n", poidsHistory)
	// fmt.Printf("Hanche: %v\n", hipHistory)
	// fmt.Printf("Poitrine: %v\n", breastHistory)
	// fmt.Printf("T: %v\n", heightHistory)

	if len(taille) > 0 {
		P.PatientMesure.Taille = taille
	}
	P.PatientMesure.Poids = poids

	fmt.Printf("T:%s / P:%s\n", P.PatientMesure.Taille, P.PatientMesure.Poids)

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

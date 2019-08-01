package db

import (
	"encoding/csv"
	"log"
	"os"
)

type Patient struct {
	PatientData
	PatientConsult
	PatientHistory
	PatientMedical
	PatientAlimentation
	PatientMesure
}
type PatientData struct {
	Nom_complet          string
	Sexe                 string
	Date_de_naissance    string
	Adresse_email        string
	Portable             string
	Pays                 string
	Adresse              string
	Code_postal          string
	Numéro_du_dossier    string
	NSS                  string
	INS                  string
	NIF                  string
	Lieu_de_consultation string
	Occupation           string
	Etiquettes           string
}

func (p *PatientData) Strings() []string {
	return []string{
		p.Nom_complet,
		p.Sexe,
		p.Date_de_naissance,
		p.Adresse_email,
		p.Portable,
		p.Pays,
		p.Adresse,
		p.Code_postal,
		p.Numéro_du_dossier,
		p.NSS,
		p.INS,
		p.NIF,
		p.Lieu_de_consultation,
		p.Occupation,
		p.Etiquettes,
	}
}

var PatientDataHeader = []string{"Nom complet", "Sexe", "Date de naissance", "Adresse e-mail", "Portable", "Pays", "Adresse", "Code postal", "Numéro du dossier", "NSS", "INS", "NIF", "Lieu de consultation", "Occupation", "Étiquettes"}

type PatientConsult struct {
	Motif_de_la_consultation string
	Attentes                 string
	Autres_informations      string
}

func (p *PatientConsult) Strings() []string {
	return []string{
		p.Motif_de_la_consultation,
		p.Attentes,
		p.Autres_informations,
	}
}

var PatientConsultHeader = []string{"Motif de la consultation", "Attentes", "Autres informations"}

type PatientHistory struct {
	// Origine
	Fumeur                                                  string
	Information_complementaire_sur_la_consommation_de_tabac string
	Qualite_du_sommeil                                      string
	Information_complementaire_sur_la_qualite_du_sommeil    string
	Consommation_d_alcool                                   string
	Information_complementaire_sur_la_consommation_d_alcool string
	etat_civil                                              string
	Information_complementaire_sur_l_etat_civil             string
	Fonction_intestinale                                    string
	Information_complementaire_sur_la_fonction_intestinale  string
	Autres_informations                                     string
}

func (p *PatientHistory) Strings() []string {
	return []string{
		p.Fumeur,
		p.Information_complementaire_sur_la_consommation_de_tabac,
		p.Qualite_du_sommeil,
		p.Information_complementaire_sur_la_qualite_du_sommeil,
		p.Consommation_d_alcool,
		p.Information_complementaire_sur_la_consommation_d_alcool,
		p.etat_civil,
		p.Information_complementaire_sur_l_etat_civil,
		p.Fonction_intestinale,
		p.Information_complementaire_sur_la_fonction_intestinale,
		p.Autres_informations,
	}
}

var PatientHistoryHeader = []string{
	"Origine",
	"Fumeur",
	"Information complémentaire sur la consommation de tabac",
	"Qualité du sommeil",
	"Information complémentaire sur la qualité du sommeil",
	"Consommation d'alcool",
	"Information complémentaire sur la consommation d'alcool",
	"État civil",
	"Information complémentaire sur l'état civil",
	"Fonction intestinale",
	"Information complémentaire sur la fonction intestinale",
	"Autres informations",
}

type PatientMedical struct {
	Pathologies            string
	Medication             string
	Antecedents_personnels string
	Antecedents_familiaux  string
	Autres_informations    string
}

func (p *PatientMedical) Strings() []string {
	return []string{
		p.Pathologies,
		p.Medication,
		p.Antecedents_personnels,
		p.Antecedents_familiaux,
		p.Autres_informations,
	}
}

var PatientMedicalHeader = []string{
	"Pathologies",
	"Médication",
	"Antécédents personnels",
	"Antécédents familiaux",
	"Autres informations",
}

type PatientAlimentation struct {
	Heure_de_lever            string
	Heure_de_coucher          string
	Aliments_preferes         string
	Aversions                 string
	Allergies_et_intolerances string
	Consommation_d_eau        string
	Autres_informations       string
}

func (p *PatientAlimentation) Strings() []string {
	return []string{
		p.Heure_de_lever,
		p.Heure_de_coucher,
		p.Aliments_preferes,
		p.Aversions,
		p.Allergies_et_intolerances,
		p.Consommation_d_eau,
		p.Autres_informations,
	}
}

var PatientAlimentationHeader = []string{
	"Heure de lever",
	"Heure de coucher",
	"Aliments préférés",
	"Aversions",
	"Allergies et intolérances",
	"Consommation d'eau",
	"Autres informations",
}

type PatientMesure struct {
	Poids  string
	Taille string
}

func (p *PatientMesure) Strings() []string {
	return []string{
		p.Poids,
		p.Taille,
	}
}

var PatientMesureHeader = []string{"Poids", "Taille"}

func Headers() []string {
	h := append(PatientDataHeader, PatientConsultHeader...)
	h = append(h, PatientHistoryHeader...)
	h = append(h, PatientMedicalHeader...)
	h = append(h, PatientAlimentationHeader...)
	h = append(h, PatientMesureHeader...)
	return h
}

func (p *Patient) Strings() []string {
	h := append(p.PatientData.Strings(), p.PatientConsult.Strings()...)
	h = append(h, p.PatientHistory.Strings()...)
	h = append(h, p.PatientMedical.Strings()...)
	h = append(h, p.PatientAlimentation.Strings()...)
	h = append(h, p.PatientMesure.Strings()...)
	return h

}

func Export(patients []*Patient) {
	records := [][]string{
		Headers(),
	}

	for _, p := range patients {
		records = append(records, p.Strings())
	}

	w := csv.NewWriter(os.Stdout)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

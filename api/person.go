package api

import (
	"time"
)

type Person struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	FirstName    string `json:"firstName"`
	Gender       Gender `json:"gender"`
	SocialNumber string `json:"socialNumber"`

	BirthDate   time.Time `json:"birthDate"`
	BirthWeight float64   `json:"birthWeight"`

	Height float64 `json:"height"`

	Email string `json:"email"`
	Phone string `json:"phone"`

	Sports         []string `json:"sports"`
	WorkConditions string   `json:"workConditions"`

	Doctor        string   `json:"doctor"`
	Antecedents   string   `json:"Antecedents"`
	Allergies     []string `json:"allergies"`
	Medecines     []string `json:"medecines"`
	PastTreatment string   `json:"pastReatments"`
	Specific      string   `json:"specific"`
}

//Gender enum
type Gender string

//Genders
const (
	GenderMale       = "Male"
	GenderFemale     = "Female"
	GenderUnknown    = "Unknown"
	GenderNonBinanry = "NonBinary"
)

//INSERT INTO `tbl Adhérent` (`N° Sécu`, `Civilité`, `Nom Patient`, `Prénom Patient`, `Date de naissance`, `SexeOptionN°`, `Sexe`, `EtatCivilOptionN°`, `Etat Civil`, `Adresse 1`, `Adresse 2`, `Code Ville`, `Téléphone Domicile`, `Téléphone Bureau`, `Portable`, `Fax`, `Email`, `Préférences`, `Dégoûts`, `Allergies`, `Intéraction médicaments`, `Alimentation`, `Dentition`, `Déglutition`, `Mastication`, `Diagnostic`, `Prescription`, `Antécédents`, `Particularités`, `Symptômes / Douleurs / Selles`, `Sports 1`, `Fréquence Sport 1`, `Sports 2`, `Fréquence Sport 2`, `Sports 3`, `Fréquence Sport 3`, `Sports 4`, `Fréquence Sport 4`, `Profession`, `Conditions de travail`, `Remarques`, `Médicament1`, `Posologie1`, `Médicament2`, `Posologie2`, `Médicament3`, `Posologie3`, `N° Médecin`, `Poids Naissance`, `Poids Adolescence`, `Poids Adulte`, `Poids Avant Grossesse`, `Poids Après Grossesse`, `Régimes Traitements passés`, `Tabac`, `Nombre d'enfants`, `Photo`)
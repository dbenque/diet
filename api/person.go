package api

import (
	"time"
)

type ConsultMeasure struct {
	BellyButton  float64 `json:"belly"`
	Height       float64 `json:"height"`
	Weight       float64 `json:"weight"`
	WeightTarget float64 `json:"weightTarget"`
	Waist        float64 `json:"waist"`   // taille
	Hip          float64 `json:"hip"`     // hanche
	Armpits      float64 `json:"armpits"` //aisselle
	Breast       float64 `json:"breast"`  //poitrine
	KneeRight    float64 `json:"kneeRight"`
	KneeLeft     float64 `json:"kneeLeft"`
	ThighRight   float64 `json:"thighRight` //Cuisse
	ThighLeft    float64 `json:"thighLeft`
	CalfRight    float64 `json:"calfRight` //Mollet
	CalfLeft     float64 `json:"calfLeft`
	AnkleRight   float64 `json:"ankleRight` // Cheville
	AnkleLeft    float64 `json:"anfleLeft`
	LeanMass     float64 `json:"leanMass"`
	FatMass      float64 `json:"fatMass"`
}
type BloodAnalysis struct {
	Cholesterol float64 `json:"cholesterol"`
	Glucose     float64 `json:"glucose"`
	Creatine    float64 `json:"creatine"`
}
type Consult struct {
	ID            string         `json:"id"`
	Date          time.Time      `json:"date"`
	PersonID      string         `json:"person"`
	Measures      ConsultMeasure `json:"measures"`
	BloodAnalysis BloodAnalysis  `json:"bloodAnalysis"`
	DCI           float64        `json:"DCI"`
	Remarks       string         `json:"remarks"`
}

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
	Work           string   `json:"work"`

	Doctor        string   `json:"doctor"`
	Antecedents   string   `json:"Antecedents"`
	Allergies     []string `json:"allergies"`
	Medecines     []string `json:"medecines"`
	PastTreatment string   `json:"pastReatments"`
	Specific      string   `json:"specific"`

	Remark string `json:"remark"`
}

type Doctor struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	FirstName string `json:"firstName"`
	Title     string `json:"title"`
	Address1  string `json:"address1"`
	Address2  string `json:"address2"`
	City      string `json:"city"`
	PostCode  string `json:"postCode"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Type      string `json:"type"`
}

type City struct {
	ID         string `json:"id"`
	Name       string `json:"city"`
	PostCode   string `json:"postCode"`
	Department string `json:"department"`
	Country    string `json:"country"`
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

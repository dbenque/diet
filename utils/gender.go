package utils

import "github.com/dbenque/diet/api"

func TextToGender(t string) api.Gender {
	switch {
	case t == "Masculin" || t == "Male":
		{
			return api.GenderMale
		}
	case t == "Féminin" || t == "Female":
		{
			return api.GenderFemale
		}
	case t == "NonBinaire" || t == "NonBinary":
		{
			return api.GenderNonBinanry
		}
	default:
		return api.GenderUnknown
	}
}

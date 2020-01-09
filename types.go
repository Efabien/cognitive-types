package cognitivetypes

type Raw map[string]map[string][]string

type Intent struct {
	Texts [][]string
	Treshold float32
}

type Intents map[string]Intent

type Keyword map[string][][]string

type Keywords map[string]Keyword

type WordWeigth struct {
	Word string
	Weigth float32	
}

type IntentWeigths struct {
	Intent string
	Weigths []WordWeigth
}

type Detection struct {
	Intent string
	Matchs []string
	Score float32
}

type Vault struct {
	Intents Intents
	Keywords Keywords
	Scope int
	Degree int
	Weigths []IntentWeigths
}
package model

type Project struct {
	ID          int    `json:"id"`
	ProjectName string `json:"project_name"`
	Category    string `json:"category"`
	ProjectType string `json:"project_type"`
	ReleaseYear int    `json:"release_year"`
	AgeCategory string `json:"age_category"`
	Duration    string `json:"duration"`
	Director    string `json:"director"`
	Producer    string `json:"producer"`
}

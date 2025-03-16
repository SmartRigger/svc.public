package response

type ApiHTTPGetToolGet struct {
	ID                       string `json:"id" yaml:"id" xml:"id"`
	Name                     string `json:"name" yaml:"name" xml:"name"`
	Category                 string `json:"category" yaml:"category" xml:"category"`
	Range                    string `json:"range" yaml:"range" xml:"range"`
	InspectionDate           string `json:"inspection_date" yaml:"inspection_date" xml:"inspection_date"`
	InspectionPersonnel      string `json:"inspection_personnel" yaml:"inspection_personnel" xml:"inspection_personnel"`
	DailyStartupVerification bool   `json:"daily_startup_verification" yaml:"daily_startup_verification" xml:"daily_startup_verification"`
	SelfMutualInspection     bool   `json:"self_mutual_inspection" yaml:"self_mutual_inspection" xml:"self_mutual_inspection"`
	MeasurementUnit          string `json:"measurement_unit" yaml:"measurement_unit" xml:"measurement_unit"`
	RegistrationDate         string `json:"registration_date" yaml:"registration_date" xml:"registration_date"`
}

type ApiHTTPGetTool struct {
	ID                       string `json:"id" yaml:"id" xml:"id"`
	Name                     string `json:"name" yaml:"name" xml:"name"`
	Category                 string `json:"category" yaml:"category" xml:"category"`
	Range                    string `json:"range" yaml:"range" xml:"range"`
	InspectionDate           string `json:"inspection_date" yaml:"inspection_date" xml:"inspection_date"`
	MeasurementUnit          string `json:"measurement_unit" yaml:"measurement_unit" xml:"measurement_unit"`
	RegistrationDate         string `json:"registration_date" yaml:"registration_date" xml:"registration_date"`
	CreatedAt                string `json:"created_at" yaml:"created_at" xml:"created_at"`
	DailyStartupVerification bool   `json:"daily_startup_verification" yaml:"daily_startup_verification" xml:"daily_startup_verification"`
	SelfMutualInspection     bool   `json:"self_mutual_inspection" yaml:"self_mutual_inspection" xml:"self_mutual_inspection"`
}

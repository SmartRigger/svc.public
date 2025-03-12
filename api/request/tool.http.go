/*
 * Copyright (C) HereweTech, Inc - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 */

/**
 * @file tool.http.go
 * @package request
 * @author Dr.NP <conan.np@gmail.com>
 * @since 03/13/2025
 */

package request

/* {{{ [svc.api/tool] */
type ApiHTTPPostToolCreate struct {
	Name                     string `json:"name" xml:"name" form:"name"`
	Category                 string `json:"category" xml:"category" form:"category"`
	Range                    string `json:"range" xml:"range" form:"range"`
	InspectionDate           string `json:"inspection_date" xml:"inspection_date" form:"inspection_date"`
	InspectionPersonnel      string `json:"inspection_personnel" xml:"inspection_personnel" form:"inspection_personnel"`
	DailyStartupVerification bool   `json:"daily_startup_verification" xml:"daily_startup_verification" form:"daily_startup_verification"`
	SelfMutualInspection     bool   `json:"self_mutual_inspection" xml:"self_mutual_inspection" form:"self_mutual_inspection"`
	MeasurementUnit          string `json:"measurement_unit" xml:"measurement_unit" form:"measurement_unit"`
	RegistrationDate         string `json:"registration_date" xml:"registration_date" form:"registration_date"`
}

type ApiHTTPPostToolUpdate struct {
	Name                     string `json:"name" xml:"name" form:"name"`
	Category                 string `json:"category" xml:"category" form:"category"`
	Range                    string `json:"range" xml:"range" form:"range"`
	InspectionDate           string `json:"inspection_date" xml:"inspection_date" form:"inspection_date"`
	InspectionPersonnel      string `json:"inspection_personnel" xml:"inspection_personnel" form:"inspection_personnel"`
	DailyStartupVerification bool   `json:"daily_startup_verification" xml:"daily_startup_verification" form:"daily_startup_verification"`
	SelfMutualInspection     bool   `json:"self_mutual_inspection" xml:"self_mutual_inspection" form:"self_mutual_inspection"`
	MeasurementUnit          string `json:"measurement_unit" xml:"measurement_unit" form:"measurement_unit"`
	RegistrationDate         string `json:"registration_date" xml:"registration_date" form:"registration_date"`
}

type ApiHTTPGetToolList struct {
	Page  int `json:"page" xml:"page" form:"page"`
	Limit int `json:"limit" xml:"limit" form:"limit"`
}

/* }}} */

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */

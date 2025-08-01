/*
 * Copyright (C) HereweTech, Inc - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 */

/**
 * @file svc.device.http.go
 * @package request
 * @author Dr.NP <conan.np@gmail.com>
 * @since 03/10/2025
 */

package request

type DeviceHTTPPostMeasure struct {
	DeviceID string `json:"device_id" xml:"device_id" form:"device_id"`
	Value    string `json:"value" xml:"value" form:"value"`
	Detail   string `json:"detail,omitempty" xml:"detail,omitempty" form:"detail,omitempty"`
}

type AdjustSpanner struct {
	IP     string `json:"ip" xml:"ip" form:"ip"`
	Pset   uint8  `json:"pset" xml:"pset" form:"pset"`
	Torque uint16 `json:"torque" xml:"torque" form:"torque"`
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */

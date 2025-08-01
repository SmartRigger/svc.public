/*
 * Copyright (C) HereweTech, Inc - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 */

/**
 * @file svc.device.http.go
 * @package response
 * @author Dr.NP <conan.np@gmail.com>
 * @since 03/10/2025
 */

package response

type DeviceHTTPPostMeasure struct {
	Status bool `json:"status" xml:"status" form:"status"`
}

type AdjustSpanner struct {
	Result bool `json:"result" xml:"result" form:"result"`
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */

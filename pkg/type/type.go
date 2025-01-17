package typeResult

type (
	SuccessResult struct {
		// ResponseCode string      `json:"code"`
		Status int         `json:"code"`
		Data   interface{} `json:"data"`
	}
	SuccessReauth struct {
		// ResponseCode string      `json:"code"`
		Status int         `json:"code"`
		Data   interface{} `json:"message"`
	}

	ErrorResult struct {
		// ResponseCode string `json:"code"`
		Status  int    `json:"code"`
		Message string `json:"message"`
	}

	ErrorResultV2 struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		// Error   interface{} `json:"err"`
	}
	ErrorResultDto struct {
		// ResponseCode string `json:"code"`
		Code    int    `json:"code"`
		Message string `json:"message"`
		Error   error  `json:"error"`
	}
)

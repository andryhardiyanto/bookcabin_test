package fibers

type ErrorType string

var (
	ErrorNotFound            = NewError(404, "not found", "not_found")
	ErrorInternalServerError = NewError(500, "internal server error", "internal_server_error")
	ErrorBadRequest          = NewError(400, "bad request", "bad_request")
	ErrorVoucherAlreadyExist = NewError(400, "voucher already exist", "voucher_already_exist")
	ErrorPageNotFound        = NewError(404, "page not found", "page_not_found")
	ErrorInvalidDate         = NewError(400, "invalid date", "invalid_date")
)

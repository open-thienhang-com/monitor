package api

const (
	PermissionDenied     = "permission denied"
	WrongID              = "wrong id"
	OperationNotAllow    = "operation not allow"
	EditFailWrongToken   = "edit fail, wrong token"
	CreateFailWrongToken = "create fail, wrong token"
	NoPermission         = "no permission"
	SiteOff              = "site is off"
)

func WrongPK(pk string) string {
	return "wrong " + pk
}

type PageError error

var (
	PageError404 PageError
	PageError500 PageError
	PageError403 PageError
	PageError401 PageError
)

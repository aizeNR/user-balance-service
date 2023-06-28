package model

type SortDirection string

const (
	Asc  = "ASC"
	Desc = "DESC"
)

var SortMap = map[string]SortDirection{
	"asc":  Asc,
	"desc": Desc,
}

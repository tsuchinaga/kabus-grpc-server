package repositories

type Setting interface {
	IsProduction() bool
	Password() string
}

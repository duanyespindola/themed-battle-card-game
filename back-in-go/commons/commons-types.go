package commons

// commons
type TUUID string

type IWithId interface {
	Id() TUUID
}

type TStatus int

type IWithStatus interface {
	Status() TStatus
}

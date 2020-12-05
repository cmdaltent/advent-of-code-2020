package common

type Passport interface {
	IsValid() bool
	MergeWith(key, value string)
}

type PassportParsed func(pp Passport)
type GeneratePassport func() Passport

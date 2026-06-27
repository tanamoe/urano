package search

type SearchRegistryRequest struct {
	RegistrationID *string
	ISBN           *string
	Title          *string
	Author         *string
	Translator     *string
	PrintAmount    *QueryNumericRange[int]
	SelfPublish    *bool
	Partner        *string
}

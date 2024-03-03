package response

type ResponseData struct{
	Country string `json:"Country"`
	OverallSpeed int `json:"OverallSpeed"`
	CountryId int `json:"CountryId"`
}
type StateResponse struct{
	StateId int
	State string
	OverallSpeed int
}
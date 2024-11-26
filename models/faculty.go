package model


type JSONAvailability struct{
	Day string
	TimePeriod map[string]string
}
type Faculty struct{
	Id int 
	Name string
	Email string
	Password string
	Max_Hours int
	Availability []JSONAvailability
	Create_At string
}
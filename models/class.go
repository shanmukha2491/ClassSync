package model


type Timings struct{
	Day string
	TimePeriod map[string]string
}
type Class struct{
	Id int 
	Subject string
	Hours_Per_Week int 
	Preferred_Timing []Timings
	Faculty_Id int
}
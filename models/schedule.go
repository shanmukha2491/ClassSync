package model

import "time"


type Schedule struct{
	Id int 
	Class_Id int
	Faculty_Id int 
	Day string
	StartTime time.Time
	EndTime time.Time
}
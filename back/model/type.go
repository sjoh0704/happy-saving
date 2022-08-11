package model

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
)

type Phase string

const (
	Coupled  Phase = "Coupled"
	Awaiting Phase = "Awaiting"
	Denyed   Phase = "Denyed"
)

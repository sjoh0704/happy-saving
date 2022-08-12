package model

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
)

type Phase string

const (
	Approved  Phase = "Approved"
	Awaiting Phase = "Awaiting"
	Denyed   Phase = "Denyed"
)

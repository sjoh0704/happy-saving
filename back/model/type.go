package model

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

type Phase string

const (
	Approved  Phase = "approved"
	Awaiting Phase = "awaiting"
	Denyed   Phase = "denyed"
)

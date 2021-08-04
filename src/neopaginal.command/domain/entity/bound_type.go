package entity

type BoundType string

const (
	InitialWithoutBound BoundType = "InitialWithoutBound"
	InBound BoundType = "InBound"
	OutBound BoundType = "OutBound"
)
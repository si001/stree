package model

type Settings struct {
	FileMask string
	FileMode int
	OrderBy  byte
}

var DataSettings Settings = Settings{
	FileMask: "*.*",
	FileMode: 2,
	OrderBy:  0,
}

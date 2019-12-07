package model

import (
	"time"
)

type FileInfo struct {
	Name    string
	Size    int64
	ModTime time.Time
	Attr    byte
}

const (
	ATTR_NOTREAD     = 0
	ATTR_FILE        = 1
	ATTR_DIR         = 2
	ATTR_ARCH        = 4
	ATTR_ERR_MESSAGE = 255
)

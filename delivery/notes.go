package delivery

import (
	svc "commonauthsvc/service"
)

type NotesHTTPHandler struct {
	NoteSvc svc.NotesSvc
}

func
package delivery

import (
	svc "commonauthsvc/service"
	"github.com/labstack/echo"
)

type NotesHTTPHandler struct {
	NoteSvc svc.NotesSvc
}

func ConfigureNoteHttpHandler(e *echo.Echo, noteSvc svc.NotesSvc) {
	notesHttpHandler := NotesHTTPHandler{
		NoteSvc: noteSvc,
	}
	notesHttpHandler.AddHandlers(e)
}

func (noteHttp *NotesHTTPHandler) AddHandlers(e *echo.Echo) {
	e.POST("/createNote", noteHttp.createNotes)
}

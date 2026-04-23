package models

import "time"

type ScriptType string

const (
	ScriptLua     ScriptType = "lua"
	ScriptPython  ScriptType = "python"
	ScriptScratch ScriptType = "scratch"
)

type Script struct {
	Id        string
	ProjectID string
	Name      string
	Type      ScriptType

	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

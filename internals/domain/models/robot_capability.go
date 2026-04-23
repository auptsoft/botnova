package models

type CommandDefinition struct {
	Name       string
	Parameters []ParameterDefinition
}

type ParameterDefinition struct {
	Name string
	Type string // "int", "float", "string", etc.
}

type PropertyDefinition struct {
	Name     string
	Type     string
	ReadOnly bool
}

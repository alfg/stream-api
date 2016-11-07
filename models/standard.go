package models

// Index Index json view.
type Index struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Version     string `json:"version,omitempty"`
	Homepage    string `json:"homepage,omitempty"`
	Bugs        string `json:"bugs,omitempty"`
	Docs        string `json:"docs,omitempty"`
}

package models

// Stream contains stream models.
type Stream struct {
	ID          int64  `db:"id" json:"id,omitempty"`
	Title       string `db:"title" json:"title,omitempty" valid:"required"`
	Type        string `db:"type" json:"type,omitempty" valid:"alphanum,required"`
	Description string `db:"description" json:"description,omitempty" valid:"-"`
	Private     bool   `db:"private" json:"private,omitempty" valid:"-"`
	StreamName  string `db:"stream_name" json:"stream_name,omitempty" valid:"alphanum,required"`
	StreamKey   string `db:"stream_key" json:"stream_key,omitempty" valid:"alphanum,required"` // StreamKey?
	StreamRTMP  string `json:"stream_rtmp_url"`
}

type StreamPrivate struct {
	ID          int64  `db:"id" json:"id,omitempty"`
	Title       string `db:"title" json:"title" valid:"required"`
	Type        string `db:"type" json:"type" valid:"alphanum,required"`
	Description string `db:"description" json:"description" valid:"-"`
	Private     bool   `db:"private" json:"private" valid:"-"`
	StreamName  string `db:"stream_name" json:"stream_name" valid:"alphanum,required"`
	StreamKey   string `db:"stream_key" json:"-" valid:"alphanum,required"` // StreamKey?
	StreamURL   string `json:"stream_url"`
}

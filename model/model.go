package model

type Index struct {
	Order      int32
	Name       string
	Unique     bool
	Type       string
	ContainKey string
	IndexType  string
	Comment    string
}

type Column struct {
	Order        int32
	Name         string
	Type         string
	CanNull      string
	DefaultValue string
	Comment      string
}

type Table struct {
	Name    string
	Comment string
	Charset string
	Engine  string
	Columns []*Column
	Indexes []*Index
	Sql     string
}

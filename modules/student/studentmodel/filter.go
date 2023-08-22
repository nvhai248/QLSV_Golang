package studentmodel

type Filter struct {
	Name string `form:"name" json:"name,omitempty"`
}

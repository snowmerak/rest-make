package spec

func Init() *Spec {
	s := &Spec{}
	s.Routes = append(s.Routes, []struct {
		Method string `yaml:"method"`
		Path   string `yaml:"path"`
	}{
		{
			Method: "GET",
			Path:   "/",
		},
		{
			Method: "GET",
			Path:   "/user",
		},
		{
			Method: "GET",
			Path:   "/user/{id}",
		},
		{
			Method: "POST",
			Path:   "/user",
		},
		{
			Method: "PUT",
			Path:   "/user/{id}",
		},
		{
			Method: "DELETE",
			Path:   "/user/{id}",
		},
	}...)

	return s
}

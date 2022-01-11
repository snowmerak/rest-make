package spec

import "fmt"

func Init() *Spec {
	s := &Spec{}
	s.Routes = append(s.Routes, []Route{
		{
			Method:     "GET",
			Path:       "/",
			Middleware: []string{},
		},
		{
			Method: "GET",
			Path:   "/user",
			Middleware: []string{
				"AdminMiddleare",
			},
		},
		{
			Method: "GET",
			Path:   "/user/{id}",
			Middleware: []string{
				"UserMiddleare",
			},
		},
		{
			Method: "POST",
			Path:   "/user",
			Middleware: []string{
				"AdminMiddleare",
			},
		},
		{
			Method: "PUT",
			Path:   "/user/{id}",
			Middleware: []string{
				"UserMiddleare",
			},
		},
		{
			Method: "DELETE",
			Path:   "/user/{id}",
			Middleware: []string{
				"AdminMiddleare",
			},
		},
	}...)

	return s
}

func (s *Spec) Parse() {
	middlewares := map[string]struct{}{}
	paths := map[string][]struct {
		method       string
		middleware   []string
		rawPath      string
		getParameter []string
	}{}

	for _, route := range s.Routes {
		paths[route.Path] = append(paths[route.Path], struct {
			method       string
			middleware   []string
			rawPath      string
			getParameter []string
		}{
			method:       route.Method,
			middleware:   route.Middleware,
			rawPath:      route.Path,
			getParameter: []string{},
		})
		startIndex := 0
		sw := false
		for i, gp := range route.Path {
			if gp == '{' && !sw {
				sw = true
				startIndex = i
			} else if gp == '}' && sw {
				sw = false
				paths[route.Path][len(paths[route.Path])-1].getParameter = append(paths[route.Path][len(paths[route.Path])-1].getParameter, route.Path[startIndex+1:i])
			}
		}
		for _, m := range route.Middleware {
			middlewares[m] = struct{}{}
		}
	}

	fmt.Println(paths)
}

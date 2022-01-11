package spec

type Route struct {
	Method     string   `yaml:"method"`
	Path       string   `yaml:"path"`
	Middleware []string `yaml:"middleware"`
}

type Spec struct {
	Routes []Route `yaml:"routes"`
}

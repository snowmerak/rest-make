package spec

type Spec struct {
	Routes []struct {
		Method string `yaml:"method"`
		Path   string `yaml:"path"`
	} `yaml:"routes"`
}

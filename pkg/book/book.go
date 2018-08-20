package Book

type Book struct {
	URL         string `yaml:"url"`
	Method      string `yaml:"method"`
	Client      string `yaml:"client"`
	Concurrency string `yaml:"concurrency"`
	Duration    string `yaml:"duration"`
}

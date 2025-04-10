package structs

// cli command struct

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

type Config struct {
	Next     string
	Previous string
}

type LocationArea struct {
	Count    int
	Next     string
	Previous string
	Results  []struct {
		Name string
		URL  string
	}
}

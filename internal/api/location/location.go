package location

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/darginmathi/pokedexcli/structs"
)

func GetLocation(cfg *structs.Config, call string) error {
	// set url to next or prev based on call
	url := ""
	if call == "next" {
		url = cfg.Next
	} else if call == "prev" {
		url = cfg.Previous
	} else {
		return fmt.Errorf("invalid direction: %v", call)
	}
	// gets data and decodes them (by default 20 set by API)
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var area structs.LocationArea

	if err := json.Unmarshal(data, &area); err != nil {
		return err
	}
	// prints all available locations
	for _, result := range area.Results {
		fmt.Println(result.Name)
	}

	// modify cfg
	cfg.Next = area.Next
	cfg.Previous = area.Previous

	return nil
}

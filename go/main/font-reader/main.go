package main

import (
	"fmt"
	"os"

	"golang.org/x/image/font/sfnt"
)

func main() {
	fontFile := "/usr/share/fonts/fonts-go/Go-Regular.ttf"
	bytes, _ := os.ReadFile(fontFile)
	font, _ := sfnt.Parse(bytes)

	fontInfo := []struct {
		nameID      sfnt.NameID
		description string
		data        string
	}{
		{
			nameID:      sfnt.NameIDFamily,
			description: "Family",
		},
		{
			nameID:      sfnt.NameIDSubfamily,
			description: "Subfamily",
		},
		{
			nameID:      sfnt.NameIDFull,
			description: "Full name",
		},
		{
			nameID:      sfnt.NameIDTypographicFamily,
			description: "Typographic Family",
		},
		{
			nameID:      sfnt.NameIDTypographicSubfamily,
			description: "Typographic Subfamily",
		},
		{
			nameID:      sfnt.NameIDWWSFamily,
			description: "WWS Family",
		},
		{
			nameID:      sfnt.NameIDWWSSubfamily,
			description: "WWS Subfamily",
		},
	}

	var b sfnt.Buffer
	for i := range fontInfo {
		item := &fontInfo[i]
		data, err := font.Name(&b, item.nameID)
		if err != nil {
			fmt.Fprintf(os.Stdout, "Error: %s (name ID: %d)\n", err, item.nameID)
			continue
		}
		item.data = data
	}

	for _, item := range fontInfo {
		fmt.Printf("%+v\n", item)
	}
}

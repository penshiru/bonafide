package main

//Tag holds then regex used to parse a law file
type Tag struct {
	name  string
	regex string
}

//NewIntroTags returns an array of Intro Tags
func NewIntroTags() []Tag {
	return []Tag{
		Tag{"Name", "^LEY DE|^C(Ó?|O?)DIGO"},
		Tag{"Number", "No\\.|N°."},
		Tag{"Aproved", "Aprobad(a|o)"},
		Tag{"Journal", "Publicad(a|o)"}, //this also fills publish date, to save a cycle.
		Tag{"Arto", `^\f?(?:Art.\s\d+|Arto.\s\d+|Artículo\s\d+|Articulo\s\d+)`},
	}
}

//NewBodyTags returns an array of Body Tags to parse a Law
func NewBodyTags() []Tag {
	return []Tag{
		Tag{"Titulo", "^\f?(TÍTULO\\s?([IVX\u00C0-\u00FF]|$)|^TITULO\\s?([IVX\u00C0-\u00FF]|$)|TITULO\\s\\w+$|TÍTULO\\s\\w+$)"},
		Tag{"Capitulo", "^\f?(?:Capítulo\\s[\u00C0-\u00FF]?\\w+$|Capí?tulo\\s?\\w{0,3}$|Capitulo\\s?\\w{0,3}$|CAP(Í?|I?)TULO\\s?)"},
		Tag{"Arto", `^\f?(?:Art.\s\d+|Arto.\s\d+|Artículo\s\d+|Articulo\s\d+)`},
		Tag{"Libro", "^\f?(LIBRO\\s[IVXLCDM]+$|LIBRO\\s\\w+$)"},
	}
}

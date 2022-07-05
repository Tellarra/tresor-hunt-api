package main

type hiddenMickey struct {
	ID         string `json:"id_mickey"`
	Land       string `json:"land"`
	PlaceWhere string `json:"place_where"`
	TypeOf     string `json:"type_of"`
}

// hiddenMickey slice to seed record mickey data.
var hiddenMickeys = []hiddenMickey{
	{ID: "1", Land: "Main street U.S.A", PlaceWhere: "Emporium", TypeOf: "Rembarde"},
}

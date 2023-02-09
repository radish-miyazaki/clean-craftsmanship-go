package rental

type VideoType int

var VideoRegistry = map[string]Movie{
	"RegularMovie":  &RegularMovie{Title: "RegularMovie"},
	"ChildrenMovie": &ChildrenMovie{Title: "ChildrenMovie"},
}

func GetType(title string) Movie {
	return VideoRegistry[title]
}

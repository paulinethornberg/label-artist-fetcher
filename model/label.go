package model

type Label struct {
	LabelName string   `json:"label_name"`
	Artists   []Artist `json:"artists"`
}

type Artist struct {
	Name string `json:"name"`
}

func NewLabel(name string, artists []Artist) Label {
	return Label{
		LabelName: name,
		Artists:   artists,
	}
}

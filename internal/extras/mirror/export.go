package mirror

type Field struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Hidden      bool   `json:"hidden"`
	/*
		possible values are `string`, `number`, `bool`, `timestamp`, `related`
		`?` postfix means nullable type
		`[]` postfix means array type; use only with `related`
	*/
	Type            string  `json:"type_name"`
	RelatedModel    *string `json:"related_model"`
	RelatedModelKey *string `json:"related_model_key"`
	RelationType    *string `json:"relation_type"`
}

type Model struct {
	Name   string  `json:"name"`
	Fields []Field `json:"fields"`
}

type Scheme struct {
	Models []Model `json:"models"`
}

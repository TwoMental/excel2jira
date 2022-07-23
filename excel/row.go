package excel

type Row interface {
	Get(key string) string
}

type RowService struct {
	rowContent []string
	relation   ColRelation
}

func (c *RowService) Get(key string) string {
	if val, ok := c.relation[key]; ok {
		defaultValue := c.relation[key].Default
		if val.Id == -1 || val.Id == 0 {
			// not in Excel
			return defaultValue
		} else {
			// in Excel
			if (len(c.rowContent) >= val.Id) && (c.rowContent[val.Id-1] != "") {
				// in Excel but no value
				return c.rowContent[val.Id-1]
			} else {
				// in Excel and has value
				return defaultValue
			}
		}
	} else {
		// relation not found
		return ""
	}
}

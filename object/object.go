package object

func ObjectMap(key []any, elements []any) map[any]any {

	var m = make(map[any]any, len(key))

	for i, v := range key {
		m[v] = elements[i]
	}
	return m
}

func ObjectMapList(maps ...map[any]any) []map[any]any {

	listMaps := make([]map[any]any, 0)
	listMaps = append(listMaps, maps...)

	return listMaps
}

package slices

func InterfaceForeach(objs []interface{}, cb func(interface{})) {
	for _, obj := range objs {
		cb(obj)
	}
}

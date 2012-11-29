package cbl

type ObjectMgr struct {
	objects map[string]*Object

	del_objects []string
}

// Public
func NewObjectMgr() *ObjectMgr {
	return &ObjectMgr{objects: make(map[string]*Object)}
}

func (om *ObjectMgr) New(name string) *Object {
	if obj, ok := om.objects[name]; ok {
		return obj
	}

	obj := NewObject()
	om.objects[name] = obj
	return obj
}

func (om *ObjectMgr) Get(name string) *Object {
	return om.objects[name]
}

func (om *ObjectMgr) Del(name string) {
	om.del_objects = append(om.del_objects, name)
}

func (om *ObjectMgr) Purge() {
	if om.del_objects != nil {
		for _, name := range om.del_objects {
			delete(om.objects, name)
		}
		om.del_objects = nil
	}
}

// Private

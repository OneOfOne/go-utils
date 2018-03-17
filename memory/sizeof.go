package memory

import "reflect"

var (
	sliceSize  = uint64(reflect.TypeOf(reflect.SliceHeader{}).Size())
	stringSize = uint64(reflect.TypeOf(reflect.StringHeader{}).Size())
)

func isNativeType(k reflect.Kind) bool {
	switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		return true
	}
	return false
}

func sizeofInternal(val reflect.Value, fromStruct bool, depth int, refmap map[uintptr]bool) (sz uint64) {

	if !val.IsValid() {
		return 0
	}

	if val.CanAddr() {
		refmap[val.Addr().Pointer()] = true
	}

	typ := val.Type()

	if !fromStruct {
		sz = uint64(typ.Size())
	}

	switch val.Kind() {
	case reflect.Ptr:
		if val.IsNil() {
			break
		}
		if refmap[val.Pointer()] {
			break
		}

		fallthrough
	case reflect.Interface:
		sz += sizeofInternal(val.Elem(), false, depth, refmap)

	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			sz += sizeofInternal(val.Field(i), true, depth, refmap)
		}

	case reflect.Array:
		if isNativeType(typ.Elem().Kind()) {
			break
		}
		sz = 0
		for i := 0; i < val.Len(); i++ {
			sz += sizeofInternal(val.Index(i), false, depth, refmap)
		}
	case reflect.Slice:
		if !fromStruct {
			sz = sliceSize
		}
		el := typ.Elem()
		if isNativeType(el.Kind()) {
			sz += uint64(val.Len()) * uint64(el.Size())
			break
		}
		for i := 0; i < val.Len(); i++ {
			sz += sizeofInternal(val.Index(i), false, depth, refmap)
		}
	case reflect.Map:
		if val.IsNil() {
			break
		}
		kel, vel := typ.Key(), typ.Elem()
		if isNativeType(kel.Kind()) && isNativeType(vel.Kind()) {
			sz += uint64(kel.Size()+vel.Size()) * uint64(val.Len())
			break
		}
		keys := val.MapKeys()
		for i := 0; i < len(keys); i++ {
			sz += sizeofInternal(keys[i], false, depth, refmap) + sizeofInternal(val.MapIndex(keys[i]), false, depth, refmap)
		}
	case reflect.String:
		if !fromStruct {
			sz = stringSize
		}
		sz += uint64(val.Len())
	}
	return
}

// Sizeof returns the estimated memory usage of object(s) not just the size of the type.
// On 64bit Sizeof("test") == 12 (8 = sizeof(StringHeader) + 4 bytes).
func Sizeof(objs ...interface{}) (sz uint64) {
	refmap := make(map[uintptr]bool)
	for i := range objs {
		sz += sizeofInternal(reflect.ValueOf(objs[i]), false, 0, refmap)
	}
	return
}

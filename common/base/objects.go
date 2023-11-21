package base

import (
	"github.com/cutexingluo/go-spring/common/reflect_util"
	"reflect"
)

// RequireNonNil require non nil
func RequireNonNil(v any) {
	if v == nil {
		panic(&ErrNilPointer{ErrMsg: "the arg is nil"})
	}
}

// Equals return true if the object is equal to the other, including  EqualsAddress and reflect.DeepEqual.
// Avoid using it in the Object implementation interface to prevent recursive calls
func Equals(v1 any, v2 any) bool {
	// Object equal
	if o1, ok := v1.(Object); ok {
		return o1.Equals(v2)
	} else if o2, ok := v2.(Object); ok {
		return o2.Equals(v1)
	}
	// equals
	return EqualsObject(v1, v2)
}

// EqualsObject return true if the object is equal to the other, including EqualsAddress and reflect.DeepEqual
func EqualsObject(v1 any, v2 any) bool {
	if EqualsAddress(v1, v2) {
		return true
	}
	return reflect.DeepEqual(v1, v2)
}

// EqualsAddress return true if the object's address is equal to the other's address
// you should use the pointer, such as EqualsAddress(&v1,&v2) if v1 and v2 are Struct
func EqualsAddress(v1 any, v2 any) bool {
	if v1 == nil || v2 == nil {
		return v1 == v2
	}
	// type equals
	if reflect.TypeOf(v1) != reflect.TypeOf(v2) {
		return false
	}
	return EqualsByPtrAddress(v1, v2) //self equals
}

// EqualsByAddress return true if the object's address is equal to the other's address
// It panics if v is nil and v's Kind is not Chan, Func, Map, Pointer, Slice, or UnsafePointer.
// structs and pointers are strictly different. like DeepEquals
//
// e.g.
//
// a := make(map[string]string);
// b := a   --> EqualsByAddress(a, b) == true, EqualsByAddress(a, a) == true, EqualsByAddress(&a, &a) == true,
// EqualsByAddress(a, &b) == false,  EqualsByAddress(&a, &b) == false,
func EqualsByAddress(v1 any, v2 any) bool {
	ptr1 := reflect_util.ValueOf(v1).Pointer()
	ptr2 := reflect_util.ValueOf(v2).Pointer()
	return ptr1 == ptr2
}

// EqualsByPtrAddress return true if the object's address is equal to the other's address
// It panics if v is nil and v's Kind is not Chan, Func, Map, Pointer, Slice, or UnsafePointer.
// It must be EqualsByPtrAddress(&self, &self) == true
//
// e.g.
//
// a := make(map[string]string);
// b := a   --> EqualsByPtrAddress(a, b) == false, EqualsByPtrAddress(a, a) == false, EqualsByPtrAddress(&a, &a) == true
func EqualsByPtrAddress(v1 any, v2 any) bool {
	ptr1 := reflect_util.GetPtr(v1).Pointer()
	ptr2 := reflect_util.GetPtr(v2).Pointer()
	return ptr1 == ptr2
}

// EqualsByStructAddress return true if the object's address is equal to the other's address
// It panics if v is nil and v's Kind is not Chan, Func, Map, Pointer, Slice, or UnsafePointer.
// false if v2 is not the v1 object and its reference
//
// e.g.
//
// a := make(map[string]string);
// c := make(map[string]string);
// b := a   --> EqualsByStructAddress(a, b) == true, EqualsByStructAddress(a, a) == true, EqualsByStructAddress(&a, &a) == true,
// EqualsByStructAddress(&a, b) == true,
// EqualsByStructAddress(a, c) == false,
func EqualsByStructAddress(v1 any, v2 any) bool {
	ptr1 := reflect_util.GetStruct(v1).Pointer()
	ptr2 := reflect_util.GetStruct(v2).Pointer()
	return ptr1 == ptr2
}

func ArrayCopy[T any](src []T, srcPos int, dest []T, destPos int, length int) {
	for i := 0; i < length; i++ {
		dest[destPos+i] = src[srcPos+i]
	}
}

// Reverse returns the reverse of the given slice.
func Reverse[T any](src []T) {
	for i := 0; i < len(src)/2; i++ {
		src[i], src[len(src)-i-1] = src[len(src)-i-1], src[i]
	}
}

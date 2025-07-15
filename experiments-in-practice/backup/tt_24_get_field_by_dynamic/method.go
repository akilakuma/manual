package main

import "reflect"

func getField(v *Vertex, field string) int {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return int(f.Int())
}

func setField(v *Vertex, field string) int {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	f.SetInt(100)
	return int(f.Int())
}

func setSTField(b *BigStruct, field string) reflect.Value {
	r := reflect.ValueOf(b)
	f := reflect.Indirect(r).FieldByName(field)
	return f
}


func setSTInternalField(b *ComplexStruct, field string) reflect.Value {
	r := reflect.ValueOf(&b.Res.Data)
	f := reflect.Indirect(r).FieldByName(field)
	return f
}

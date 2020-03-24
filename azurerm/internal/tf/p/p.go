package p

func Bool(input bool) *bool {
	return &input
}

func BoolI(i interface{}) *bool {
	b := i.(bool)
	return &b
}

func Int(input int) *int {
	return &input
}

func IntI(input interface{}) *int {
	i := input.(int)
	return &i
}

func Int32(input int32) *int32 {
	return &input
}

func Int32I(i interface{}) *int32 {
	i32 := i.(int32)
	return &i32
}

func Int64(input int64) *int64 {
	return &input
}

func Int64I(i interface{}) *int64 {
	i32 := i.(int64)
	return &i32
}

func Float64(input float64) *float64 {
	return &input
}

func Float64I(input interface{}) *float64 {
	f := input.(float64)
	return &f
}

func String(input string) *string {
	return &input
}

func StringI(i interface{}) *string {
	s := i.(string)
	return &s
}

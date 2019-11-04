package unmodifiableSlice


//暂时未开发
import (
	"errors"
	"learn_compiler/my_log"
)

type UnmodifiableSlice interface {

	append(interface{},interface{})interface{}

}

type MyUnmodifiableSlice struct {
	slice interface{}
}

func Convert (src []interface{}) []interface{}{

	return src
}

func (us *MyUnmodifiableSlice) append (src interface{},ele interface{}) interface{}{
	my_log.LogFatal(errors.New("can't modify this data structure"))
	return src
}






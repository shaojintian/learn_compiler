package my_log

//封装log.Fatalf() ,log.Printf(),显示日期，时间，错误行号，错误信息

import (
	"log"
	"runtime"
)



func LogFatal(err error){
	pc,_,line,ok := runtime.Caller(1)//1 表示显示LogFatal被调用的地方的信息
	if ok{
		funcName := runtime.FuncForPC(pc).Name()
		log.Fatalf("funcName:%s,err_line:%d,err:%v",funcName,line,err)
	}else{
			log.Fatalln(err)
	}

}

func LogPrint(err error){
	pc,_,line,ok := runtime.Caller(1)
	if ok{
		funcName := runtime.FuncForPC(pc).Name()
		log.Printf("funcName:%s,err_line:%d,err:%v",funcName,line,err)
	}else{
		log.Println(err)
	}
}



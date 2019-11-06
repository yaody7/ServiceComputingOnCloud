package main

import (
	"github.com/urfave/negroni"	//use negroni
	"net/http"				//use net/http
	"fmt"
	"io"
	"strconv"				//use to transform int to string/ string to int 
//	"strings"
)
func main() {
	n := negroni.New()			//get a negroni instance
	n.UseFunc(printFunctionInfo)	//push printFunctionInfo into MiddleWare chain
	n.UseFunc(errorFilter)		//push errorFilter into MiddleWare chain
	n.UseFunc(myCalculate)		//push myCalculate into MiddleWare chain

	router:=http.NewServeMux()		//get a router
	router.Handle("/",handler())	//route "\" to handler()

	n.UseHandler(router)		//push router into MiddleWare chain

	n.Run(":1234")			//set port: 1234

}


func printFunctionInfo(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc){			//print a "calculator" shape
	io.WriteString(rw,"\n");
	io.WriteString(rw,"-----------------------------------------------------------------------------------\n")
	io.WriteString(rw," ██████╗ █████╗ ██╗      ██████╗██╗   ██╗██╗      █████╗ ████████╗ ██████╗ ██████╗\n")
        io.WriteString(rw,"██╔════╝██╔══██╗██║     ██╔════╝██║   ██║██║     ██╔══██╗╚══██╔══╝██╔═══██╗██╔══██╗\n")
        io.WriteString(rw,"██║     ███████║██║     ██║     ██║   ██║██║     ███████║   ██║   ██║   ██║██████╔╝\n")
        io.WriteString(rw,"██║     ██╔══██║██║     ██║     ██║   ██║██║     ██╔══██║   ██║   ██║   ██║██╔══██╗\n")
        io.WriteString(rw,"╚██████╗██║  ██║███████╗╚██████╗╚██████╔╝███████╗██║  ██║   ██║   ╚██████╔╝██║  ██║\n")
        io.WriteString(rw," ╚═════╝╚═╝  ╚═╝╚══════╝ ╚═════╝ ╚═════╝ ╚══════╝╚═╝  ╚═╝   ╚═╝    ╚═════╝ ╚═╝  ╚═╝\n")
	io.WriteString(rw,"-----------------------------------------------------------------------------------\n")		//print the usage of this function
	io.WriteString(rw,"How to use? You can choose add, sub, mul, div\n");
	io.WriteString(rw,"http:127.0.0.1:1234/?operation=add&left=1&right=2\n")
	io.WriteString(rw,"You will get the answer:3\n")
	io.WriteString(rw,"-----------------------------------------------------------------------------------\n")
	next(rw,r)
}


func errorFilter(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc){
	r.ParseForm()					//analyze parameter
	my_operator:=r.Form["operation"]		//get parameter of "operation"
	left:=r.Form["left"]			//get parameter of "left"
	right:=r.Form["right"]			//get parameter of "right"
	fmt.Println(my_operator)			//print operator
	fmt.Println(left)				//print left operand
	fmt.Println(right)				//print right operand
	_,err1:=strconv.Atoi(left[0])		//left operand's string to int 
        _,err2:=strconv.Atoi(right[0])		//right operand's string to int

	if my_operator[0]!="add"&&my_operator[0]!="sub"&&my_operator[0]!="mul"&&my_operator[0]!="div"{		//detect whether if an error in operator
		io.WriteString(rw,"Please choose correct operation\n")
	} else if err1!=nil || err2!=nil {	//detect whether if an error in operand
		io.WriteString(rw,"Please choose number to calculate\n")
	} else{
		next(rw,r)			//if all correct, get next middleWare
	}
}

func myCalculate(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc){
	r.ParseForm()					//analyze parameters	
        my_operator:=r.Form["operation"]	//get parameter of "operation"
        left:=r.Form["left"]			//get parameter of "left"
        right:=r.Form["right"]			//get parameter of "right"
        a,_:=strconv.Atoi(left[0])		//left operand's string to int
        b,_:=strconv.Atoi(right[0])		//right operand's string to int
	answer:=0		
	true_operator:=" ";
	io.WriteString(rw,"Your answer is :\n")	//output hint
	if my_operator[0]=="add"{			//set operation base on "operation"
		answer=a+b
		true_operator="+"
	}else if my_operator[0]=="sub"{
		answer=a-b
		true_operator="-"
	}else if my_operator[0]=="div"{
		answer=a/b
		true_operator="/"

	}else{
		answer=a*b
		true_operator="*"
	}
	io.WriteString(rw,strconv.Itoa(a))   //print a format like "a + b = c"
	io.WriteString(rw," ")
	io.WriteString(rw,true_operator)
	io.WriteString(rw," ")
	io.WriteString(rw,strconv.Itoa(b))
	io.WriteString(rw," = ")
	io.WriteString(rw,strconv.Itoa(answer))
	io.WriteString(rw,"\n")

}

func handler() http.Handler{			//http.Handler middleware
	return http.HandlerFunc(myHandler)
}

func myHandler(rw http.ResponseWriter, r *http.Request) {	//set the output style
	rw.Header().Set("Content-Type", "text/plain")

}































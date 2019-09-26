package main

import (
	"bufio"
	"io"
	flag "github.com/spf13/pflag"
	"fmt"
	"os"
	"strings"
)

type selpg_Args struct{		//Argument needed by selpg
	start_page int
	end_page int
	lines_number int
	f bool
	help bool
	input string
	output string
}

var sa selpg_Args		//create an instance of selpg_Args

func init() {		//set the flag
	flag.IntVar(&sa.start_page,"s",1,"start page")
	flag.IntVar(&sa.end_page,"e",0,"end page")
	flag.IntVar(&sa.lines_number,"l",-1,"the number of lines per page")
	flag.BoolVar(&sa.f,"f",false,"default number of lines")
	flag.BoolVar(&sa.help,"h",false,"help")
	flag.StringVar(&sa.output,"d","","output")
}

func main(){
	flag.Parse()		//do the init work for flag
	if sa.help{		// -h situtation
		flag.PrintDefaults()
		return
	}
	if flag.NArg()==1{	//get the input file
		sa.input=flag.Arg(0)
	}
	check_error()
	my_work()
}

func check_error(){		//check whether if the argument are right
	check:=true
	if sa.start_page>sa.end_page {	//start page > end page?
		check= false
	}
	if sa.start_page<1 {		//start page < 1 ?
		check= false
	}
	if sa.f && sa.lines_number!=-1{	//-f && -l ?
		check= false
	}	
	if flag.NArg()!=0 && flag.NArg()!=1{	//# of no-flag arguments bigger than 1 ?
		check=false
	}
	if check==false{			//if false, exit
		fmt.Println("Args Error!")
		os.Exit(1)
	}
}

func my_work(){
	my_input:=os.Stdin		//init the argument needed by the work
	my_output:=os.Stdout
	current_page:=1
	current_line:=0
	var my_error error
	//set input
	if sa.input!=""{
		my_input,my_error=os.Open(sa.input)
		if my_error!=nil{
			fmt.Println(my_error)
			fmt.Println("Cannot open the input file")
			os.Exit(1)
		}
		defer my_input.Close()
	}

	//set output
	if sa.output!=""{			//ouput file
		my_output,my_error=os.OpenFile(sa.output,os.O_WRONLY,os.ModeAppend)
		if my_error!=nil{
			fmt.Println(my_error)
			fmt.Println("cannot open the output file")
			os.Exit(1)
		}
		defer my_output.Close()
	}
	if sa.f {		// -f situation
		reader:=bufio.NewReader(my_input)
		for true{

			if current_page>sa.end_page{
				fmt.Println("Success!")
				os.Exit(0)
			}
			buffer,my_error:=reader.ReadString('\f')
			if my_error==io.EOF{		// the end of the file
				if current_page>=sa.start_page&&current_page<=sa.end_page{
					fmt.Fprintf(my_output,"%s",buffer)
				}
				fmt.Println("Success!")
				os.Exit(0)
			}
			if my_error!=nil{		// get wrong
				fmt.Println(my_error)
				fmt.Println("Read error")
				os.Exit(1)
			}
			if current_page>=sa.start_page{
				fmt.Fprintf(my_output,"%s",buffer)	//write the file
			}
			current_page++
		}
	} else{			// -l situation
		reader:=bufio.NewReader(my_input)
		for true{
			if current_page>sa.end_page{
				fmt.Println("Success!")
				os.Exit(0)
			}
			buffer,my_error:=reader.ReadString('\n')
			if my_error==io.EOF{		// the end of the file
				if current_page>=sa.start_page&&current_page<=sa.end_page{
					buffer=strings.Replace(buffer,"\f","",-1)		//remove the "\f" in buffer
					fmt.Fprintf(my_output,"%s",buffer)
				}
				fmt.Println("Success!")
				os.Exit(0)
			}
			if my_error!=nil{		//get wrong
				fmt.Println(my_error)
				fmt.Println("Read error")
				os.Exit(1)
			}
			before:=len(buffer)
			buffer=strings.Replace(buffer,"\f","",-1)		//remove the "\f" in buffer
			after:=len(buffer)
			if before!=after{
				current_page++
			}
			if current_page>=sa.start_page{
				current_line++
				fmt.Fprintf(my_output,"%s",buffer)	//write the file
			}
			if current_line==sa.lines_number{
				current_line=0
				fenye:="\f"
				fmt.Fprintf(my_output,"%s",fenye)
			}
		}
	}
}

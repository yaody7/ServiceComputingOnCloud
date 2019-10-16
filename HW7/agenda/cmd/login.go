/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)



// registerCmd represents the register command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login",
	Long:  `login to use the agenda`,
	Run: func(cmd *cobra.Command, args []string) {
//	init()
	username, _ := cmd.Flags().GetString("username")
	password, _ := cmd.Flags().GetString("password")
		readFile()
		for _,value:=range user_info{
			if value.Username==username{
				if value.Password==password{
					fmt.Println("Login succeed!")
					f,_:=os.OpenFile("currentUser.txt",os.O_RDWR|os.O_CREATE|os.O_TRUNC,0644)
					defer f.Close()
					data:=[]byte(username)
					f.Write(data)
					return
				}else{
					fmt.Println("Wrong password!")
					return
				}
			}
		}

		fmt.Println("You have not registered!")
		
	},
}


func init() {
	fileptr,e :=  os.Open("entity/User.json")
	defer fileptr.Close()
	if e!=nil&&fileptr!=nil{
		fmt.Println(e)
	}
	if fileptr==nil{
		fileptr,e:=os.Create("entity/User.json")
		defer fileptr.Close()
		if e!=nil{
			fmt.Println(e)
		}
	}
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("username","","","")
	loginCmd.Flags().StringP("password","","","")


}

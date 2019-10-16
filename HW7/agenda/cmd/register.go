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
	"encoding/json"
	"github.com/spf13/cobra"
	"os"
)

type User struct{
	Username string
	Password string
	Email string
	Phone string
}

var user_info []User

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Sign up",
	Long:  `Sign up your account, so that you can use the agenda`,
	Run: func(cmd *cobra.Command, args []string) {
//	init()
	username, _ := cmd.Flags().GetString("username")
	password, _ := cmd.Flags().GetString("password")
        email, _ := cmd.Flags().GetString("email")
        phone, _ := cmd.Flags().GetString("phone")
		readFile()
		for _,value:=range user_info{
			if value.Username==username{
				fmt.Println("Fail: This Username is used!")
				return
			}
		}
		fileptr, err := os.OpenFile("entity/User.json", os.O_APPEND|os.O_WRONLY|os.O_TRUNC,0644)
		defer fileptr.Close()
		encoder := json.NewEncoder(fileptr)
		var temp User
		temp.Username=username
		temp.Password=password
		temp.Email=email
		temp.Phone=phone
		user_info=append(user_info,temp)
		err=encoder.Encode(user_info)
		if err!=nil{
			fmt.Println(err)
		}
		fmt.Println("Succeed!")	
	},
}

func readFile(){
	fileptr, _:=os.Open("entity/User.json")
	defer fileptr.Close()
	decoder := json.NewDecoder(fileptr)
	decoder.Decode(&user_info)
//	fmt.Println(user_info)
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
	rootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("username","","","")
	registerCmd.Flags().StringP("password","","","")
	registerCmd.Flags().StringP("email","","","")
	registerCmd.Flags().StringP("phone","","","")

}

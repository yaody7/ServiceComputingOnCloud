# HW7 CLI命令行实用程序开发实战 - Agenda

### 1. 实验准备

> **配置环境：**
>
> 打开 **$GOPATH** ，使用 **git clone** 将 **sys** 和 **text** 两个包下载下来。
> 
> 接着使用 **go install github.com/spf13/cobra/cobra** 安装 **cobra**
> 
> 接着要做初始化工作
> 
> 本次作业是做 **agenda** ，所以使用命令 **cobra init agenda --pkg-name=agenda/cmd** 做初始化。
> 
>这样我们的项目就初始化完毕了，可以开始代码设计了。



### 2.设计过程

> 本次实验，我选择实现的命令式 **login**、**logout** 以及 **register** ，接下来将分别介绍其实现过程。
>
> ### register
> 
> > **register** 比较复杂，我们首先要读取用户信息，接着遍历所有用户的名字，检查是否与当前的用户信息冲突。若不冲突则可以注册，否则则返回失败。
> >
> > 我们读取用户信息使用到的是 **json** 的框架，我们通过 **json.NewDecoder** 来绑定读取文件，并使用 **Decode** 函数来将用户信息传入我们先前定义好的数组中。接着我们就可以在这个数组中遍历所有的用户名，来查询冲突。
> >
> > **具体代码如下：**
> >
> > ```go
> > type User struct{
> > 	Username string
> > 	Password string
>> 	Email string
> > 	Phone string
> > }
> > 
> > var user_info []User
> > 
> > // registerCmd represents the register command
> > var registerCmd = &cobra.Command{
> > 	Use:   "register",
> > 	Short: "Sign up",
> > 	Long:  `Sign up your account, so that you can use the agenda`,
> > 	Run: func(cmd *cobra.Command, args []string) {
> > //	init()
> > 	username, _ := cmd.Flags().GetString("username")
> > 	password, _ := cmd.Flags().GetString("password")
> >         email, _ := cmd.Flags().GetString("email")
> >         phone, _ := cmd.Flags().GetString("phone")
>> 		readFile()
> > 		for _,value:=range user_info{
> > 			if value.Username==username{
> > 				fmt.Println("Fail: This Username is used!")
> > 				return
> > 			}
> > 		}
> > 		fileptr, err := os.OpenFile("entity/User.json", os.O_APPEND|os.O_WRONLY|os.O_TRUNC,0644)
> > 		defer fileptr.Close()
> > 		encoder := json.NewEncoder(fileptr)
> > 		var temp User
> > 		temp.Username=username
> > 		temp.Password=password
> > 		temp.Email=email
> > 		temp.Phone=phone
> > 		user_info=append(user_info,temp)
> > 		err=encoder.Encode(user_info)
> > 		if err!=nil{
> > 			fmt.Println(err)
> > 		}
> > 		fmt.Println("Succeed!")	
> > 	},
> > }
> > 
>> func readFile(){
> > 	fileptr, _:=os.Open("entity/User.json")
> > 	defer fileptr.Close()
> > 	decoder := json.NewDecoder(fileptr)
> > 	decoder.Decode(&user_info)
> > //	fmt.Println(user_info)
> > }
> > ```
> 
> 
> 
> ### login
> 
> > login就比较简单了，我们可以调用之前已经写好的 **readFile** 函数，然后查询要登陆的用户是否在数组中，若存在则对比密码，若不存在则直接返回失败。登录成功之后需要更新 **currentUser.txt** 的信息。
> >
> > **具体代码如下：**
> >
> > ```go
> > var loginCmd = &cobra.Command{
> > 	Use:   "login",
> > 	Short: "login",
> > 	Long:  `login to use the agenda`,
> > 	Run: func(cmd *cobra.Command, args []string) {
> > //	init()
> > 	username, _ := cmd.Flags().GetString("username")
> > 	password, _ := cmd.Flags().GetString("password")
> > 		readFile()
> > 		for _,value:=range user_info{
> > 			if value.Username==username{
> > 				if value.Password==password{
> > 					fmt.Println("Login succeed!")
> > 					return
> > 				}else{
> > 					fmt.Println("Wrong password!")
> > 					return
> > 				}
> > 			}
> > 		}
> > 		f,_:=os.OpenFile("currentUser.txt",os.O_RDWR|os.O_CREATE|os.O_TRUNC,0644)
> > 		defer f.Close()
> > 		data:=[]byte(username)
> > 		f.Write(data)
> > 		fmt.Println("You have not registered!")
> > 		
> > 	},
> > }
> > ```
> 
> 
> 
> ### logout
> 
> logout是最简单的一部分，我们只需更改 **currentUser.txt** 里的信息为 **no one** 即可。
> 
> **具体代码如下：**
> 
> ```go
> var logoutCmd = &cobra.Command{
> 	Use:   "logout",
> 	Short: "logout",
> 	Long:  `logout the agenda`,
> 	Run: func(cmd *cobra.Command, args []string) {
> //	init()
> 
> 	f,_:=os.OpenFile("currentUser.txt",os.O_RDWR|os.O_CREATE|os.O_TRUNC,0644)
> 
> 
> 	defer f.Close()
> 	data:=[]byte("no one")
> 	f.Write(data)
> 	fmt.Println("No one login now!")
> 		
> 	},
> }
> 
> ```



### 3. 测试结果

> ### register
>
> ![1571243061842](C:\Users\89481\AppData\Roaming\Typora\typora-user-images\1571243061842.png)
>
> ![1571243126337](C:\Users\89481\AppData\Roaming\Typora\typora-user-images\1571243126337.png)
>
> ### login
>
> ![1571243836942](C:\Users\89481\AppData\Roaming\Typora\typora-user-images\1571243836942.png)
>
> ![1571244373127](C:\Users\89481\AppData\Roaming\Typora\typora-user-images\1571244373127.png)
>
> 
>
> ### logout
>
> ![1571244490296](C:\Users\89481\AppData\Roaming\Typora\typora-user-images\1571244490296.png)
>
> ![1571244498701](C:\Users\89481\AppData\Roaming\Typora\typora-user-images\1571244498701.png)




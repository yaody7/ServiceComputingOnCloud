# 基于Negroni框架的cloudgo应用

本次实验是基于Negroni框架的应用。我设计了一个简单的四则运算应用。

这个应用设计主要分为两部分， **中间件设计** 以及 **main函数的设计** 。接下来就分别对这两个部分进行介绍。

### 中间件设计

> **printFunctionInfo**
>
> > 这个中间件的功能是 **打印应用使用方法** ，以及在开头输出一个 **”calculator“** 的大型字样。
> >
> > 使用到的是 **io.WriteString** 来写 **response** 
> >
> > **使用方法：**
> >
> > 访问：http:127.0.0.1:1234/?operation=add&left=1&right=2
> >
> > 其中 **opertaion** 可以是 **mul、div、add、sub**
> >
> > 服务端将返回计算结果。
>
>  
>
> **errorFilter**
>
> > 这个中间件的功能是 **检查输入参数错误** 
> >
> > **operation** 的检查十分简单，只需要判断是否是 **mul、div、add、sub** 中的一个即可
> >
> > **left、right** 的检测用到了 **strconv库** 中的 **Atoi** 函数，若是转换后 **error** 为空，那么就证明输入参数无错。
> >
> > 若是这两项检测有错误，则打印错误信息，并停止服务
> >
> > 若都没有错误，则进入下一个中间件
>
>  
>
> **myCalculate**
>
> > 这个中间件的功能就是 **计算** 了
> >
> > 通过 **operation** 来确定我们所作运算的种类，接着设置一个变量 **answer** 来存储答案。
> >
> > 在输出的时候，同样要使用到 **io.WriteString** 函数。
> >
> > 这里注意到的是，该函数的第二个参数是一个 **string** 类型。所以我们要使用到 **strconv** 库中的 **Itoa** 函数进行 **int** 到 **string** 的转换。
>
>  
>
> 还有另外的 **http.handler** 中间件，是照搬老师的课程博客的，这里就不再赘述了。



### main函数设计

> main函数设计主要就是要将设计好的中间件，使用 **UseFunc** 函数加入到中间件链中即可。其余的就和课程博客一样，设置路由，并且设置端口而已。
>
> 这里使用的端口是 **1234**
>
> ```go
> func main() {
> 	n := negroni.New()			//get a negroni instance
> 	n.UseFunc(printFunctionInfo)	//push printFunctionInfo into MiddleWare chain
> 	n.UseFunc(errorFilter)		//push errorFilter into MiddleWare chain
> 	n.UseFunc(myCalculate)		//push myCalculate into MiddleWare chain
> 
> 	router:=http.NewServeMux()		//get a router
> 	router.Handle("/",handler())	//route "\" to handler()
> 
> 	n.UseHandler(router)		//push router into MiddleWare chain
> 
> 	n.Run(":1234")			//set port: 1234
> 
> }
> ```



### 效果展示

![1573006706078](C:\Users\89481\AppData\Roaming\Typora\typora-user-images\1573006706078.png)




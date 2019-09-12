# HW3 安装go语言开发环境

### 1. 安装VSCode

> 我们可以到VSCode官网进行下载
>
>  https://code.visualstudio.com/ 
>
> ![1568027205435](https://github.com/yaody7/ServiceComputingOnCloud/blob/master/HW3/pics/1568027205435.png)
>
> 在这个位置，我们应该使用Linux x64选项下的.deb进行下载。
>
> 
>
> 接着只需要点击下载的.deb文件进行安装即可。



### 2.安装 golang

> ### 方法1. 命令行安装（Ubuntu 18.04）
>
> 1. 安装
>
> > `sudo apt-get install golang`
> >
> > `sudo apt-get install golang-go`
>
> 2. 验证安装成功与否：
>
> > `go version`
>
> 3. 创建工作空间
>
> > `mkdir $HOME/gowork`
>
> 4. 配置环境变量
>
> > `vim ~/.profile`
> >
> > 添加以下内容：
> >
> > `export GOPATH=$HOME/gowork`
> >
> > `export PATH=$PATH:$GOPATH/bin`
> >
> > 执行配置：
> >
> > `source $HOME/.profile`
>
> ### 方法2.（搜索博客获得，据说可以获得更新版本）
>
> https://blog.csdn.net/runner668/article/details/80958496
>
> > 1. 从课程网站中给出的链接下载压缩包
> > 2. 打开终端，并获取root权限
> >
> > ![1568030135182](https://github.com/yaody7/ServiceComputingOnCloud/blob/master/HW3/pics/1568030135182.png)
> >
> > 3. 将安装包解压到/usr/local目录下：
> >
> >    `	tar -xzf go.tar.gz -C/usr/local`
> >
> > 4. 配置环境变量
> >
> > >  `vim ~/.bashrc`
> > >
> > > 将以下内容插入
> > >
> > > `export GOPATH=/opt/go
> > > export GOROOT=/usr/local/go
> > > export GOARCH=386
> > > export GOOS=linux
> > > export GOBIN=$GOROOT/bin/
> > > export GOTOOLS=$GOROOT/pkg/tool/
> > > export PATH=$PATH:$GOBIN:$GOTOOLS`
> >
> > 5. 重新加载profile文件。
> >
> >    `source ~/.bashrc`



### 3. 编程与测试

> 1. 使用vim编写代码
>
>    ` vim hello.go`
>
>    代码如下：
>
>    `package main
>    import "fmt"
>    func main(){
>            fmt.Printf("hello, world\n")
>    }`
>
> 2. 测试
>
>    `go run hello.go`
>
>    ![1568030189031](https://github.com/yaody7/ServiceComputingOnCloud/blob/master/HW3/pics/1568030189031.png)



### 4. 安装gotour

> 在这个网站上下载所需文件

> https://bitbucket.org/mikespook/go-tour-zh/downloads/

> 使用git将下面两个仓库clone下来

> https://github.com/golang/tools
> https://github.com/golang/net

> 打开$GOPATH，并按照如下的结构安排刚才下载的文件

![在这里插入图片描述](https://img-blog.csdnimg.cn/20190912153354816.png)

> 打开go-tour-zh中的gotour并执行`go install`

> 完成后，gotour就会出现在bin中，就代表安装完成了！

### 5. 实验心得

​	本次实验中，我们为以后的go编程打下了基础——配置了go的运行环境。因为网络的原因，教学网站上的链接总是点不开。所以只好到网上找教程，发现就是要下载一个压缩包，而且这个压缩包可能在国外的服务器上。因为自己的PC机上没有科学上网的条件，我选择了在手机上下载这个压缩包再传回电脑进行解压。可是当我解压缩完成，环境变量也按教程配置好后，运行代码还是出现了bug。没有办法只好使用命令行进行安装了，之后的过程也就顺风顺水，第一个helloworld也顺利跑出来了。

​	我认为还需要努力的地方就是环境变量的配置问题。虽然把环境配置好了，但是还是对教学网站上那些奇奇怪怪的符号代码不是很了解。之后有机会还是要钻研一下，否则以后如果需要自己配置环境就没有办法解决了。





> 


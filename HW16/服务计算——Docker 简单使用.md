# Docker 简单使用

> Docker是一个应用容器引擎，使用者可以将其应用以及所有需要的依赖打包到一个包中，然后发布到机器上进行运作。接下里我们就一步步了解一些Docker的使用

### Docker安装

> - $ sudo apt-get install \
>
>   ​    apt-transport-https \
>
>   ​    ca-certificates \
>
>   ​    curl \
>
>   ​    software-properties-common
>
> - $ curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
>
> - $ sudo add-apt-repository \
>
>     "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
>
>     $(lsb_release -cs) \
>
>     stable"
>
> - $ sudo apt-get update
>   $  sudo apt-get install docker-ce

### Docker运行

> $ docker run -d -p 80:80 httpd
>
> 得到以下界面：
>
> ![在这里插入图片描述](https://img-blog.csdnimg.cn/20191216112645438.png)
>
> 接着直接在浏览器中访问我们机器本地ip：127.0.0.1
>
> ![在这里插入图片描述](https://img-blog.csdnimg.cn/20191216112704596.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3UwMTE0MzA5MzI=,size_16,color_FFFFFF,t_70)
>
> 这样就代表我们的Docker已经成功安装并运行了

### 运行第一个容器

> 我们使用docker自带的hello-world镜像来进行测试
>
> $ docker run hello-world
>
> 就会出现以下的界面：
>
> ![在这里插入图片描述](https://img-blog.csdnimg.cn/20191216112722702.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3UwMTE0MzA5MzI=,size_16,color_FFFFFF,t_70)
>
> $ docker run -it ubuntu bash
>
> ![在这里插入图片描述](https://img-blog.csdnimg.cn/20191216112827208.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3UwMTE0MzA5MzI=,size_16,color_FFFFFF,t_70)

### Docker基本操作

> **显示本地镜像库内容**
>
> $ docker images
>
> ![在这里插入图片描述](https://img-blog.csdnimg.cn/20191216112842408.png)
>
> **显示所有容器**
>
> $ docker ps -a
>
> ![在这里插入图片描述](https://img-blog.csdnimg.cn/20191216112855608.png)
>
> 其中还有比较常用的指令有
>
> - container
> - image
> - network
> - volume
>
> 我们可以通过 -help 来获取其具体参数使用，这里就不一一赘述了

### Docker & MySQL

> 我们使用Docker的时候，很多时候也会需要使用到数据库，接下里就是在Docker上使用MySQL的一些操作：
>
> 首先需要拉取MySQL的镜像：
>
> ![在这里插入图片描述](https://img-blog.csdnimg.cn/20191216112912947.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3UwMTE0MzA5MzI=,size_16,color_FFFFFF,t_70)
>
> 接着打开运行MySQL：
>
> ![在这里插入图片描述](https://img-blog.csdnimg.cn/20191216113014992.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3UwMTE0MzA5MzI=,size_16,color_FFFFFF,t_70)
>
> 这样就可以操作我们的MySQL数据库了
>
> MySQL的具体使用就应该查看MySQL的文档了。
>
> 而我们的数据库文件就会保存在**/var/lib/mysql** 中
>
> 这个卷也会在主机的 **/var/lib/docker/volumes** 中

### 监控容器

可以使用：

> docker info
>
> docker info -format {{.ServerVersion}}

来监测docker的状态

可以使用：

> docker stats

来监测各个容器的信息

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191216112926264.png)

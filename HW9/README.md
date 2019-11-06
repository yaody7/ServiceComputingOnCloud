# 基于Negroni框架的cloudgo应用

Negroni是一个很好用的框架，其中间件框架十分清晰。我们只需要设计好我们自己想要的中间件，并将其加入到中间链中，Negroni就会帮我们将各个中间件串起来依次执行了。而且Negroni还兼容原生的http.Handler，方便我们的使用。老师的教学博客中，同样应用了这个框架，所以本次实验我选择使用Negroni框架。



## curl测试

![1573006706078](C:\Users\89481\AppData\Roaming\Typora\typora-user-images\1573006706078.png)



## ab测试

17343141$>ab -n 1000 -c 100 127.0.0.1:1234/?operation=add\&left=1\&right=6

This is ApacheBench, Version 2.3 <$Revision: 1807734 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:        
Server Hostname:        127.0.0.1
Server Port:            1234

Document Path:          /?operation=add&left=1&right=6
Document Length:        1643 bytes

Concurrency Level:      100
Time taken for tests:   0.161 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      1762000 bytes
HTML transferred:       1643000 bytes
Requests per second:    6199.78 [#/sec] (mean)
Time per request:       16.130 [ms] (mean)
Time per request:       0.161 [ms] (mean, across all concurrent requests)
Transfer rate:          10667.98 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   1.0      0       4
Processing:     0   15  12.9     10      87
Waiting:        0   14  12.8      9      87
Total:          0   15  13.0     10      88

Percentage of the requests served within a certain time (ms)
  50%     10
  66%     16
  75%     21
  80%     25
  90%     35
  95%     40
  98%     51
  99%     56
 100%     88 (longest request)


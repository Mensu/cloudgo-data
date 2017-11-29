# cloudgo-data

> 课程《服务计算》作业六：用 Go 开发 web 应用程序 [cloudgo-data](http://blog.csdn.net/pmlpml/article/details/78602290)

## Install

```
go get github.com/Mensu/cloudgo-data
```

## 任务要求

- [x] 使用 ``xorm`` 或 ``gorm`` 实现本文的程序，从编程效率、程序结构、服务性能等角度对比 ``database/sql`` 与 ``orm`` 实现的异同！
  1. ``orm`` 是否就是实现了 ``dao`` 的自动化？
  2. [x] 使用 ``ab`` 测试性能

## ``database/sql`` VS ``orm``

本次作业使用 [``gorm``](http://jinzhu.me/gorm/) 实现文章中的程序

### 编程效率

如果是数据模型简单的项目，那么需要重复编写简单增删改查 ``SQL`` 以及 ``Scan`` 出每一行的值的 ``database/sql`` 会显得十分繁琐，编程效率低于 ``orm``。其实这种情况下用 No-SQL 会更合适

如果是数据模型复杂的项目，``orm`` 需要加一大堆结构标签来定制查询模型，而且需求稍微一变化，很多时候就难以找到正确的调用姿势来让 ``orm`` 执行期望的 SQL 语句，结果还是得用回 SQL 语句。相比之下，``database/sql`` 写起来会显得更加直观，易于维护，编程效率更高一些

### 程序结构

``database/sql`` 一般需要编写诸如 ``INSERT INTO`` 和 ``创建对象`` 这样，``SQL 语句`` 和 ``数据存取接口`` 之间的转换

``orm`` 通过提供 API 的方式完成了这种转换

### 服务性能

``orm`` 依赖于反射，会做很多额外的工作完成数据和结构之间的转换，性能上不如什么都不做的 ``database/sql``

## 1. ``orm`` 是否就是实现了 ``dao`` 的自动化？

``orm`` 只是将数据库中的实体映射到结构，只不过将**写 SQL 语句**换做是**函数调用**，相对来说进行的操作还比较底层，不能完全实现 ``dao`` 的自动化。但不可否认，如果数据模型较为简单，如本次作业，那 ``orm`` 确实在一定程度上确实是实现了 ``dao`` 的自动化

## 2. [x] 使用 ``ab`` 测试性能

### POST /service/userinfo

测试数据位于 ``testdata/post``

- ``database/sql``

```plain
$ ab -n 10000 -c 100 -p testdata/post -T "application/x-www-form-urlencoded" "http://127.0.0.1:8080/service/userinfo"
This is ApacheBench, Version 2.3 <$Revision: 1796539 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        127.0.0.1
Server Port:            8080

Document Path:          /service/userinfo
Document Length:        4262 bytes

Concurrency Level:      100
Time taken for tests:   6.003 seconds
Complete requests:      10000
Failed requests:        9983
   (Connect: 0, Receive: 0, Length: 9983, Exceptions: 0)
Non-2xx responses:      36
Total transferred:      2616777 bytes
Total body sent:        2050000
HTML transferred:       1377065 bytes
Requests per second:    1665.87 [#/sec] (mean)
Time per request:       60.029 [ms] (mean)
Time per request:       0.600 [ms] (mean, across all concurrent requests)
Transfer rate:          425.70 [Kbytes/sec] received
                        333.50 kb/s sent
                        759.20 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   0.6      0      11
Processing:     1   59  46.1     50     340
Waiting:        1   59  46.0     50     339
Total:          1   59  46.2     51     341
WARNING: The median and mean for the initial connection time are not within a normal deviation
        These results are probably not that reliable.

Percentage of the requests served within a certain time (ms)
  50%     51
  66%     77
  75%     88
  80%     95
  90%    115
  95%    140
  98%    181
  99%    211
 100%    341 (longest request)
```

- ``orm``

```plain
ab -n 10000 -c 100 -p testdata/post -T "application/x-www-form-urlencoded" "http://127.0.0.1:8080/service/userinfo?orm="
This is ApacheBench, Version 2.3 <$Revision: 1796539 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        127.0.0.1
Server Port:            8080

Document Path:          /service/userinfo?orm=
Document Length:        118 bytes

Concurrency Level:      100
Time taken for tests:   6.268 seconds
Complete requests:      10000
Failed requests:        9997
   (Connect: 0, Receive: 0, Length: 9997, Exceptions: 0)
Total transferred:      2468622 bytes
Total body sent:        2100000
HTML transferred:       1228622 bytes
Requests per second:    1595.48 [#/sec] (mean)
Time per request:       62.677 [ms] (mean)
Time per request:       0.627 [ms] (mean, across all concurrent requests)
Transfer rate:          384.63 [Kbytes/sec] received
                        327.20 kb/s sent
                        711.83 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.6      0      28
Processing:     0   61  51.8     49     387
Waiting:        0   61  51.8     49     387
Total:          1   62  51.8     50     388

Percentage of the requests served within a certain time (ms)
  50%     50
  66%     82
  75%     95
  80%    103
  90%    130
  95%    156
  98%    190
  99%    212
 100%    388 (longest request)
```

### GET /service/userinfo

- ``database/sql``

```plain
$ ab -n 10000 -c 100 "http://127.0.0.1:8080/service/userinfo"
This is ApacheBench, Version 2.3 <$Revision: 1796539 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        127.0.0.1
Server Port:            8080

Document Path:          /service/userinfo
Document Length:        4454 bytes

Concurrency Level:      100
Time taken for tests:   15.419 seconds
Complete requests:      10000
Failed requests:        9989
   (Connect: 0, Receive: 0, Length: 9989, Exceptions: 0)
Non-2xx responses:      136
Total transferred:      120021624 bytes
HTML transferred:       118989856 bytes
Requests per second:    648.54 [#/sec] (mean)
Time per request:       154.194 [ms] (mean)
Time per request:       1.542 [ms] (mean, across all concurrent requests)
Transfer rate:          7601.39 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   4.7      0     383
Processing:     1  148 152.8    111    1083
Waiting:        1  147 152.7    110    1082
Total:          2  148 153.0    112    1083

Percentage of the requests served within a certain time (ms)
  50%    112
  66%    169
  75%    208
  80%    237
  90%    332
  95%    475
  98%    604
  99%    683
 100%   1083 (longest request)
```

- ``orm``

```plain
$ ab -n 10000 -c 100 "http://127.0.0.1:8080/service/userinfo?orm="
This is ApacheBench, Version 2.3 <$Revision: 1796539 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        127.0.0.1
Server Port:            8080

Document Path:          /service/userinfo?orm=
Document Length:        12002 bytes

Concurrency Level:      100
Time taken for tests:   19.054 seconds
Complete requests:      10000
Failed requests:        13
   (Connect: 0, Receive: 0, Length: 13, Exceptions: 0)
Non-2xx responses:      13
Total transferred:      120949302 bytes
HTML transferred:       119919133 bytes
Requests per second:    524.81 [#/sec] (mean)
Time per request:       190.544 [ms] (mean)
Time per request:       1.905 [ms] (mean, across all concurrent requests)
Transfer rate:          6198.80 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.4      0       9
Processing:     3  189 105.2    185     666
Waiting:        3  188 105.1    184     665
Total:          3  189 105.3    185     666

Percentage of the requests served within a certain time (ms)
  50%    185
  66%    229
  75%    257
  80%    274
  90%    324
  95%    370
  98%    442
  99%    491
 100%    666 (longest request)
```

### 性能测试评价

``database/sql`` 和 ``orm`` 在处理请求的 web 服务上并没有对性能造成太大的影响，毕竟请求主要的时间是耗在 IO，而不是 go 语言的执行上。使用 ``nodejs`` 也是差不多的效果

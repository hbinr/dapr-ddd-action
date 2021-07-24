# go-chi、gin 路由性能测试：

## 测试命令
```sh
 wrk -t12 -c100 -d30s http://localhost:8090/user/1
```

### gin 表现
```sh
Running 30s test @ http://localhost:8090/user/1
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   141.89ms  237.67ms   1.99s    89.42%
    Req/Sec   116.87     49.60   353.00     67.44%
  41897 requests in 30.09s, 7.67MB read
  Socket errors: connect 0, read 0, write 0, timeout 29
Requests/sec:   1392.18
Transfer/sec:    261.03KB
```

timeout 数量有29

dapr 运行时，有不少额外的报错内容:
```sh
2021/07/23 23:46:36 failed to send the request: Post "http://localhost:9411/api/v2/spans": EOF
2021/07/23 23:46:36 failed to send the request: Post "http://localhost:9411/api/v2/spans": EOF
2021/07/23 23:46:37 failed to send the request: Post "http://localhost:9411/api/v2/spans": EOF
2021/07/23 23:46:38 failed to send the request: Post "http://localhost:9411/api/v2/spans": EOF
2021/07/23 23:46:39 failed to send the request: Post "http://localhost:9411/api/v2/spans": EOF
2021/07/23 23:46:40 failed to send the request: Post "http://localhost:9411/api/v2/spans": EOF
```

### go-chi 表现

```sh
Running 30s test @ http://localhost:8090/user/1
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   117.77ms  187.42ms   1.71s    88.62%
    Req/Sec   143.94     55.58   430.00     67.33%
  51680 requests in 30.08s, 7.49MB read
Requests/sec:   1717.92
Transfer/sec:    255.00KB
```
QPS更多，没有和gin一样的报错内容， 也没有timeout


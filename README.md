# 计时跟踪器

一个简单的计时跟踪器实现
```go
    tracker := NewTimeTracker("查询商品信息")
		tracker.StepStart("第一步查询商品信息")
		time.Sleep(time.Duration(int64(rand.Intn(100)*1000*10000)))
		tracker.StepEnd()
		tracker.StepStart("第二步查询商品关联的用户信息")
		time.Sleep(time.Duration(int64(rand.Intn(100)*1000*10000)))
		tracker.StepEnd()
		tracker.StepStart("第三步查询商品关联的产地信息")
			tracker.StepStart("第三步查询商品关联的产地信息之省份信息")
			time.Sleep(time.Duration(int64(rand.Intn(100)*1000*10000)))
			tracker.StepEnd()
			tracker.StepStart("第三步查询商品关联的产地信息之城市信息")
			time.Sleep(time.Duration(int64(rand.Intn(100)*1000*10000)))
			tracker.StepEnd()
		tracker.StepEnd()
	tracker.Close()
    tracker.PrintBeautiful()
```
结果如下:
```text
------ 查询商品信息 0h 0m 2s 748ms 
--------- 第一步查询商品信息 0h 0m 0s 810ms 
--------- 第二步查询商品关联的用户信息 0h 0m 0s 871ms 
--------- 第三步查询商品关联的产地信息 0h 0m 1s 66ms 
------------ 第三步查询商品关联的产地信息之省份信息 0h 0m 0s 472ms 
------------ 第三步查询商品关联的产地信息之城市信息 0h 0m 0s 593ms
```
####enjoy
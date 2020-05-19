
concurrent包含知识点：

`sample1` 并发最简单使用：
    goroutine + sync.WaitGroup
    
`sample2` 并发中互斥锁使用：
    `sample2` + sync.Mutex

`sample3` 并发中使用管道管理一组goroutine：
    `sample1` + channel
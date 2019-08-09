# select语句

## 概述

本文描述了select的用法，通过select可以监听多个channel的读写事件。这很类似于linux系统编程的select函数。但在Go中，实现的机制明显是不同的。linux系统编程的select是轮训的机制，而且监控的是文件描述符，且有数量的限制。Go中的select和channel配合使用，监控的是channel的读写状态。

## select的要点

1. select会阻塞在多个channel上，对多个channel的读/写事件进行监控。
2. select中对case语句的判断不是顺序进行的。
3. 在select中执行case语句时，不会自动的fall through。
4. 在select中所有channel的读和写都被认为是同时进行的。
5. case中的channel的事件包括：读取的时候，channel被close，或写入时channel没有空间。
6. 当所有channel都没有数据读取时，select阻塞，当其中有一个channel有数据时则进行处理。
7. 可以为select设置一个超时时间，当select超时时，可以完成一些其他工作。

## select的基本格式

    var c1, c2 <-chan interface{}
    var c3 chan<- interface{}
    select {
        case <- c1:         //监听  channel的读事件
        // Do something
        case <- c2:         //读事件
        // Do something
        case c3<- struct{}{}:   //监控  channel的写事件
        // Do something
    }

通过select来检测channel的关闭事件

    func TestSelect1() {
        start := time.Now()
        c := make(chan interface{})

        go func() {
            time.Sleep  (2*time.Second)
            close(c)
        }()

        fmt.Println("Blocking on    read...")
        select {
        case <-c:
            fmt.Printf("Unblocked   %v later.\n",     time.Since(start))
        }
    }

注意：当close channel时，读取channel的一方会从channel中读取到value,false，此时的value一般情况下为nil。
该例子也可以用来通知当不使用channel时，关闭channel的情况。

多个channel同时准备好读的情况
当多个channel同时准备好，select的行为是怎样的呢？我们通过一个例子来看一下：

func TestMultiChannel() {
    c1 := make(chan interface{}); close(c1)
    c2 := make(chan interface{}); close(c2)
    c3 := make(chan interface{}); close(c3)

    var c1Count, c2Count, c3Count int
    for i := 1000; i >= 0; i-- {
        select {
        case <-c1:
            c1Count++
        case <-c2:
            c2Count++
        case <-c3:
            c3Count++
        }
    }
    fmt.Printf("c1Count: %d\nc2Count: %d\nc3Count: %d\n", c1Count, c2Count, c3Count)
}

输出：

c1Count: 337
c2Count: 319
c3Count: 345

多运行几次，可以看出，几个数字相差都不是很大。
以上例子，同时有3个channel可读取，从以上的输出可以看出，select对多个channel的读取调度是基本公平的。让每一个channel的数据都有机会被处理。

没有任何channel准备好，处理超时
在很多情况下，当channel没有准备好时，我们希望能够设置一个超时时间，并在等待channel超时时进行一些处理。此时就可以按以下方式来进行编码：

    func TestProcTimeOut() {
        var c <-chan int

        for {
            select {
            case <-c:
            case <-time.After(1 * time.Second):
                fmt.Println("Timed out.Do something.")
            }
        }
    }

该代码会每隔1秒钟，打印出：

    Timed out.Do something.
    Timed out.Do something.
    Timed out.Do something.
    ...

这样select就变成了“非阻塞”模式，我们可以设定一个时间，当没有channel可处理时，可以处理超时时间。这也是后台服务器编程常用的处理方式。

没有任何channel准备好，处理默认事件
当没有任何channel准备好数据时，可以设置是执行默认的处理代码。

    func TestDefaultProc() {
        start := time.Now()
        var c1, c2 <-chan int
        select {
        case <-c1:
        case <-c2:

        default:
            fmt.Printf("In default after %v\n\n", time.Since(start))
        }
    }
注意：default和处理超时不同，当没有channel可读取时，会立即执行default分支。而超时的处理，必须要等到超时，才处理。

通过channel通知，从而退出死循环

    func TestExitLoop() {
        done := make(chan interface{})

        go func() {
            time.Sleep(2*time.Second)
            close(done)
        }()

        workCounter := 0
    loop:
        for {
            select {
            case <-done:
                break loop
            default:
            }

            // Simulate work
            workCounter++
            time.Sleep(1*time.Second)
        }

        fmt.Printf("在通知退出循环时，执行了%d次.\n", workCounter)
    }

启动一个goroutine，该goroutine在2s后，关闭channel。此时，主协程会在select中的case <-done分支中得到通知，跳出死循环。而在此之前，会执行default分支的代码，这里是什么都不做。

永远等待

    select{}

以上的语句会永远等待，直到有信号中断。s
package main

import (
	"fmt"
	"sync"
	"time"
)

type User struct {
	Name   string
	Locker *sync.Mutex
}

func (u *User) SetName(wait *sync.WaitGroup, name string) {
	defer func() {
		fmt.Println("Unlock set name:", name)
		u.Locker.Unlock()
		wait.Done()
	}()

	u.Locker.Lock()
	fmt.Println("Lock set name:", name)
	time.Sleep(1 * time.Second)
	u.Name = name
}

func (u *User) GetName(wait *sync.WaitGroup) {
	defer func() {
		fmt.Println("Unlock get name:", u.Name)
		u.Locker.Unlock()
		wait.Done()
	}()

	u.Locker.Lock()
	fmt.Println("Lock get name:", u.Name)
	time.Sleep(1 * time.Second)
}

func main() {
	user := User{}
	user.Locker = new(sync.Mutex)
	wait := &sync.WaitGroup{}
	names := []string{"a", "b", "c"}
	for _, name := range names {
		//等待2个goroutine的执行完之后，再继续遍历
		wait.Add(2)
		go user.SetName(wait, name)
		// time.Sleep(1 * time.Second)
		go user.GetName(wait)
	}
	// 让主goroutine一直阻塞等待
	wait.Wait()
}

// wait默认没有等待任何goroutine，为了让子协程们运行，会在生产子协程的地方，增加等待个数。
// 对应的，主函数尾部添加wait()函数，代表一直在等待子协程，没有则结束，有则等
// 子协程，也需要承担通知的义务，别人在等自己，自己执行完不能一声不吭。具体做法是把wait对象传进入，执行完毕
// 进行wait.Done()或者wait.Add(-1)

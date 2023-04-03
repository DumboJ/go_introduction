package main

import "fmt"

/**
select 语句使一个go 进程可以等待多个信道操作
select 可以阻塞到某个分支可以继续执行时才开始执行该分支；多个分支都准备好，会随机选择一个分支执行

select 监听多个信道 ，实现fibonacci数列
*/
func main() {
	c := make(chan int)
	quit := make(chan int)
	//匿名函数执行，往c channel和quit中添加元素
	//1.开启进程匿名方法
	go func() {
		//2.此处阻塞住，当前信道C中并没有准备好元素
		//select会阻塞到c信道中有数据后才开始执行，将c信道中元素读出，后往quit信道中加入
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		//6.循环完成，开始往quit中放入元素
		quit <- 0
	}()
	//3.执行该函数
	fibonacci_select(c, quit)
}

/*select 计算数值并放入信道c中*/
func fibonacci_select(c chan int, quit chan int) {
	m, n := 0, 1
	for {
		//4.循环体内select 阻塞channel执行
		select {
		//5.循环往channel c中放入计算元素，每次放入步骤2都会执行一次
		case c <- m:
			m, n = n, m+n
		//7.一直等到被阻塞的主方法中循环体次数完成，往quit中加入了0标记值，执行quit 的case
		case <-quit:
			fmt.Println("program quit")
			return
		}
	}
}

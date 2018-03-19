package threads

import "fmt"

func Run() {
	//syncGoroutine()
	cacheChan()
}

func syncGoroutine() {
	g := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case i := <-g:
				fmt.Println(i + 1)
			case <-quit:
				fmt.Println("B quit")
				return
			}
		}
	}()

	for i := 0; i < 3; i ++ {
		g <- i
	}

	quit <- true

	fmt.Println("Main quit")

}

/*
数据的输出有时候需要做扇出／入（Fan In／Out），
但是在函数中调用常常得修改接口，而且上下游对于数据的依赖程度非常高，
所以一般使用通过channel进行Fan In／Out，这样就可以轻易实现类似于shell里的管道。
*/
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()

	return c
}

func cacheChan() {
	jobs := make(chan int)
	done := make(chan bool)

	go func() {
		fmt.Println("GoStart")
		for i := 1; ; i++ {
			fmt.Println("GoforSTART", i)
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
			fmt.Println("GoForEnd", i)
		}
	}()

	for j := 1; j <= 3; j++ {
		fmt.Println("OutFOR", j)
		jobs <- j
		fmt.Println("send job", j)
	}

	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}

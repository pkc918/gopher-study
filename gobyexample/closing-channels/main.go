package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				// 表示还有值
				fmt.Println("sent job", j)
			} else {
				fmt.Println("received all job")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent jobs", j)
	}

	close(jobs)
	fmt.Println("sent all jobs")
	<-done

	_, ok := <-jobs
	fmt.Println("received more jobs", ok)
}

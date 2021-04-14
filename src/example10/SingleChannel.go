package main

func main() {

	ch := make(chan int)

	// 双向channel能隐式的转化成单向channel
	var writeCh chan<- int = ch
	var readCh <-chan int = ch

	writeCh <- 666
	<-writeCh // invalid operation: <-writeCh (receive from send-only type chan<- int)

	<-readCh
	readCh <- 666 // invalid operation: readCh <- 666 (send to receive-only type <-chan int)

	var ch2 chan int = writeCh // cannot use writeCh (type chan<- int) as type chan int in assignment
}

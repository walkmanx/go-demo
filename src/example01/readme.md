## 模块之间的调用

<home>/
 |-- greetings/
 |-- hello/

1. go mod init command	
	$ go mod init example.com/hello

2. go mod edit command	
	$go mod edit -replace=example.com/greetings=../greetings

replace example.com/greetings => ../greetings
3. go mod tidy command
	$ go mod tidy
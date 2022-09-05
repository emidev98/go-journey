# Testing

Go has a testing module integrated in the language it can be found inside the [testing package](https://pkg.go.dev/testing) which allows you to create different testing strategies by asserting the output of the functions you expect to have. To define tests a testing module must be created and the testing functions must receive the parameter **testing.T** to be able to execute these testing cases with the following command:
> go test

## Coverage

The test coverage allows to calculate on a very easy manner how many of your lines of code are being executed in your test cases. To calculate how much test coverage you have you can execute the following command:
> go test -cover

If you executed the previous command it will report the amount of lines are covered by tests. In case you want to export a report to a file the following command can help you:
> go test -coverprofile=coverage.out

Since that report may be hard to read, there is another tool that allow you to read the previously generated file and have a more human-readable format of it:
> go tool cover -func=coverage.out

In case you like to have explicit information about what lines of coder are covered by your tests it can be done very easy with the following command:
> go tool cover -html=coverage.out

## Profiling

Profiling helps to improve the code optimization. As you can see the Fibonacci function takes a while to execute. In order to detect this "malfunctions" in the code you can use the following command to generate a report:
> go test -cpuprofile=cou.out

Afterwards you can create a prompt to the **cou.out** file by using the following command:
> go tool pprof cou.out

The prompt mentioned in the below sentence allows to have the output of how long it took every function to execute by using the command **top**, allows you to see detailed information of execution time by using **list Fibonacci** (as you already noticed it is the function that took a while to execute) and even by using the command **web** you can see a detailed chart of how long it took to execute a specific part of your test on web format.

> <small> in case you are running Linux you may need to install graphviz to generate the diagram from **web** command. </small>
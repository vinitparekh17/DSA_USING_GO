# Stack
> Stack is a linear data structure which follows a particular order in which the operations are performed. The order may be LIFO(Last In First Out) or FILO(First In Last Out).

Stack is very useful and important concept in general programming. So now let's try to understand by assuming a scenario.

Well I am gonna go through really easy scenario so you won't get stuck into it.
Suppose you have simple todo list which you made it using sticky notes. when you add a new task you add a new sticky note to the top of the list. And when you complete a task you remove the sticky note from the top of the list. This is called **stack**.

Let's turn back to our actual topic, Stack is the most easier DS, you will be using it simuntaniously in your programming life without even realising it!


## The call stack

You PC uses a stack internally called **call stack**. It keeps track of what function is currently running and what function called that function and what function called that function and so on.

Remember Recursion function? So when you call a function, it's added to the top of the call stack. And when the function returns, it's removed from the top of the call stack.

Have a look at following table which shows the call stack of the following code:

```go
func main() {
    fmt.Println("main() started")
    foo()
    fmt.Println("main() stopped")
}

func foo() {
    fmt.Println("foo() started")
    bar()
    fmt.Println("foo() stopped")
}

func bar() {
    fmt.Println("bar() started\nbar() stopped")
}
```

| Call stack |
| --- |
| bar() |
| foo() |
| main() |

So basically, your PC is adding func on top of the first func, and when it's done, it gets popped off the stack from the top, decently simple right?

## Stack implementation

Unfortunately, Go doesn't have a built-in stack data structure and slice has dynamic size, so we can't use it as a stack. But we can implement our own stack using slice. and ask user to enter the size of the stack. 

Check out code from [main.go](main.go) file.
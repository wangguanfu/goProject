## err
    1.Error vs Exception
    2.Error Type
    3.Handling Error
    4.Go 1.13 errors
    5.Go 2 Error Inspection
    6.References


### 1.Error vs Exception
> Go error 就是普通的一个接口，普通的值。
>
> 我们经常使用 errors.New() 来返回一个 error 对象。

        
    1.Go 的处理异常逻辑是不引入 exception，支持多参数返回，所以你很容易的在函数签名中带上实现了 error interface 的对象，交由调用者来判定。
    
    2.如果一个函数返回了 (value, error)，你不能对这个 value 做任何假设，必须先判定 error。唯一可以忽略 error 的是，如果你连 value 也不关心。
    
    3.Go 中有 panic 的机制，如果你认为和其他语言的 exception 一样，那你就错了。当我们抛出异常的时候，相当于你把 exception 扔给了调用者来处理。
    比如，你在 C++ 中，把 string 转为 int，如果转换失败，会抛出异常。或者在 Java 中转换 String 为 Date 失败时，会抛出异常。
    
    4.Go panic 意味着 fatal error（就是挂了）。不能假设调用者来解决 panic，意味着代码不能继续运行。
    
    5.使用多个返回值和一个简单的约定，Go 解决了让程序员知道什么时候出了问题，并为真正的异常情况保留了 panic。

```
** 对于真正意外的情况，那些表示不可恢复的程序错误，例如索引越界、不可恢复的环境问题、栈溢出，我们才使用 panic。
对于其他的错误情况，我们应该是期望使用 error 来进行判定。**

简单
考虑失败，而不是成功（plan for failure, not success）
没有隐藏的控制流
完全交给你来控制 error
Error are values
```



### 2.Error Type
**Error are values**
    
    1. 预定义的特定错误，我们叫为 sentinel error，
            使用 sentinel 值是最不灵活的错误处理策略



    2.Error type 是实现了 error 接口的自定义类型。例如 MyError 类型记录了文件和行号以展示发生了什么。
        调用者要使用类型断言和类型 switch，就要让自定义的 error 变为 public。这种模型会导致和调用者产生强耦合，从而导致 API 变得脆弱。
        结论是尽量避免使用 error types，虽然错误类型比 sentinel errors 更好，因为它们可以捕获关于出错的更多上下文，但是 error types 共享 error values 许多相同的问题。
        因此，我的建议是避免错误类型，或者至少避免将它们作为公共 API 的一部分
    
    
    3.Opaque errors  不透明错误
        
        在少数情况下，这种二分错误处理方法是不够的。例如，与进程外的世界进行交互（如网络活动），需要调用方调查错误的性质，以确定重试该操作是否合理。
        在这种情况下，我们可以断言错误实现了特定的行为，而不是断言错误是特定的类型或值。
            
            
### 3.Handling Error
    1.获取上下文信息
    2.更好的处理错误
    3.判定error
    
    
        
##### Wrap erros
        没有生成错误的 file:line 信息。没有导致错误的调用堆栈的堆栈跟踪

        You should only handle errors once. Handling an error means inspecting the error value, and making a single decision.

        
        日志记录与错误无关且对调试没有帮助的信息应被视为噪音，应予以质疑。记录的原因是因为某些东西失败了，而日志包含了答案。
        The error has been logged.
        The application is back to 100% integrity.
        The current error is not reported any longer.
        
        错误要被日志记录。
        应用程序处理错误，保证100%完整性。
        之后不再报告当前错误。
        
        github.com/pkg/errors


        使用 errors.Cause 获取 root error，再进行和 sentinel error 判定。
        总结:
        Packages that are reusable across many projects only return root error values.
            选择 wrap error 是只有 applications 可以选择应用的策略。具有最高可重用性的包只能返回根错误值。此机制与 Go 标准库中使用的相同（kit 库的 sql.ErrNoRows）。
        If the error is not going to be handled, wrap and return up the call stack.
            这是关于函数/方法调用返回的每个错误的基本问题。如果函数/方法不打算处理错误，那么用足够的上下文 wrap errors 并将其返回到调用堆栈中。例如，额外的上下文可以是使用的输入参数或失败的查询语句。确定您记录的上下文是足够多还是太多的一个好方法是检查日志并验证它们在开发期间是否为您工作。
        Once an error is handled, it is not allowed to be passed up the call stack any longer.
            一旦确定函数/方法将处理错误，错误就不再是错误。如果函数/方法仍然需要发出返回，则它不能返回错误值。它应该只返回零（比如降级处理中，你返回了降级数据，然后需要 return nil）。



#### go1.13

    go1.13 errors 包包含两个用于检查错误的新函数：Is 和 As。











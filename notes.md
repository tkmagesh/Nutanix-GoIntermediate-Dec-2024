# Go Intermediate

## Magesh Kuppan
- tkmagesh77@gmail.com

## Schedule
Commence    : 9:30 AM
Tea Break   : 11:00 AM (15 mins)
Lunch Break : 12:30 PM (45 mins)
Tea Break   : 02:45 PM (15 mins)
Wind up     : 04:30 PM

## Methodology
- No powerpoint
- Code & Discussion
- Floor is open at all times during the class for Q & A

## Repository
- https://github.com/tkmagesh/Nutanix-GoIntermediate-Dec-2024

## Software Requirements
- Go Tools (https://go.dev/dl)
- Visual Studio Code

## Pre-requisites 
- Data Types
- Programming Constructs
- Functions
    - Variadic functions
    - Anonymous functions
    - Function Types
    - Higher Order Functions
    - Deferred functions
- Pointers
- Error Handling
- Panic & Recovery
- Structs
    - Struct Composition
- Methods
- Type Assertion
- Interfaces
- Modules & Packages

## Recap
### Panic & Recovery
#### Panic
- A state of the application where the application execution cannot proceed further

## Concurrency
- Managed Concurrency (built in scheduler)
### Wait Group
- Semaphore based counter
- Has the ability to block the execution of the current function until the counter becomes 0
### Channel (data type)
- facilitates "share memory by communicating" strategy
#### Declaration
```go
var <var_name> chan <data_type>
// ex:
var ch chan int
```
#### Initialization
```go
<var_name> = make(chan <data_type>)
// ex:
ch = make(chan int)
```
#### Declaration & Initializatio
```go
var ch chan int = make(chan int)
// OR
var ch = make(chan int)
// OR
ch := make(chan int)
```
### Channel Operator ( <- )
#### Send Operation
```go
<var_name> <- <data>
// ex:
ch <- 100
```
#### Receive Operation
```go
<- <var_name>
// ex:
data := <- ch
```

## Context
### Create a root context (non-cancellable)
    - context.Background()
### Create children (cancellable)
    - context.WithCancel(parentCtx)
    - context.WithTimeout(parentCtx, time.Duration)
    - context.WithDeadline(parentCtx, time.Time)
### Create children (non-cancellable)
    - context.WithValue(parentCtx, key, value)

# Singleton Pattern

Singleton is a creational design pattern, which ensures that only one object of its kind exists and provides a single point of access to it for any other code.

Singleton has almost the same pros and cons as global variables. Although they’re super-handy, they break the modularity of your code.

You can’t just use a class that depends on a Singleton in some other context, without carrying over the Singleton to the other context. Most of the time, this limitation comes up during the creation of unit tests.

## Theory

### Problem

The Singleton pattern solves two problems at the same time, violating the Single Responsibility Principle:

1. **Ensure that a class has just a single instance.** Why would anyone want to control how many instances a class has? The most common reason for this is to control access to some shared resource—for example, a database or a file.

   Here’s how it works: imagine that you created an object, but after a while decided to create a new one. Instead of receiving a fresh object, you’ll get the one you already created.

   Note that this behavior is impossible to implement with a regular constructor since a constructor call **must** always return a new object by design.

2. **Provide a global access point to that instance.** Remember those global variables that you (all right, me) used to store some essential objects? While they’re very handy, they’re also very unsafe since any code can potentially overwrite the contents of those variables and crash the app.

   Just like a global variable, the Singleton pattern lets you access some object from anywhere in the program. However, it also protects that instance from being overwritten by other code.

   There’s another side to this problem: you don’t want the code that solves problem #1 to be scattered all over your program. It’s much better to have it within one class, especially if the rest of your code already depends on it.

Nowadays, the Singleton pattern has become so popular that people may call something a singleton even if it solves just one of the listed problems.

### Solution
All implementations of the Singleton have these two steps in common:

* Make the default constructor private, to prevent other objects from using the new operator with the Singleton class.

* Create a static creation method that acts as a constructor. Under the hood, this method calls the private constructor to create an object and saves it in a static field. All following calls to this method return the cached object.

If your code has access to the Singleton class, then it’s able to call the Singleton’s static method. So whenever that method is called, the same object is always returned.

### Real-World Analogy
The government is an excellent example of the Singleton pattern. A country can have only one official government. Regardless of the personal identities of the individuals who form governments, the title, “The Government of X”, is a global point of access that identifies the group of people in charge.

### Applicability
* **Use the Singleton pattern when a class in your program should have just a single instance available to all clients; for example, a single database object shared by different parts of the program.**

* The Singleton pattern disables all other means of creating objects of a class except for the special creation method. This method either creates a new object or returns an existing one if it has already been created.

* **Use the Singleton pattern when you need stricter control over global variables.**

* Unlike global variables, the Singleton pattern guarantees that there’s just one instance of a class. Nothing, except for the Singleton class itself, can replace the cached instance.

Note that you can always adjust this limitation and allow creating any number of Singleton instances. The only piece of code that needs changing is the body of the `getInstance` method.

### How to Implement
1. Add a private static field to the class for storing the singleton instance.

2. Declare a public static creation method for getting the singleton instance.

3. Implement “lazy initialization” inside the static method. It should create a new object on its first call and put it into the static field. The method should always return that instance on all subsequent calls.

4. Make the constructor of the class private. The static method of the class will still be able to call the constructor, but not the other objects.

5. Go over the client code and replace all direct calls to the singleton’s constructor with calls to its static creation method.

### Pros
* You can be sure that a class has only a single instance.
* You gain a global access point to that instance.
* The singleton object is initialized only when it’s requested for the first time.
### Cons
* Violates the Single Responsibility Principle. The pattern solves two problems at the time.
* The Singleton pattern can mask bad design, for instance, when the components of the program know too much about each other.
* The pattern requires special treatment in a multithreaded environment so that multiple threads won’t create a singleton object several times.
* It may be difficult to unit test the client code of the Singleton because many test frameworks rely on inheritance when producing mock objects. Since the constructor of the singleton class is private and overriding static methods is impossible in most languages, you will need to think of a creative way to mock the singleton. Or just don’t write the tests. Or don’t use the Singleton pattern.
 
### Relations with Other Patterns

* A **Facade** class can often be transformed into a **Singleton** since a single facade object is sufficient in most cases.

* **Flyweight** would resemble **Singleton** if you somehow managed to reduce all shared states of the objects to just one flyweight object. But there are two fundamental differences between these patterns:

   1. There should be only one Singleton instance, whereas a Flyweight class can have multiple instances with different intrinsic states.

   2. The Singleton object can be mutable. Flyweight objects are immutable.

* **Abstract Factories**, **Builders** and **Prototypes** can all be implemented as Singletons.

## Conceptual Example
Usually, a singleton instance is created when the struct is first initialized. To make this happen, we define the `getInstance` method on the struct. This method will be responsible for creating and returning the singleton instance. Once created, the same singleton instance will be returned every time the `getInstance` is called.

**How about goroutines?** The singleton struct must return the same instance whenever multiple goroutines are trying to access that instance. Because of this, it’s very easy to get the singleton design pattern implemented wrong. The example below illustrates the right way to create a singleton.

Some points worth noting:

* There is a `nil`-check at the start for making sure `singleInstance` is empty first time around. This is to prevent expensive lock operations every time the `getinstance` method is called. If this check fails, then it means that the `singleInstance` field is already populated.

* The `singleInstance` struct is created within the lock.

* There is another `nil`-check after the lock is acquired. This is to ensure that if more than one goroutine bypasses the first check, only one goroutine can create the singleton instance. Otherwise, all goroutines will create their own instances of the singleton struct.

### single.go: Singleton
```
package main

import (
    "fmt"
    "sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {
    if singleInstance == nil {
        lock.Lock()
        defer lock.Unlock()
        if singleInstance == nil {
            fmt.Println("Creating single instance now.")
            singleInstance = &single{}
        } else {
            fmt.Println("Single instance already created.")
        }
    } else {
        fmt.Println("Single instance already created.")
    }

    return singleInstance
}
```

### main.go: Client code
```
package main

import (
    "fmt"
)

func main() {

    for i := 0; i < 30; i++ {
        go getInstance()
    }

    // Scanln is similar to Scan, but stops scanning at a newline and
    // after the final item there must be a newline or EOF.
    fmt.Scanln()
}
```
### output.txt: Execution result
```
Creating single instance now.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
```

## Another Example
There are other methods of creating a singleton instance in Go:

1. init function

   We can create a single instance inside the init function. This is only applicable if the early initialization of the instance is ok. The init function is only called once per file in a package, so we can be sure that only a single instance will be created.

2. sync.Once

   The sync.Once will only perform the operation once. See the code below:

### syncOnce.go: Singleton
```
package main

import (
    "fmt"
    "sync"
)

var once sync.Once

type single struct {
}

var singleInstance *single

func getInstance() *single {
    if singleInstance == nil {
        once.Do(
            func() {
                fmt.Println("Creating single instance now.")
                singleInstance = &single{}
            })
    } else {
        fmt.Println("Single instance already created.")
    }

    return singleInstance
}
```

### main.go: Client code
```
package main

import (
    "fmt"
)

func main() {

    for i := 0; i < 30; i++ {
        go getInstance()
    }

    // Scanln is similar to Scan, but stops scanning at a newline and
    // after the final item there must be a newline or EOF.
    fmt.Scanln()
}
```

### output.txt: Execution result
```
Creating single instance now.
Single instance already created.
Single instance already created.
```
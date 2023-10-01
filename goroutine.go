// two goroutines which have a race condition when executed concurrently. \

package main

import (
    "fmt"
    "sync"
)

var (
    counter int
    wg      sync.WaitGroup
)

func increment() {
    for i := 0; i < 10000; i++ {
        counter++
    }
    wg.Done()
}

func main() {
    wg.Add(2)

    go increment()
    go increment()

    wg.Wait()

    fmt.Println("Final Counter Value:", counter)
}


/*
Race Condition Explanation:

A race condition occurs when multiple goroutines access a shared resource concurrently, and at least one of them modifies that resource. In this example, both increment() goroutines are attempting to increment the counter variable without any synchronization mechanism, such as mutex locks or channels.

The race condition can occur because:

Both goroutines read the current value of counter simultaneously.
Both goroutines increment the value.
Both goroutines write the incremented value back to counter.
Depending on the scheduling and timing of these operations, various outcomes are possible:

One goroutine's increment might get lost, leading to an incorrect final value of counter.
Both goroutines might increment the variable correctly, leading to the expected final value.
The order of reads and writes might get interleaved in unpredictable ways, causing the final value of counter to be different from what is expected.
Race conditions can result in unexpected and inconsistent behavior in your program, making it difficult to reason about and debug. To prevent race conditions, you should use synchronization mechanisms such as mutex locks or channels to coordinate access to shared resources among goroutines.
*/
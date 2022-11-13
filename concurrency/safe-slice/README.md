# Introduction
This is a look into how we can safely append slices in a concurrent manner.

# Key takeaways
## Attempt 0
- If a process is expected to completely process, always have a method to wait until the process is finish. ie, use `sync.WaitGroup`.

## Attempt 1
- Data race can occur when you are appending an array with multiple goroutines trying to add their part to the array. During this time, use `mutex` to lock the data access to the shared memory location. (or even better, use channels)

# Conclusion
Using `WaitGroup` and `mutex`, we ensure our goroutines all finish and no data race occur when writing to the same memory.

Note that this does not take into consideration of the order of data write.

# Link
https://blog.devgenius.io/how-to-safely-append-data-to-the-same-slice-concurrently-in-golang-df467e1ebc9c
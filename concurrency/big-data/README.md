# Introduction
This one teaches us how to write a processing tool in a map reduced manner (split, analyse, merge) using GoLang.

Specifically, it will be revolving around the analysis of a giant Excel sheet using concurrent manners.

# Links
https://medium.com/@snassr/processing-large-files-in-go-golang-6ea87effbfe2

# Key takeaways
## Fan-in Fan-out
When wanting to do any form of data processing in a concurrent manner, there will always be 2 main ways to move data around.
- Fan out: where we split up a large input into several goroutines that take a chunk of those input for processing (split of a map reduce)
- Fan in: where we redirect a bunch of goroutines into a single process for processing (merge of a map reduce)

In the example of ours, we first have a single data reader that reads row by row of a large set of data, then packages them into a defined size to be sent for processing. 

The worker receiving end will then wait for data to come into that channel and being processing it immediately. once data is processed, it is then sent out to the output channel.

Another aggregator receiving end will loop through all worker output channel and pass any of them (regardless of order) into the aggregator's output channel (fan-in) and processed that way.

## Channel closing principle
In short
- Don't close channel at the receiving end, but at the sending end. ie, when I'm out of data to send, I'll close my channel. The receiving end should process this using `range` because is a channel is closed, the `range` will exit.
- Don't close a channel if the channel has mulitple concurrent sender.

Pretty nice blog: https://go101.org/article/channel-closing.html
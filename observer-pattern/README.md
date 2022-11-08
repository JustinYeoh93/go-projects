# Introduction
This is an observer pattern example from https://medium.com/mindware-engineering/how-to-implement-observer-pattern-in-go-621365ae8d03

# Key take away
Observer pattern is a one-to-many dependency between a notifier and observers. Where notifier is responsible for notifying any state changes to the observers.

This is usually handled via a map of observers stored in the notifier. Once a state change occurs, the notifier runs through all the observers in its map, and notify them.

We can enhance this further by making the notifying process run in multiple goroutine instead of looping one by one.
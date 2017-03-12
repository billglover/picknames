# picknames
Pick random name(s) from a file

A couple of times in the last week I've been asked to pick some random names out of a list. Every time it has happened, I've found myself in a position where I know some of the people in the list and am conscious that in situations like this it is almost impossible to avoid bias (either positive or negative). Naturally I reached for a quick script to pick names at random from the list, problem solved.

Or was it...

Whilst my approach worked it clearly wasn't the most efficient way of picking random lines out of a text file. Out of sheer curiosity, I'm planning on benchmarking a couple of different approaches.

 > Given a file containing a list of names (one per line), return a number `n` names picked at random from the list. No name should be picked more than once.

## Attempt 1

My first attempt wasn't written in Go, but for the sake of comparison between alternative approaches, I'm re-writing it here to enable comparison with alternative approaches.

 1. read the contents of the file into memory
 2. assign each entry a random number
 3. sort the list of names according to the number
 4. pick the first `n` names from the list


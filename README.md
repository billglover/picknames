# picknames

Pick random name(s) from a file

A couple of times in the last week I've been asked to pick some random names out of a list. Every time it has happened, I've found myself in a position where I know some of the people in the list and am conscious that in situations like this it is almost impossible to avoid bias (either positive or negative). Naturally I reached for a quick script to pick names at random from the list, problem solved.

Or was it...

Whilst my approach worked it clearly wasn't the most efficient way of picking random lines out of a text file. Out of sheer curiosity, I'm planning on benchmarking a couple of different approaches.

 > Given a file containing a list of names (one per line), return a number `n` names picked at random from the list. No name should be picked more than once.

## Attempt 1 - the first result on Google

My first attempt wasn't written in Go, but for the sake of comparison between alternative approaches, I'm re-writing it here to enable comparison with alternative approaches.

 1. read the contents of the file into memory
 1. assign each entry a random number
 1. sort the list of names according to the number
 1. pick the first `n` names from the list

Based on this PowerShell script: [Use PowerShell to Pick Random Winning Users from Text](https://blogs.technet.microsoft.com/heyscriptingguy/2011/09/08/use-powershell-to-pick-random-winning-users-from-text/)

## Attempt 2 - pick lines with probability 1/n

Now that we have a baseline, I wanted to try something different. Rather than trying to replicate the PowerShell script I found, I wanted to try implementing this based on the approach that best fits how I would have told someone else to solve it.

 1. count the number of lines in the file, `lines`
 1. assume each line has equal probability of being picked `1 / lines`
 1. repeat `n` times:
    1. pick a random number `r` in the range `1..lines`
    1. if that number hasn't already been picked, return `lines[r]`

***note:** for descriptive purposes we have assumed the first line is referenced as `lines[1]`*

## Benchmarking Results

| Attempt | Pick 1     | Pick 5     | Pick 10    |
| ------- | ---------- | ---------- | ---------- |
| 1       | 5424 ns/op | 5368 ns/op | 5629 ns/op |
| 2       |            |            |            |
|         |            |            |            |

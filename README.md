# ThirdTime

## What is ThirdTime?

Here’s an example of the basic procedure:

1. Note the time, or start a stopwatch
2. Work for as long or short as you like, until you want or need to break
3. Suppose you worked for 45 minutes. This earns you 45 ÷ 3 = 15 minutes off; so set an alarm for 15 minutes
4. Break until the alarm goes off
5. Go back to step 1.

<!-- add image -->
![ThirdTime](
    docs/breaks.png)

You can read more about the technique 
[here](https://www.lesswrong.com/posts/RWu8eZqbwgB9zaerh/third-time-a-better-way-to-work)

## How to use ThirdTime? 

1. Download the latest binary from the [releases page](https://github.com/ghodsizadeh/third-time/releases), or by using `go install`:
```bash
$ go install github.com/ghodsizadeh/third-time
```
2. Run the binary
```bash
$ ./third-time

Elapsed: 30m00.000s

s stop • t start rest • r reset • q quit
```
3. When you finish your work, press `t` to start the rest time for 1/3 of the work time
```bash
Resting: 10m00.000s

s stop • t start rest • r reset • q quit
```
4. 
> But instead of one-third, you can use whatever fraction you like, such as:
> - 1/2: 40 mins work + 20 mins breaks per hour. Working 2/3 of the time. Lazy
> - 1/3: 45 mins work + 15 mins breaks per hour. Working 3/4 of the time. Standard
> - 1/4: 48 mins work + 12 mins breaks per hour. Working 4/5 of the time. Industrious
> - 1/5: 50 mins work + 10 mins breaks per hour. Working 5/6 of the time. Hard
> -  1/6: 51½ mins work + 8½ mins breaks per hour. Working 6/7 of the time. Grinding 
   
   You can change the fraction of the rest time by adding `-f` flag
```bash
$ ./third-time -f 4
```
5. You can check history by adding `x` flag
```bash

$ ./third-time x
Duration,Date,Fraction
2,2024-05-23,21:37,3
1,2024-05-23,21:55,3
Average time: 0 minutes
```



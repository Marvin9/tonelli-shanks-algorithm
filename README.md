[This](https://discuss.codechef.com/t/lcasqrt-editorial/82141) problem on codechef motivated me to implement [tonelli-shanks](https://en.wikipedia.org/wiki/Tonelli%E2%80%93Shanks_algorithm#Speed_of_the_algorithm).

```
Given non-zero integar n and prime number p, It finds R such that,

(R)^2 congruence n (mod p)
```

[Algorithm](https://en.wikipedia.org/wiki/Tonelli%E2%80%93Shanks_algorithm#The_algorithm)

## Algorithm in nutshell:

> Input: n, p
> Goal: R, where (R)^2 ≡ n (mod p); If solution exists else -1

1. Check if solution exists using Euler's criterion. Return -1 if not.

2. Find ```Q and S``` such that ```p - 1 = Q * (2)^S```
> divide p - 1 by 2 until it's modulo 2 is not equal to 0. Keep count number of division, store it as ```S```, remaining value (of ```p-1```) is ```Q```.

3. Find smallest ```z``` which is quadratic non-residue.

> continue ```i = 2 to (p-1)``` until euler's criterion (i, p) is ```1```. Store ```z = i``` at end.

4. Define some variables
```bash
m = S
c = (z)^Q
t = (n)^Q
R = (n)^((Q+1)/2)
```

> Note: Make sure all operations don't go beyond ```p```. For example, (3)^3 and ```p = 10``` => ans = ```[[[[(1 * 3)%10]*3]%10]*3]%10```

5. Infinite loop,

- if `t = 0`, return `0`
- if `t = 1`, return `R`
- For i = ```(1 to m - 1)```, such ```(t)^((2)^i) % p = 1```
> Note: Operations must not go beyond ```p```. 
- `b = (c)^((2)^(M - i - 1))`
- `M = i`
- `c = (b)^2 % p`
- `t = t * (b)^2 % p`
- `R = (R * b) % p`
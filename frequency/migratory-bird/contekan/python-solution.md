count = [0]*6, maksudnya apa?
inisiasi array of int 
len(arr) = 6

count.index(max(count))

count.index() ini apa?
max(count) ini apa?

max(count) = nilai maximum dari array of int
```python
input()
arr = map(int,input().strip().split())

count = [0]*6
for t in arr:
    count[t] += 1

return count.index(max(count))




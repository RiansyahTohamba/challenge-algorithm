```python
noUppSpec = "abcdexa111$A"
print(minimumNumber(8,noUppSpec))

def minimumNumber(n, password):
    count = 0    

    if any(i.isdigit() for i in password)==False:
        count+=1

    if any(i.islower() for i in password)==False:
        count+=1

    if any(i.isupper() for i in password)==False:
        count+=1

    if any(i in '!@#$%^&*()-+' for i in password)==False:
        count+=1

    return max(count, 6-n)

```
# any() apa nih?
any(i.isdigit() for i in password)

i.isdigit() for i in password 
    cek i.isdigit?
        
## syntax apa nih?
The `any() function returns True` if any item in an iterable are true, otherwise it returns False. 
If the iterable object is empty, the any() function will return False.

any() menggunakan iterable function

# max(count,6-n) 
max itu apa?

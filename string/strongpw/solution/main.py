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

    return max(count,6-n)

# apa padanan dari any() di golang?
# apa padanan dari max() di golang?

# ga ada uppercase
# ga ada special character
# noUppSpec = "abcdexa11111"
# print(minimumNumber(8,noUppSpec))
# # complete
noUppSpec = "abcdexa11A"
print(minimumNumber(8,noUppSpec))

# print(minimumNumber(3, "lalalal"))
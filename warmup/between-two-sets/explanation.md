Sample Input

len dari arr1 dan arr2 = [2 3]
[2 4] = len(arr1) = 2
[16 32 96] = len(arr2) = 3

Sample Output
3


Explanation
2 and 4 divide evenly into 4, 8, 12 and 16.
2 dan 4 dibagi rata menjadi 4, 8, 12 dan 16.

2 dan 4 adalah faktor pada 16,32,96

4, 8, 16 , ini apa?



16
32
96



4, 8 and 16 divide evenly into 16, 32, 96.
4, 8 and 16 dibagi rata menjadi 16, 32, 96.


4, 8 and 16 are the only `three numbers` for which each `element of a` is `a factor` and `each is a factor of all elements of b.`

4,8,16 merupakan 3 nomor yang untuk setiap element a adalah sebuah faktor dan setiap faktor untuk semua element b.




# lagi

There will be two arrays of integers. Determine all integers that satisfy the following two conditions:

1. The elements of the first array are all factors of the integer being considered
[4,8,16]

2. 'The integer being considered' is a factor of all elements of the second array
[4,8,16]


These numbers are referred to as being between the two arrays. 
Determine how many such numbers exist.

Example
a = [2,6]
b = [24,36]

There are two numbers between the arrays:  [6 dan 12].

6 % 2 = 0, 
6 % 6 = 0, 


24 % 6 = 0 and 
36 % 6 = 0 for the first value.

12 % 2 = 0, 
12 % 6 = 0, 


24 % 12 = 0,  
and
36 % 12 = 0,  for the second value. 

Return 2.


2 dan 6 adalah faktor dari [6,12]
[6,12] adalah faktor dari [24,36]
begitu rumusnya.

maka kalau ketemu soal
[2 4]
[16 32 96]

algebra nya adalah:
2,4 adalah faktor dari [w,x,y,z]
[x,y,z] adalah faktor dari [16,32,96]

jawabannya:
len(element) = 3

2 dan 4 faktor bagi [4,8,12,16] 
[4, 8, 16] adalah faktor dari [16, 32, 96]
kenapa 2 ga masuk?

2 dan 4 merupakan faktor [4,8,12,16] 










## apa itu x being considered?
1. The elements of the first array are all factors of the integer being considered.
2. 'The integer being considered' is a factor of all elements of the second array

return dari fungsi ini adalah `single integer`, bukan `array of integers`. maka integer being considered adalah `that returned integer`

1. faktor dari integer yang mau dicari hasilnya
2. integer yang mau dicari hasilnya adalah faktor dari semua elements yang ada pada array ke dua


element dari first array adalah faktor dari x
semua element dari second array adalah 

contoh 
1. [2,6]
2. [24,36]

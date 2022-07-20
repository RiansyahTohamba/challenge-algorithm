SLICE
pangkat-2 fixed size for array

# kenapa perlu capacity?
jarak antar pangkat-2 yang jauh

misalkan saya punya array dengan panjang 17
maka array akan dibuatkan capacity nya sebesar 32.

mari kita lihat deret bilangan per-pangkatan 2

2, 4, 8, 16, 32,...

maka jika kita sudah yakin panjang fix nya, kita boleh define capacity nya, agar lebih efisien.

# kenapa pangkat-2?
karena untuk menampung semua kemungkinan bilangan bit yang terdiri dari 2 angka, 1 or 0.

2 = 2 ^ 1 = 1 bit = 1 ; 0
4 = 2 ^ 2 = 2 bit = 00; 11 ; 10 ; 01
8 = 2 ^ 3 = 3 bit = 000 ; 111 ; 001 ; ...

jika 1 bit maka berarti hanya menampung 1 atau 0
jika 2 bit maka berarti hanya menampung 00, 01, 10 atau 11

# pada kasus stack, kenapa capacity sudah fix?
karena tidak mungkin kita masukkan 17 elemen pada stack, lalu stack nya mendapatkan tambahan 18..32.
bahkan jika 18..32 itu sudah berisi 0.

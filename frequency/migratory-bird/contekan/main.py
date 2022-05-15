def migratoryBirds(arr):
    count = [0]*6
    # t = value dari array
    for t in arr:
        count[t] += 1

    print(count)
    print("max_value(count) == frekuensi tertinggi")
    print("apakah max.value dari array count?")
    maxVal = max(count)
    print(maxVal)
    idxMaxVal = count.index(maxVal)
    print("apakah index untuk max value?")
    print(idxMaxVal)
    return idxMaxVal


if __name__ == '__main__':
    arr = [1,2,1,2,4,3,3]
    result = migratoryBirds(arr)
    print("apa saja himpunan dari bird id?")
    print(arr)
    print("apakah id terkecil + frekuensi terbesar pada array bird?")
    print(result)

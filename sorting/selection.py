def selection_sort(a):
    smallest = 0

    for i in range(len(a)):
        for j in range(i, len(a)):
            if a[j] < a[smallest]:
                smallest = j
        a[smallest], a[i] = a[i], a[smallest]
    
    return a


arr = [10, 100, 1, 82, 102, 89, 84]
print(selection_sort(arr))

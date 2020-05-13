def main():
    a = [10, 4, 2, 1, 5, 6]
    s = 10
    print(find_continued_sub_arr(a, s))

def find_continued_sub_arr(A, S):

    i, j, sum = 0, 0, 0

    while i < len(A) and j < len(A) and sum != S:
        sum += A[j]
        if sum < S:
            j += 1
            print("sum: ", sum)
        elif sum > S:
            print("sum larger than S: ", sum)
            if i < j:
                sum -= A[i]
                i += 1
            else:
                sum = 0
                i = j = i + 1


    return None if sum != S else (i,j)


if __name__ == '__main__':
    main()

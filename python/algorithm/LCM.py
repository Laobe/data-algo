# coding: utf-8

""" 最大公共子矩阵
"""

A = [[0, -2, -7, 0],
     [9, 2, -6, 2],
     [-4, 1, -4, 1],
     [-1, 8, 0, -2],]
N = 4
F = []
Q = []
P = []


def LCS(lists=None):
    if lists is None:
        return 0

    max_sum = 0
    all_sum = 0
    for i in lists:
        all_sum += i
        if all_sum < 0:
            all_sum = 0
        else:
            max_sum = max(max_sum, all_sum)
    return max_sum


def LCM():
    max_sum = 0
    for i in range(len(A)):
        for j in range(i+1, len(A)):
            F = []
            for k in range(len(A[j])):
                F.append(sum(A[k][i:j]))
            if max_sum <= LCS(F):
                max_sum = LCS(F)
                right = i
                left = j
    return max_sum, right, left


if __name__ == '__main__':
    print(LCM())

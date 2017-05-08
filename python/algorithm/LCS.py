# coding: utf-8

""" 最大子序列和    The longest common subsequence
"""


def MaxSubSequence1(lists=None):
    """ 枚举法 O(N^3)
    """
    if lists == None:
        lists = []
    max_sum = 0
    for i in range(len(lists)):
        this_sum = 0
        for j in range(i, len(lists)):
            this_sum = sum(lists[i:j+1])
            max_sum = max(max_sum, this_sum)
    return max_sum


def MaxSubSequence2(lists=None):
    """ 枚举法 O(N^2)
    """
    if lists == None:
        lists = []
    max_sum = 0
    for i in range(len(lists)):
        this_sum = 0
        for j in lists[i+1:]:
            this_sum += j
            max_sum = max(max_sum, this_sum)
    return max_sum


def MaxSubSequence4(lists=None):
    """ 动态规划法 O(N)
    """
    if lists == None:
        lists = []
    max_sum = 0
    this_sum = 0
    for j in lists:
        this_sum += j
        if this_sum <= 0:
            this_sum = 0
        else:
            max_sum = this_sum
    return max_sum



if __name__ == '__main__':
    lists = [11, -32, 14]
    print(MaxSubSequence1(lists))
    print(MaxSubSequence2(lists))
    print(MaxSubSequence4(lists))

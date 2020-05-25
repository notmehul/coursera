def isPalindrome(s):
    '''
    str -> bool
    true is 's' is a palindrome, otherwise not
    '''
    i = 0
    j = len(s) -1
    # i an j have been given value of the twi end of the string
    # from the below while loop, i and j will get closer nd closer in value till the middle
    # i is starting from the left side, so i is always less than j
    # when i becomes larger than j, it can be assumed that the array check has ended lol obviously
    while i < j and s[i] == s[j]:
        i += 1
        j -= 1
    # if the middle of the string was reached, then the string is a palindrome, else it's not 
    # because there is an and clause in while
    # hence i and j won't reach case where j comes close to i
    return j <= i 

x = str(input())

print(isPalindrome(x))
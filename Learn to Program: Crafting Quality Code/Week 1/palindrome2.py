def isPalindrome(s):
    '''
    str -> bool
    this function returns true if the string 's' is a palindrome else, false
    '''
    # n is the number os chars in s
    n = len(s)
    #compare the first half to the second half
    #and omitting the middle is it's an odd number
    return s[ : n // 2] == reverse(s[ n - n//2: ])

def reverse(s):
    '''
    str -> str
    this function reverses the string which has been input
    ''' 
    rev = ''
    # for ever character in s, add it to the beginning of the rev
    for ch in s:
        rev = ch + rev
    return rev

string = str(input())

print(isPalindrome(string))
def isPalindrome(s):
    '''
    str -> bool
    this function returns true if the string 's' is a palindrome else, false
    '''
    return reverse(s) == s

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
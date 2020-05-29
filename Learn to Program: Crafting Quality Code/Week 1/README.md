# Notes

# Table of Contents:
1. [Approaching the Problem](#palindrome-approaching-the-problem)
2. [Algorithm 1](#palindrome-algorithm-1)
    * [Code](#code)
3. [Algorithm 2](#palindrome-algorithm-2)
    * [Code](#code-1)
4. [Algorithm 3](#palindrome-algorithm-3)
	* [Code](#code-2)
5. [Restaurant Recomendation](#the-restaurant-recommendation-problem)
	* [Task](#task)
	* [The Problem](#the-problem)
	

# Palindrome: Approaching the Problem
## Introduction

A palindrome is a word that is read the same front-to-back or back-to-front. For instance, R A C E C A R reversed is also R A C E C A R. However, BLUE reversed is EULB, and thus not a palindrome.
Algorithms

An algorithm is a sequence of steps that accomplish a task. We present three algorithms for detecting palindromes:
* Algorithm 1: Reverse the string. Compare the reversed string to the original string.

    For example, the reverse of the string "noon" is "noon". Since the reversed string is the same as the original string, "noon" is a palindrome.

* Algorithm 2: Split the string into two halves. Reverse the second half. Compare the first half to the reversed second half.

    For example, the first half of the string "noon" is "no" and the second half is "on". The reverse of the second half is "no". Since the first half is equal to the reverse of the second half, "noon" is a palindrome.
    For a string with an odd length, let's consider "racecar". When splitting the string into two halves, we omit the middle character, "e". The first half of the string is "rac" and the second half of the string is "car". The reverse of the second half is "rac". Since the first half is equal to the reverse of the second half, "racecar" is a palindrome.

* Algorithm 3: Compare the first character to the last, the second to the second last, and so on. Stop when the middle of the string is reached.

    For example, for string "noon", we compare the first character ("n") to the last character ("n"), and the second character ("o") to the second last ("o"). Since both pairs of characters that were compared are equal, "noon" is a palindrome.
    For a string with an odd length, let's consider "racecar". We compare the first character ("r") to the last character ("r"), the second character ("a") to the second last ("a"), and the third character ("c") to the third last character ("c"). The middle character, "e" does not need to be compared with anything. Since all pairs of characters that were compared are equal, "racecar" is a palindrome.

In upcoming lectures, we will implement all three algorithms. Following the Recipe for Designing Functions, we implemented the function header:

    Examples
    >>> is_palindrome('noon')
    True
    >>> is_palindrome('racecar')
    True
    >>> is_palindrome('dented')
    False
    Type Contract
    (str) --> boolean
    Header
    def is_palindrome(s):
    Description
    Return True if and only if s is a palindrome
    Body
    This will be discussed in future lectures.
    Test
    This will be discussed in future lectures.

# Palindrome: Algorithm 1
## Overview

To determine whether a string is a palindrome, the first algorithm we explored was:

    Reverse the string.
    Compare the reversed string to the original string.

For example, the reverse of the string "noon" is "noon". Since the reversed string is the same as the original string, "noon" is a palindrome.

For the string "dented", the reverse of the string "dented" is "detned". Since the reversed string is not the same as the original string, "dented" is not a palindrome.

### Code

```python
def reverse(s):
    """ (str) -> str

    Return s reversed.
    
    >>> reverse('hello')
    'olleh'
    """

    s_reversed = ''
    for ch in s:
        s_reversed = ch + s_reversed
    return s_reversed


def is_palindrome_v1(s):
    """ (str) -> bool

    Return True if and only if s is a palindrome.

    >>> is_palindrome_v1('noon')
    True
    >>> is_palindrome_v1('radar')
    True
    >>> is_palindrome_v1('kayaks')
    False
    """

    s_reversed = reverse(s)
    return s == s_reversed
```

# Palindrome: Algorithm 2
## Overview

To determine whether a string is a palindrome, the second algorithm we explored was:

    Split the string into two halves.
    Reverse the second half.
    Compare the first half to the reversed second half.

For example, the first half of the string "noon" is "no" and the second half is "on". The reverse of the second half is "no". Since the first half is equal to the reverse of the second half, "noon" is a palindrome.

For a string with an odd length, let's consider "racecar". When splitting the string into two halves, we omit the middle character, "e". The first half of the string is "rac" and the second half of the string is "car". The reverse of the second half is "rac". Since the first half is equal to the reverse of the second half, "racecar" is a palindrome.

Finally, for a string that is not a palindrome, let's consider "dented". The first half of the string is "den" and the second half of the string is "ted". The reverse of the second half is "det". Since the first half is not equal to the reverse of the second half, "dented" is not a palindrome.

### Code
```python
def is_palindrome_v2(s):
    """ (str) -> bool

    Return True if and only if s is a palindrome.

    >>> is_palindrome_v2('noon')
    True
    >>> is_palindrome_v2('racecar')
    True
    >>> is_palindrome_v2('dented')
    False
    """

    # The number of chars in s.
    n = len(s)

    # Compare the first half of s to the reverse of the second half.
    # Omit the middle character of an odd-length string.
    return s[:n // 2] == reverse(s[n - n // 2:])

def reverse(s):
    """ (str) -> str

    Return a reversed version of s.

    >>> reverse('hello')
    'olleh'
    >>> reverse('a')
    'a'
    """

    rev = ''

    # For each character in s, add that char to the beginning of rev.
    for ch in s:
        rev = ch + rev

    return rev
```
# Palindrome: Algorithm 3
## Overview

To determine whether a string is a palindrome, the third algorithm we explored was:

    Compare the first character to the last, the second to the second last, and so on.
    Stop when the middle of the string is reached.

For example, for string "noon", we compare the first character ("n") to the last character ("n"), and the second character ("o") to the second last ("o"). Since both pairs of characters that were compared are equal, "noon" is a palindrome.

For a string with an odd length, let's consider "racecar". We compare the first character ("r") to the last character ("r"), the second character ("a") to the second last ("a"), and the third character ("c") to the third last character ("c"). The middle character, "e" does not need to be compared with anything. Since all pairs of characters that were compared are equal, "racecar" is a palindrome.

Finally, for a string that is not a palindrome, let's consider "dented". We compare the first character ("d") to the last character ("d"), the second character ("e") to the second last ("e"), and the third character ("n") to the third last character ("t"). Since not all pairs of characters that were compared are equal, "dented" is a not palindrome.

### Code

```python
def is_palindrome_v3(s):     """ (str) -> bool

    Return True if and only if s is a palindrome.

    >>> is_palindrome_v3('noon')
    True
    >>> is_palindrome_v3('racecar')
    True
    >>> is_palindrome_v3('dented')
    False
    """

    # s[i] and s[j] are the next pair of characters to compare.
    i = 0
    j = len(s) - 1

    # The characters in s[:i] have been successfully compared to those in s[j:].
    while i < j and s[i] == s[j]:
        i = i + 1
        j = j - 1

    # If we exited because we successfully compared all pairs of characters,
    # then j <= i.
    return j <= i
```

# The Restaurant Recommendation Problem

While learning, students are often told which functions to write, including the function names, parameters and return types.

The next set of videos focus on an exercise where you get to decide yourself which functions to write. The example also gives you more practice with dictionaries, lists, and files.

### Task

The problem we'll be tackling is a restaurant recommendation system.

We are given a list of restaurants that contains:

    The name of the restaurant.
    The percentage of people who recommended the restaurant.
    The price range of the restaurant.
    The type of food served by the restaurant. 

The program will make a recommendation to the user based on this data.

We'll tell you a bit about the main function, but this task is complex and helper functions will make writing the main function much easier.
### The Problem:

Write a function that has three parameters:

    a restaurant file that is open for reading,
    the price range (one of $, $$, $$$ and $$$$), and
    a list of cuisines. 

The function returns a list of restaurants (in that price range, serving at least one of those cuisines), and their ratings sorted from highest to lowest.

## Representing the Data

The first step to solving the restaurant recommendations problem is choosing data structures to store the information on restaurant prices, ratings, and cuisines.
Examining the Data File

Here is the restaurant data from a sample file:
Georgie Porgie
87%
$$$
Canadian, Pub Food

Queen St. Cafe
82%
$
Malaysian, Thai

Dumplings R Us
71%
$
Chinese

Mexican Grill
85%
$$
Mexican

Deep Fried Everything
52%
$
Pub Food

### Examining the Data
We'll organize the data by grouping it according to tasks that we would like to perform.

#### Rating Information

For each restaurant, we want to be able to look up the rating, so we'll keep track of that information:
Georgie Porgie: 87
Queen St. Cafe: 82
Dumplings R Us: 71
Mexican Grill: 85
Deep Fried Everything: 52

#### Price Range Information

We'll also want to look up restaurants by price, so we'll make a list of that data:
$: Queen St. Cafe, Dumplings R Us, Deep Fried Everything
$$: Mexican Grill
$$$: Georgie Porgie
$$$$:

#### Cuisine Information

Recommendations are made based on types of cuisine as well, so we need to keep track of that information:
Canadian: Georgie Porgie
Pub Food: Georgie Porgie, Deep Fried Everything
Malaysian: Queen St. Cafe
Thai: Queen St. Cafe
Chinese: Dumplings R Us
Mexican: Mexican Grill

### Choosing the data structure

What data structures can we use to store this information? We could use strings, lists, tuples or dictionaries. That is a design decision that we need to make.
Rating Information

Our "Rating Information" looks a lot like a Python dictionary, where each key is a restaurant name and each value is a rating.

Let's add some braces and commas to make the structure look more like a Python dictionary. Let's also add quotes around the restaurant names to make them strings. Finally, let's create a variable name name_to_rating that refers to this dictionary:
Georgie Porgie: 87
Queen St. Cafe: 82
Dumplings R Us: 71
Mexican Grill: 85
Deep Fried Everything: 52
	  	→ 	  	
name_to_rating = {'Georgie Porgie': 87,
'Queen St. Cafe': 82,
'Dumplings R Us': 71,
'Mexican Grill': 85,
'Deep Fried Everything': 52}

Now, to find ratings, we can use the restaurant's name.

>>>name_to_rating['Queen St. Cafe']
82

This dictionary type can be written as: dict of {str: int}

### Pricing Information

Pricing information also looks a lot like a Python dictionary. We would like to be able to look up a price and get all the restaurants in the price range.

The keys (price ranges) look like strings, and the values (restaurant names) look like strings too; however, there can be zero, one or more than one restaurant associated with each price range. Therefore, each value will be a list of str.

Again, adding quotes, commas, brackets, braces and a variable name gives us:
$: Queen St. Cafe, Dumplings R Us, Deep Fried Everything
$$: Mexican Grill
$$$: Georgie Porgie
$$$$:
	  	→ 	  	
price_to_names = {'$': ['Queen St. Cafe', 'Dumplings R Us', 'Deep Fried Everything'],
'$$': ['Mexican Grill'],
'$$$': ['Georgie Porgie'],
'$$$$': []}

We can use this dictionary to find restaurants in a given price range.
>>>price_to_names['$']
['Queen St. Cafe', 'Dumplings R Us', 'Deep Fried Everything']

This dictionary type can be written as: dict of {str: list of str}
Cuisine Information

A dictionary is also suitable for representing cuisine information.

In this case, the type will be dict of {str: list of str}. Each key will be a cuisine and each value will be a list of str, since there can be more than one restaurant for each type of cuisine.

Once again, adding quotes, commas, brakets, braces and a variable name gives us:
Canadian: Georgie Porgie
Pub Food: Georgie Porgie, Deep Fried Everything
Malaysian: Queen St. Cafe
Thai: Queen St. Cafe
Chinese: Dumplings R Us
Mexican: Mexican Grill
	  	→ 	  	
cuisine_to_name = {'Canadian': ['Georgie Porgie'],
'Pub Food': ['Georgie Porgie', Deep Fried Everything'],
'Malaysian': ['Queen St. Cafe'],
'Thai': ['Queen St. Cafe'],
'Chinese': ['Dumplings R Us'],
'Mexican': ['Mexican Grill']}

We can use this dictionary to find restaurants that serve a particular type of cuisine.
>>>cuisine_to_name['Chinese']
['Dumplings R Us'] 

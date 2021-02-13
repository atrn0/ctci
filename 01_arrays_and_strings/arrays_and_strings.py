from collections import Counter
from typing import Any, List, Union


"""
ハッシュテーブルの実装はここから引用してきましたー
https://tkr1205.hatenablog.com/entry/2020/12/17/155503
"""

class LinkedList:
    def __init__(self, value):
        self.value = value
        self.next = None

class MyDict:
    def __init__(self):
        self.table = [None] * 1000

    def hash(self, key):
        # 1000で割った余りをhash値として扱う
        return key % 1000

    def add(self, key, value):
        hashed_key = self.hash(key)
        if self.table[hashed_key]:  # すでにキーにデータが存在したら
            ll = self.table[hashed_key]  # すでに存在しているデータ(連結リストの先頭)
            while ll:  # 連結リストの最後尾までループして，最後にデータを追加する
                if not ll.next:  # 連結リストの最後の場合
                    ll.next = LinkedList(value)  # 新しい値を連結リストに連結
                    break
                else:
                    ll = ll.next
        else:
            self.table[hashed_key] = LinkedList(value)  # データが既に存在していない場合，連結リストとしてデータを挿入

    def get(self, key):
        values = []
        hashed_key = self.hash(key)
        ll = self.table[hashed_key]
        if not ll:  # 指定したキーにデータが存在しない場合は-1を返す
            return -1
        while ll:  # 連結リストが存在する場合
            values.append(ll.value)  # 連結リストの値をリストに追加
            if not ll.next:  # もし連結リストの最後尾だったら
                return values
            else:
                ll = ll.next

    def remove(self, key, value):
        hashed_key = self.hash(key)
        ll = self.table[hashed_key]
        if not ll:  # 指定したキーにデータが存在しなかったら
            print('No Data')
            return
        if ll.value == value:  # 連結リストの先頭を削除する場合
            if ll.next:  # 先頭を削除して，連結リストを一つ前にずらす
                self.table[hashed_key] = ll.next
            else:  # データが一つだけの場合
                self.table[hashed_key] = None
            print(f'Key:{key}, Value:{value} Removed')
            return
        ll_prev = ll
        ll = ll_prev.next
        while ll:  # 指定したキーに複数連結リストが存在する場合
            if ll.value == value:
                ll_prev.next = ll.next
                print(f'Key:{key}, Value:{value} Removed')
                return
            else:
                ll_prev = ll
                ll = ll.next
        print('Data not Found')


def merge(word: list, more: list) -> list:
    sentence = []
    for w in word:
        sentence.append(w)
    for w in more:
        sentence.append(w)
    return sentence

def joinWords1(words: list) -> str:
    sentence = ""
    for w in words:
        sentence += w
    return sentence

def joinWords2(words: list) -> str:
    return "".join(words)


"""
1.1 Is Unique:

Implement an algorithm to determine if a string has all unique characters. What if you cannot use additional data structures?
"""

# O(n)(or O(1) cuz it will not more than 128)
def IsUnique(input: str) -> bool:
    if len(input) > 128:
        return False
    char_set = [None]*128
    for i in input:
        val = ord(i)
        if char_set[val]:
            return False
        char_set[val] = True
    return True

#O(n^2) without data stractures
def IsUnique2(input: str) -> bool:
    for i in range(len(input)):
        for j in range(i+1,len(input)):
            if input[i] == input[j]:
                return False
    return True

#or sort it first O(nlogn)

"""
boolean isUnique(String str){
  int checker = 0;
  for (int i = 0; i < str.length(); i++){
    int val = str.charAt(i) - "a";              #実はここが読めない...
    if ((checker & (1 << val)) > 0) {
      return false
    }
    checker |= (1 << val);
  }
  return true;
}
"""



"""
1.2 Check Permutation:

Given two strings, write a method to decide if one is a permutaion of the other.
"""

def CheckPermutation(str1: str, str2: str) -> bool:
    if len(str1) != len(str2):
        return False
    temp1 = sorted(str1)
    temp2 = sorted(str2)
    return temp1 == temp2
    

def CheckPermutation2(str1: str, str2: str) -> bool:
    if len(str1) != len(str2):
        return False
    temp1 = Counter(str1)
    temp2 = Counter(str2)
    return temp1 == temp2    


"""
1.3 URLify:

Write a method to replace all spaces in a string with '%20'.
You may assume that the string has sufficient space at the end to hold the additional characters,
and that you are given the "true" length of the string.
(Note: If implementing in Java, please use a charaster array so that you can perform this operation in place.)
EXAMPLE
Input:      "Mr John Smith    ", 13
Output:     "Mr%20John%20Smith"
"""

def URLify(sentence: str, len: int) -> str:
    if sentence == "":
        return ""
    res = sentence[:len]
    return res.replace(" ", "%20")


"""
1.4 Palindrome Permutation:
Given a string, write a function to check if it is a permutation of a palindrome.
A palindrome is a word or phrase that is the same forwards and backwords.
A permutation is a rearrangement of letters. The palindrome does not neet to be limited to just dictionary words.
EXAMPLE
Input:      Tact Coa
Output:     True (permutations: "taco cat", "atco ata", etc.)
"""

def PalindromePermutation(input: str) -> bool:
    if len(input) <= 1:
        return True
    temp = Counter(input.replace(" ","").lower())
    flag = False
    for i in temp:
        if temp[i] % 2 == 1:
            if flag:
                return False
            else:
                flag = True
    return True

"""
1.5 One Way:

There are three types of edits that can be performed on strings: Insert a character, remove a character, pr replace a character.
Given two string, write a function to check if they are one edit (or zero edits) away.
EXAMPLE
pale,   ple     -> true
pales, pale    -> true
pale, bale      -> true
pale, bake      -> false
"""

def OneWayReplace(s1: str, s2: str) -> bool:
    flag = False
    for i in range(len(s1)):
        if s1[i] != s2[i]:
            if flag:
                return False
            else:
                flag = True
    return True

#insert a charater to s1 in order to make s2
def OneWayInsert(s1: str, s2: str) -> bool:
    index1 = 0
    index2 = 0
    while index1 < len(s1) and index2 < len(s2):
        if s1[index1] != s2[index2]:
            if index1 != index2:
                return False
            index2 += 1
        index1 += 1
        index2 += 1
    return True


def OneWay(s1: str, s2: str) -> bool:
    if len(s1) == len(s2):
        return OneWayReplace(s1 ,s2)
    elif len(s1) + 1 == len(s2):
        return OneWayInsert(s1, s2)
    elif len(s2) + 1 == len(s1):
        return OneWayInsert(s2, s1)
    else:
        return False

#print(OneWay("aaadfg","aabdfg"))

"""
1.6 String Compression:

Implement a method to perform basic string compression using the counts of repeated charaters.
For example, the string aabcccccaaa would become a2b1c5a3. 
If the "compressed" string would not become smaller thatn the original string, your method should return the original string.
You can assume the strig has only uppercase and lowercase letters (a-z).
"""

def StringCompression(input: str) -> str:
    res = ""
    cnt = 0
    for i in range(len(input)):
        cnt += 1
        if i+1 >= len(input) or input[i] != input[i+1]:
            res += input[i] + str(cnt)
            cnt = 0
    if len(res) < len(input):
        return res
    else:
        return input


"""
1.7 Rotate Matrix:

Given an image represented by an NxN matrix, where each pixel in the image is 4 bytes, 
write a method to rotate the image by 90 degrees. Can you do this in place?
"""

def RotateMatrix(matrix: list) -> bool:
    if len(matrix) == 0 or len(matrix) != len(matrix[0]):
        return False
    n = len(matrix)
    for layer in range(n//2):
        first = layer
        last = n - 1 - layer
        for i in range(first, last):
            offset = i - first
            top = matrix[first][i]  #save top
            #left -> top
            matrix[first][i] = matrix[last - offset][first]
            #bottom -> left
            matrix[last - offset][first] = matrix[last][last - offset]
            #right -> bottom
            matrix[last][last - offset] = matrix[i][last]
            #top -> right
            matrix[i][last] = top
    return True

"""
1.8 Zero Matrix:

Write an algorithm such that if an elemnt in an MxN matrix is 0, its entire row and column are set to 0.
"""

def nullifyRow(matrix, row):
    for i in range(len(matrix[0])):
        matrix[row][i] = 0

def nullifyColumn(matrix, column):
    for i in range(len(matrix)):
        matrix[i][column] = 0

def ZeroMatrix(matrix: list) -> None:
    row = []
    column = []
    for i in range(len(matrix)):
        for j in range(len(matrix[0])):
            if matrix[i][j] == 0:
                row.append(i)
                column.append(j)
    for i in row:
        nullifyRow(matrix, i)
    for i in column:
        nullifyColumn(matrix, i)


"""
1.9 String Rotation:

Assume you have a method isSubstring which checks if one word is a substring of another. Given two strings, s1 and s2,
write code to check if s2 is a rotation of s1 using only one call to isSubstring
(e.g., "waterbottle" is a rotation of "erbottlewat").
"""

def StringRotation(s1: str, s2: str) -> bool:
    if len(s1) == len(s2) and len(s1) > 0:
        s1s1 = s1+s1
        return isSubstring(s1s1, s2)
    return False

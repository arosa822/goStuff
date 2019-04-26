# HW - Lesson 1

*   Write a program to take a list of files from from the command line 
and count the words in them.  A word is anything seperated by whitespace (so
detached punctuation will count)

You can test our first program with WC:\n
`$ wc testing.txt`

*   Then Crate an option to count only unique words. 

you can test this with: \n
`awk {for (i=1;i<=NF;i++) {print $i}}' testing.txt|sort|uniq|wc`

*   You may want to use iouti.ReadFile and strings.fields

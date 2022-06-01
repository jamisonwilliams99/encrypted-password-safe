# Encrypted Password Safe
- A CLI password safe that allows users to store and encrypt their passwords using a custom encryption algorithm.

## User Notes

### To create the executable:
```
    go build -o pws.exe
```
- In order for this to work, Go must be installed on your computer.
- the -o flag allows you to name the .exe to whatever you want (in this case, pws.exe)
- - omitting the -o flag will cause the .exe to have the name of the cwd

### Adding a password to the password safe:
```
    pws add <password> <encryption key> <what the password is used for>
```
- Please note that due to the current operations of the encryption algorithm, the password MUST be 16 characters and the key MUST be 8 characters
- - I plan to fix these limitations in future versions

Example:
```
    pws add passwordpassword 5A8D91A5 Facebook
```
- the password "passwordpassword" will be encrypted with the key "5A8D91A5", where the encrpyted password will be stored in the database
- "Facebook" is just an additional note that will be stored alongside the password that will let the user know what the password is for.

#### Adding a password using the "show" flag:
- adding the flag "show=true" to the flag command will show the ecrypted password that is stored:
```
> pws add passwordpassword 5A8D91A5 Facebook --show=true

passwordpassword   ->   1"443+. !↕$$,$'↓

Added a password to the password safe
```

### Listing the passwords that are currently stored in the password safe:
- This will list the ID of every password that is stored in the database along with what the password is used for

```
> pws list
1. Facebook
2. Twitter
3. Reddit
```
- The ID is the unique identifier for each password in the database.
- - Used to retrieve the password after it is stored

### Retrieving a password from the password safe
```
    pws get <password ID> <encryption key>
```

- Command used to retrieve and decrypt and encrypted password that is stored in the database

```
> pws get 1 5A8D91A5
requested password: passwordpassword
```

- Note that even if the incorrect key is passed in, the password will still be decrypted, but with incorrect results:

```
> pws get 1 3FE4!291
requested password: k\nn⌂wzlo`rr{svh
```

#### Retrieving a password using the "show" flag:
- Like with the add command, the "show" flag can be used with the get command to show the encrpyted password:

```
> pws get 1 5A8D91A5 --show=true

1"443+. !↕$$,$'↓   ->   passwordpassword

requested password: passwordpassword
```

### Future commands (Not yet implemented):
- I plan to add the following commands in future versions:
- - remove: removes a password from the password safe
- - update: updates a password that is stored in the password safe (used if you changed the password to the corresponding account)

### The Encrpytion algorithm:

The encryption/decryption algorithm uses concepts from the very simple Caesar Cypher.

- The 16-character password is split into 4 character segments
- The key is split into 2 character segments
- Each segment of the password is mapped to one of the key segments.
- Each character in the key is used to determine how the password segment will be ciphered. 
- - The first character is used to determine the direction of the cypher by first converting the character to its corresponding ASCII integer. The LSB of the binary form determines the direction. If the LSB is 0, the each character in the password segment will be shifted left. If the LSB is 1, they will be shifted right.
- - The second character is used to determine the how much the characters in the password segment will be shifted by. This shifting occurs with respect to the ASCII table, and more specifically, the integer values of each ASCII character. The ASCII table has 128 character, with integer IDS in the range (0-127). The Characters in the password segment will be shifted along this table in the direction and by the amount specified by the key segment. If one of the two boundaries (0 or 127) are exceeded by this shift, it will simply wrap around to the other side.

The ASCII table: https://www.asciitable.com/

**Example**

Consider the following password and key:
```
passwordpassword    5A8D91A5
```

The password and key are segmented and mapped as follows:
pass -> 5A
word -> 8D
pass -> 91
word -> A5

For the sake of this example, I will just cover the first segment since the others follow a similar scheme.

First, the key segment, 5A, must be decoded by first converting the characters to their ASCII ID integers.

5 -> 53
A -> 65

To determine the direction, convert the integer ID of the first character (53) to binary:

53 -> 11010**1**

The LSB of this binary string is **1**, so this means that each character in this password segment will be shifted to the **right**

The shift amount is simply derived from the integer ID of the second character (65). This means that each character in the password segment will be shifted to the **right** by **65** units on the ASCII table.

The result:
pass -> 1"44

This process is repeated for each password, key segment pair to fully encrpyt the password.

### Decrypting
The decryption process is nearly identical. The same key is used on the encrypted password, except the shift direction scheme is reversed. In the decrpytion algorithm, a 0 for the LSB of the direction character in the key segment corresponds to a right shift, and a 1 corresponds to a left shift.

**Note**: This encryption algorithm is likely not very secure and was created as a fun project to experiment with cryptography. I would not recommend using it for encrypting any really sensitive data (like a password). A more secure algorithm, like AES, could easily be dropped in place of the current algorithm to make this applicaton useable for storing sensitive data. This functionality may be added in future versions.



alphabet = 'abcdefghijklmnopqrstuvwxyz'
ALPHABET = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'

# Your string to brute force "Ugew gnwj zwjw Oslkgf"    
ciphertext = "Ugew gnwj zwjw Oslkgf"

def decrypt(n, ciphertext):
    result = ''
    for l in ciphertext:
        try:
                if l.isupper():
                    index= ALPHABET.index(l)
                    i = (index - n) % 26
                    result += ALPHABET[i]

                else:
                    index = alphabet.index(l)
                    i = (index - n) % 26
                    result += alphabet[i]
        except ValueError:
            result += l
    return result

#brute force while for loop testing every shift combo
for key in range(26):
    dec = decrypt(key, ciphertext)
    print("%d . %s" % (key,dec))
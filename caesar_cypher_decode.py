def decipher_char(character):
    base = ord('a')
    return chr((ord(character) - base - key_number) % 26 + base)

ciphertext = input("Enter a ciphertext: ").lower()
key_number = int(input("Enter a key number: "))

plaintext = ""

print("\nDecrypting...\n")

for char in ciphertext:
    if char.isalpha():
        plaintext += decipher_char(char)
    else:
        plaintext += char

print("Plaintext:", plaintext)

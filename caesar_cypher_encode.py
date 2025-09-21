def cipher_char(character):
    base = ord('a')
    return chr((ord(character) - base + key_number) % 26 + base)

plaintext = input("Enter a plaintext: ").lower()
key_number = int(input("Enter a key number: "))

ciphertext = ""

print("\nEncrypting...\n")

for char in plaintext:
    if char.isalpha():
        ciphertext += cipher_char(char)
    else:
        ciphertext += char

print("Ciphertext:", ciphertext)

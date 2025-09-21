def decipher_char(character, key_number):
    base = ord('a')
    return chr((ord(character) - base - key_number) % 26 + base)

def decode(ciphertext, key_number):
    plaintext = ""
    for char in ciphertext:
        if char.isalpha():
            plaintext += decipher_char(char, key_number)
        else:
            plaintext += char
    return plaintext

def brute_force_caesar_cipher_decode(ciphertext):
    print("\nBrute forcing...\n")
    for key_number in range(1, 26):  # 1–25
        plaintext = decode(ciphertext, key_number)
        print(f"Key {key_number}: {plaintext}")

ciphertext = input("Enter a ciphertext: ").lower()

print("\nOptions:")
print("1. Decode with key")
print("2. Brute force all keys")

choice = input("Choose option (1/2): ")

if choice == "1":
    key_number = int(input("Enter a key number: "))
    print("\nDecrypting...\n")
    print("Plaintext:", decode(ciphertext, key_number))
elif choice == "2":
    brute_force_caesar_cipher_decode(ciphertext)
else:
    print("Invalid choice.")

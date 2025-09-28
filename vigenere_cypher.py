ALPHABET = "abcdefghijklmnopqrstuvwxyz0123456789"
ALPHABET_LEN = len(ALPHABET)

def vigenere(text: str, key: str, decrypt: bool = False) -> str:
    result = []
    key = key.lower()
    key_i = 0  # index по ключу

    for ch in text:
        if ch.lower() in ALPHABET:
            ti = ALPHABET.index(ch.lower())
            ki = ALPHABET.index(key[key_i % len(key)])

            if decrypt:
                new_i = (ti - ki) % ALPHABET_LEN
            else:
                new_i = (ti + ki) % ALPHABET_LEN

            result.append(ALPHABET[new_i])
            key_i += 1
        else:
            result.append(ch)
    return "".join(result)

while True:
    print("\n=== Шифр Віженера ===")
    print("1 - Шифрувати")
    print("2 - Розшифрувати")
    choice = input("Ваш вибір: ").strip()

    if choice not in ["1", "2"]:
        print("Невірний вибір!")
        continue

    key = input("Введіть ключ: ").strip()
    text = input("Введіть текст: ").strip()

    if choice == "1":
        result = vigenere(text, key)
        print("Зашифрований текст:", result)
    else:
        result = vigenere(text, key, decrypt=True)
        print("Розшифрований текст:", result)
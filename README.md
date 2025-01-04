# Passphrase Generator
This is a beginner-level passphrase generator built in Go as a learning project for a first contact to the language. The tool generates cryptographically secure passphrases.

## Usage
### **Command-Line Flags**

- `-list`:
  - Choose a wordlist to generate the passphrase. Available options:
    - `eff-long`
    - `bip39`
    - `bip39-pt`
  - Default: `eff-long`.

- `-len`:
  - Specify the number of words in the passphrase.
  - Default: `10`.

### **Example**
Generate a passphrase with `bip39` wordlist and `5` words:
```bash
passphrase -list bip39 -len 5
```

**Output**:
```
vote random admit aware situate (55.00 bits)
```

> `55.00 bits` means 55 bits of entropy.

---

## Wordlists
The following wordlists are available:

- **[EFF Long](https://www.eff.org/files/2016/07/18/eff_large_wordlist.txt)**: Based on the Electronic Frontier Foundation's long wordlist.
- **[BIP39](https://github.com/bitcoin/bips/blob/master/bip-0039/english.txt)**: Standard wordlist used in cryptocurrency wallets for mnemonic phrases.
- **[BIP39-PT](https://github.com/bitcoin/bips/blob/master/bip-0039/portuguese.txt)**: Portuguese version of the BIP39 wordlist.

---

## Installation
1. **Build the Executable**:
   Clone this repository and build the executable using `go build`:
   ```bash
   go build
   ```

2. **Run the Tool**:
   Execute the binary from the terminal:
   ```bash
   ./passphrase.exe
   ```
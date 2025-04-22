# uuid

A simple CLI tool that generates a v4 UUID, prints it to the console, and automatically copies it to your clipboard.

---

## 📋 Prerequisites

- Go 1.18+ installed  
- A `/usr/local/bin/` directory (or equivalent) in your PATH

---

## 📦 Features

- **Secure UUID Generation**  
  Uses `github.com/google/uuid` to generate a standards-compliant v4 UUID.  
- **Automatic Clipboard Copy**  
  Instantly copies the new UUID to your clipboard via `github.com/atotto/clipboard`.

---

## ⚙️ Installation

1. **Clone the repo**  
   ```bash
   git clone https://github.com/Plasmatech-Studios/uuid.git
   cd uuid
   ```

2. **Build**  
   ```bash
   go build -o uuid main.go
   ```

3. **(Optional) Install globally**  
   ```bash
   sudo mv -f uuid /usr/local/bin/uuid
   ```

> You can also install directly with:  
> ```bash
> go install github.com/Plasmatech-Studios/uuid@latest
> ```

---

## 🚀 Usage

```bash
# Generate a UUID and copy it to your clipboard
uuid
```

**What happens:**
- The new UUID is printed to your terminal.
- It’s already in your clipboard—just press ⌘V / Ctrl+V to paste.

---

## 🤝 Contributing

PRs, stars ❤️, and issues are all welcome!

---

## 📜 License

This project is licensed under the MIT License.  

See [LICENSE](LICENSE) for details.
---


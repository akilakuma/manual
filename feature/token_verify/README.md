# VGA Code Checker

A simple Golang utility to verify 12-character tokens based on a hash pattern.  
Useful for access validation, license keys, or lightweight DRM schemes.

---

## 🔍 What It Does

Given a 12-character `code`, the first 4 characters must match specific positions in a `pattern` string.  
The remaining 8 characters are interpreted as 4 index pairs, used to extract characters from `pattern`.  
If all 4 extracted characters match the front 4 characters, the code is considered valid.

---

## 📁 Project Structure

```
refactored_vga_checker/
├── main.go                 # Entry point example
├── vga/
│   ├── checker.go          # Core validation logic
│   └── checker_test.go     # Unit and benchmark tests
└── README.md               # This documentation
```

---

## 🧪 Example

```go
code := "ddea02631215"
pattern := "67d8309e4026e15a7a9fa79a0a33ed01e490ff15bea67e8dd474e21e93c5d7cd"

if vga.Check(code, pattern) {
    fmt.Println("✅ Code is valid!")
} else {
    fmt.Println("❌ Code is invalid.")
}
```
```
步驟 1：分割 code
strFront := code[:4]   // 前 4 碼：ddea
strBack := code[4:]    // 後 8 碼：02631215

步驟 2：把後 8 碼每 2 碼轉換為整數索引
// 分成 4 組兩位數
"02" → 2  
"63" → 63  
"12" → 12  
"15" → 15

indexArr = [2, 63, 12, 15]

步驟 3：取出 pattern 中對應位置的字元
從 pattern 取出：
pattern[2]  = d  
pattern[63] = e  
pattern[12] = e  
pattern[15] = a


步驟 4：比對這些字元與 strFront 的每一位
strFront    = d  d  e  a
pattern取出 = d  e  e  a
               ↑  ↑  ↑  ↑
              OK NOK OK OK

```



---

## 🧪 Benchmark

Run:

```bash
go test -bench=. ./vga
```

---

## 📦 Installation

Clone and run:

```bash
go run main.go
```

---

## 📄 License

MIT License

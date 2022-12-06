# Xoodyak Tools
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/pedroalbanese/xoodyak/blob/master/LICENSE.md) 
[![GoDoc](https://godoc.org/github.com/pedroalbanese/xoodyak?status.png)](http://godoc.org/github.com/pedroalbanese/xoodyak)
[![Go Report Card](https://goreportcard.com/badge/github.com/pedroalbanese/xoodyak)](https://goreportcard.com/report/github.com/pedroalbanese/xoodyak)

 Xoodyak is a lightweight, versatile, cryptographic scheme suitable in constrained environments. It can be used for hashing, encryption, MAC computation and authenticated encryption.

Xoodyak builds upon the Xoodoo permutations and the duplex construction. 

### Xoodyak Encryption Tool
<pre>Usage of xoodyak:
xoodyak [-d] -p "pass" [-i N] [-s "salt"] -f <file.ext>
  -a string
        Additional Associated data.
  -d    Decrypt instead Encrypt.
  -f string
        Target file. ('-' for STDIN)
  -i int
        Iterations. (for PBKDF2) (default 1024)
  -k string
        128-bit key to Encrypt/Decrypt.
  -p string
        PBKDF2.
  -r    Generate random 128-bit cryptographic key.
  -s string
        Salt. (for PBKDF2)</pre>

### Xoodyak Recursive Hasher
<pre>Usage of xoodyaksum:
xoodyaksum [-c &lt;hash.ext&gt;] [-r] &lt;file.ext&gt;
  -c string
        Check hashsum file.
  -r    Process directories recursively.</pre></H3>

### Xoodyak MAC
<pre>Usage of xoodyakmac:
xoodyakmac [-k &lt;secret&gt;] -f &lt;file.ext&gt;
  -f string
        Target file. ('-' for STDIN)
  -k string
        Secret key.</pre>

## License

This project is licensed under the ISC License.

##### Industrial-Grade Reliability. Copyright (c) 2020-2023 ALBANESE Research Lab.

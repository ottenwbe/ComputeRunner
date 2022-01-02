# ComputeRunner

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/ottenwbe/computerunner/blob/main/LICENSE)

This is a small, minimalistic distributed middleware to execute code. It is inspired by my experience with FaaS, PaaS, and stream processing.

Note, this is a small testing playground for me to test out things in distributed middelwares and PaaS/FaaS computing platforms.

## How to run code on the platform

```
curl -d '{"code":"console.log(\"Hello World\");"}' -X POST localhost:8080/runtime  -H "Content-Type: application/json" -v

curl -d '{"code":"node(\"A\", 50*50 "}' -X POST localhost:8080/runtime  -H "Content-Type: application/json" -v  
```

## Disclaimer

I created this project for the purpose of educating myself and personal use. If you are interested in the outcome, feel free to contribute; this work is published under the MIT license.



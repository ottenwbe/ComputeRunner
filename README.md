# ComputeRunner

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/ottenwbe/computerunner/blob/main/LICENSE)

This is a small, minimalistic PaaS/FaaS middleware to execute (JS) code.

The purpose of this project is to provide a small testing playground to test out things in distributed middelwares and PaaS/FaaS computing platforms.

## How to run code on the platform

```
curl -d '{"code":"console.log(\"Hello World\");"}' -X POST localhost:8080/runtime  -H "Content-Type: application/json" -v
```

## Disclaimer

I created this project for the purpose of educating myself and personal use. If you are interested in the outcome, feel free to contribute; this work is published under the MIT license.



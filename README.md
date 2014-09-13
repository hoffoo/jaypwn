jaypwn
======

Make json pretty

```sh
# example with consul
$ curl http://localhost:8500/v1/kv/somejson -s | jaypwn
[
    {
        "CreateIndex": 16,
        "Flags": 0,
        "Key": "somejson",
        "LockIndex": 0,
        "ModifyIndex": 19,
        "Value": "c3R1ZmY="
    }
]
```

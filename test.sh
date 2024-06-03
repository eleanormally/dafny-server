time curl -s -X POST -H "Content-Type: application/json" --data '{
    "requester": "postman",
    "files": [
        {
            "name": "test.dfy",
            "content": "method sumn(n: int) returns (t: int) \\n    requires n >= 0\\n    ensures t == n * (n + 1) /2\\n{\n    var i := 0;\\n    t := 0;\\n    while( i < n ) \\n        invariant 0 <= i <= n\\n        invariant t == i*(i+1)/2\\n    {\\n        i := i+1;\\n        t := t+i;\\n    }\\n    return t;\\n}\\n"
        }
    ]
}' localhost/compile > /tmp/tmp.txt

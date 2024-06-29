# Dafny Server
Single dockerfile to verify arbitrary Dafny code.

# Endpoints
## `/health` (GET)
Returns the current queue length of compilations. Can be used for health detection and load balancing in k8s.
## `/compile` (POST)
Compiles the given dafny files. Does not give any verification that the service is up until it is finished compiling (long queue can generate wait times as high as 10 seconds)

The body of a compilation request should adhere to the following JSON format (an example can be seen in `demo.js`):
```js
{
    requester: "<identifier of requesting body>", // not currently in use
    files: [
        //list of dafny files, all will be automatically linked and compiled as a unit.
        {
            name: "<filename>.dfy",
            content: "<file contents>"
        }
    ]
}
```

let dafnyCode = `
method sumn(n: int) returns (t: int) 
    requires n >= 0
    ensures t == n * (n + 1) /2
{
    var i := 0;
    t := 0;
    while( i < n ) 
        invariant 0 <= i <= n
        invariant t == i*(i+1)/2
    {
        i := i+1;
        t := t+i;
    }
    return t;
}
`

let requestObj = {
  requester: "demo-tester",
  files: [
    {
      name: "demo.dfy",
      content: dafnyCode
    }
  ]
}

let get = async () => {
  let response = await fetch("http://localhost/compile", {
    body: JSON.stringify(requestObj),
    method: "POST",
  })
  console.log(await response.text())
}
get()

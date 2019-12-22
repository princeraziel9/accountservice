I've been following along a great resource which explains microservices in the context of go. I'm having some problems running my web app in a docker container as I'm new to docker. I could run the web server outside of docker. I can successfully buld the docker file but when I run the container it says the file isn't found.

$ go version
    go version go1.12.2 gccgo (GCC) 9.2.0 linux/amd64

I do:
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o accountservice-linux-amd64
    
To compile the code and generate a binary called 'accountservice-linux-amd64'

I then successfully build following Dockerfile:

    FROM iron/base

    EXPOSE 6767
    ADD ./accountservice-linux-amd64 /
    ENTRYPOINT ["./accountservice-linux-amd64"]
    
With:
    $sudo docker build -t raziel9/accountserver .
    
When I run it with:
    $sudo docker run --rm raziel9/accountserver
    
Error message:    
    docker: Error response from daemon: OCI runtime create failed: container_linux.go:346: starting container process caused "exec: \"./accountservice-linux-amd64\": stat ./accountservice-linux-amd64: no such file or directory": unknown.

What gives?
    

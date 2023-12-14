# k6-terminus
K6 extension to execute the terminus commands in K6 test scripts

### Install K6:
    go install go.k6.io/xk6/cmd/xk6@latest

### Compile the binary with the extension:
    xk6 build --with k6-terminus=.

### Compile the binary with other k6 extension:
    xk6 build v0.47.0 \
        --with github.com/pantheon-systems/k6-terminus \
        --with github.com/szkiba/xk6-chai

### Execute the test: 
    ./k6 run test.js

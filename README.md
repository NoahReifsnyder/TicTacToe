# Simple Webapp Tutorial (Golang Version)

This tutorial relates to creating a simple webapp using the Go language

## Getting Started: Setting the GOPATH
Naturally, you will need to install Go.  If you've never used Go before, it's 
important to understand that the Go tools expect you to work in a file tree that 
is structured a certain way.  In particular, if we are building a program called
"webapp", and the root of our Go project is "/a/b/c/", then the code for our 
program should be in "/a/b/c/src/webapp".  In addition, for everything to work, 
it is necessary for the GOPATH environment variable to be set to "/a/b/c".

The easiest way to do this is to follow these steps:

1. Create a file in "/a/b/c" called "env.sh"
2. The contents of "/a/b/c/env.sh" should just be one line:

    export GOPATH=`pwd`

3. Whenever you start a new terminal session, "source" env.sh:

    . env.sh
    
If you followed those instructions correctly, then you should be able to type

    go build webapp
    
From the "/a/b/c" folder.  The result should be an executable in "/a/b/c"
called "webapp" or "webapp.exe".

## Getting Started: Configuration
The webapp needs to know two things: where is the folder whose files should
be served via the "public/" route, and what port number should be used for 
the server.  The easiest way to provide this information is via a config.json 
file, stored in "/a/b/c":

    {
        "AppPort" : "8080",
        "FilePath" : "files"
    }

This says that files in the "files" subdirectory will be served, and that the 
server will run on port 8080.

Note: you can use any folder that you would like, but be sure it exists.  In our
case, from "/a/b/c", you should type:

    mkdir files
    cp src/webapp/*.go files
    
This will ensure you have a folder called "files", and that there is some content 
in it.

## Running the App
If you've got a properly built `config.json`, then you can start your web application by
clicking on it, or by typing `./webapp` from the command line.

To see your app in action, you can start by visiting http://localhost:8080 in your web
browser.  It's going to be pretty unexciting.  But if you put files in the 'files' folder of 
your directory tree, you'll be able to find them at http://localhost:8080/public.  You can also 
use `curl` to interact with the data routes:

    curl -X POST http://localhost:8080/data
    curl -X GET  http://localhost:8080/data
 
 ## Next Steps
 This code is intended to be part of a tutorial.  While it is a very basic introduction to
 building web apps and web services, it is very far from complete.  If you are planning to
 do more with web apps, talk to Professor Spear in order to get some advice on next
 steps.
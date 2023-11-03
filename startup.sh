#!bin/bash

api = "api"


case $1 in 
    "build")            
        if [$2 == $api];then 
            exec go build ./go-api/
        elif [$2 == "wasm"];then
            exec go build ./go-wasm/
        fi
        echo "Building api.";;
    "run")      
        echo "Running";;
    "serve")    
        echo "Serving";;
    *)          
        echo "Command not found";;
esac

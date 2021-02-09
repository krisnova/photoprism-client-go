#!/bin/bash

token="c558cccdd25917056e8b7b72a2a3e5f40215d707a6fac1aa"
server="localhost"
port="8080"

function photoget() {
    url="http://${server}:${port}/${1}"
    curl --header "X-Session-Id: ${token}" --header "Content-Type: application/json" ${url}
}

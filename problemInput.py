#!/usr/bin/env python3
import requests
import sys

def get_input(day):
    url = "https://adventofcode.com/2021/day/"+str(day)+"/input"
    headers = {'Cookie': 'session=53616c7465645f5fb0abbb5bbc7c396781eb2cf92514a3c126037b9ea8e9f0b3f3a7a9e7a5ae9d617ee74ac5c067f10e'}
    r = requests.get(url, headers=headers)

    if r.status_code == 200:
        return r.text 
    else:
        sys.exit(f"/api/alerts response: {r.status_code}: {r.reason} \n{r.content}")


#!/usr/bin/env python3
import requests
import sys
import os

sessionToken = os.getenv('AoC_token')

def get_input(day):
    url = "https://adventofcode.com/2021/day/"+str(day)+"/input"
    headers = {'Cookie': 'session='+sessionToken}
    r = requests.get(url, headers=headers)

    if r.status_code == 200:
        return r.text 
    else:
        sys.exit(f"/api/alerts response: {r.status_code}: {r.reason} \n{r.content}")


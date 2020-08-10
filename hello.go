package main
/*
import json
import copy

a = """
reviewer_homework
"""
q = a.split()



s = "POST /_reindex?wait_for_completion=false"
data = {
    "size" : 10000,

    "source": {
        "remote": {
            "host": "http://es-cn-v641dcsmy0011zuhe.elasticsearch.aliyuncs.com:9200",
            "username": "elastic",
            "password": "9xv4uAzrX$GU"
        },
        "index": "ut-v3-2019-37",
        "query": {
            "match_all": {}
            },
        "sort": {
      "@timestamp": "desc"
    }
    },
    "dest": {
       "index": "ut-v3-2019-37"
    }
}

for i in q:


    print(s)
    data["source"]["index"] = i
    data["dest"]["index"] = i


    newdata = copy.deepcopy(data)
    newdata = json.dumps(newdata)

    print(newdata)
    print()

 */
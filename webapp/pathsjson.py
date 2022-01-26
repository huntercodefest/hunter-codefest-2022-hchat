import json
import os

rel = "src/components/CUNY-CLASSES-JSON"
root_dir = os.listdir(rel)

def createJson(name, path):
    return {"name": name, "path": path}

for d in root_dir:
    dir = os.listdir(rel+"/"+d)
    writelist = []
    for files in dir:
        writelist.append(createJson(files.strip(".json"), files))
    out = open(rel+"/"+d+"/"+"_PATHS.json", "w")
    json.dump(writelist, out, indent=4)
    

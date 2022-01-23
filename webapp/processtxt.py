import json
import os

input_file = os.listdir("CUNY-CLASSES-TXT")
store = []

def createJson(school, key, major, desc):
    return {"school": school, "major": major, "room_num": key, "room_desc": desc}

out_folder_name = "CUNY-CLASSES-JSON/"
os.mkdir(out_folder_name)

for file in input_file:
    
    os.mkdir(out_folder_name+(os.path.basename(file).strip(".txt")))
    input_file_path = os.path.join("CUNY-CLASSES-TXT", file)
    with open(input_file_path) as f:
        line = f.readline()
        ind, ind2 = line.index(':'), line.index(',')
        school = file.strip(".txt")
        key = line[:ind]
        major = line[ind+1:ind2].strip()
        desc = line[ind2+1:].strip('"').replace('\"\n', "")

        outfile = major + ".json"
        folder_file_path = out_folder_name + (os.path.basename(file).strip(".txt"))+ '/' + outfile.strip()
        output = open(folder_file_path, "w")
        jobj = createJson(school, key, major, desc)
        store.append(jobj)
        
        
        temp = major

        file_lines = f.readlines()
        for d in file_lines: 
            ind = d.index(':')
            ind2 = d.index(',')
            school = file.strip(".txt")
            key = d[:ind]
            major = d[ind+1:ind2].strip()
            desc = d[ind2+1:].strip('"').replace('\"\n', "")

            if major == temp:
                jobj = createJson(school, key, major, desc)
                store.append(jobj)
                temp = major
                if d == file_lines[-1]:
                    json.dump(store, output, indent=4)
                    output.close()
                    store = []
                    print("Completed: " + temp)   
                
            else:
                json.dump(store, output, indent=4)
                output.close()
                store = []
                print("Completed: " + temp)
                outfile = major + ".json"
                folder_file_path = out_folder_name + (os.path.basename(file).strip(".txt"))+'/'+outfile.strip()
                output = open(folder_file_path, "w")
                jobj = createJson(school, key, major, desc)
                store.append(jobj)

                temp = major
    
            
            

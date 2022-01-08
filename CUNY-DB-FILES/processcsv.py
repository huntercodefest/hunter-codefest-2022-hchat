import os
import shutil

dir="/home/daniel/go/hunter-codefest-2022-hchat/CUNY-DB-FILES"
def processFiles(ext):
    filenames=[]
    for filename in os.listdir(dir):
        if filename.endswith(ext):
            filenames.append(filename)
    return filenames

def deleteFirstLine(filenames):
    for f_name in filenames:
        with open(f_name, 'r') as f:
            n = os.path.splitext(f_name)[0]+".txt"
            write_to = open(n, 'w')
            f.readline()
            write_to.writelines(f.readlines())
def findLargestFile(filenames):
    max = 0
    name = ""
    for filename in filenames:
        f = open(filename, 'r')
        length = len(f.readlines())
        if length > max:
                max = length
                name = filename
        f.close()
    return name, max


def countLines(filenames):
    max = 0
    total = 0
    f_name = ""
    for filename in filenames:
        read_from = open(filename, 'r')
        lines = read_from.readlines()
        s = int(lines[-1].split(":")[0])
        if s > max:
            max = s
            f_name = filename
        total += len(lines)
    return max, total, f_name

def formatLines(filenames):
    count = 0
    for filename in filenames:
        read_from = open(filename, 'r')
        write_to = open(os.path.splitext(filename)[0]+".txt", 'w')
        read_from.readline()
        lines = read_from.readlines()
        for line in lines:
            count+=1
            li = line.split(",")
            li[0] = li[0].replace("\"","")
            write_to.write(str(count)+": " + li[0]+","+li[1])
        write_to.close()
        read_from.close()

def getLines(filename):
    f = open(filename, 'r')
    lines = f.readlines()
    f.close()
    return lines

def main():
    # this function reads in csv files then deletes first line then converts them to correct format
    csvfilenames = processFiles(".csv")
    deleteFirstLine(csvfilenames)
    txtfilenames = processFiles(".txt")
    formatLines(csvfilenames)
        
if __name__ == "__main__":
     main()
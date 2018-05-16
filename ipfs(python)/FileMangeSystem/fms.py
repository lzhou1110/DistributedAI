
import sys
import get
import add
import time

if __name__ == '__main__':

    time_start = time.time()
    argc = len(sys.argv)-1
    if argc < 1 or argc > 3:
        print("Invalid arguments!")
    else:
        if sys.argv[1] == 'add' and argc == 2:
            filename = sys.argv[2]
            add.splitfile(filename)
        elif sys.argv[1] == 'get' and argc > 1:
            filename = sys.argv[2]
            targetDir = ""
            if argc == 3:
                targetDir = sys.argv[3]
            get.get(filename, targetDir)

        elif sys.argv[1] == 'help' and argc == 1:
            print("\nVaild argument: ")
            print("python fms.py add <filename> : add a file to ipfs.                       ----------E.g. python fms.py add PM10Data.zip")
            print("python fms.py get <HashFileName>: get a file to current directory.       ----------E.g. python fms get HashRecord.txt")
            print("python fms.py get <HashFileName> <OutputDir>: get a file to OutputDir.   ----------E.g. python fms get HashRecord.txt D:/pyproject/fms/recover")
            print("python fms.py help")

        else:
            print("Invalid arguments!")
    time_end = time.time()
    print('totally cost', time_end - time_start)
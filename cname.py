import subprocess

a = open("404.txt").read().split('\n')
c = open("possible.txt", "a+")

class color:
    yellow = '\33[33m'
    green = '\33[32m'
    red = '\33[31m'

for domain in a:
    if (domain == ""):
        print(color.red+"No Domain Found For Check CNAME")
        continue

    cmd = subprocess.check_output(['host', domain])

    text = "".join(map(chr, cmd))

    if("is an alias for" in text):
        point = text.split('\n')
        p = point[0].split("alias for")
        print(color.green+"Takeover may be Possible for "+domain)
        c.write(domain+"@"+p[1]+"\n")
    else:
        print(color.red+"Takeover not Possible for "+domain)

# SubSniper
A simple tool created with Golang. 

## Package Ised
- Sublist3r
- SubFinder
- GoLang
- dnsutils

## Installation
```sh
git clone https://github.com/inadislam/SubSniper.git
cd SubSniper
```
Install SubList3r into SubSniper Folder
```sh
git clone https://github.com/aboul3la/Sublist3r.git
cd Sublist3r && pip install -r requirements.txt
```
Install SubFinder
```sh
go get -v github.com/projectdiscovery/subfinder/cmd/subfinder
```

Install DNSUtils
```sh
apt install dnsutils
```

## Usage
Without Building
```sh
go run main.go [site_url]
```

By Building
```sh
go build main.go
./main [site_url]
```

## Example
```sh
go run main.go example.com
```

By building
```sh
./main example.com
```

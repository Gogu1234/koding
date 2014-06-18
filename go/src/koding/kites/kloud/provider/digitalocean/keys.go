package digitalocean

import "text/template"

var t = template.Must(template.New("hosts").Parse(hosts))

const (
	// name of the key saved in DigitalOcean
	keyName = "koding-deployment"

	// RSA key pair that we add to the newly created machine for
	// provisioning.
	publicKey = `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDYFQFq/DEN0B2YbiZqb3jr+iQphLrzW6svvBjQLUXiKA0P0NfgedvNbqqr2WQcQDKqdZQSHJPccfYYvjyy0wEwD7hq8BDkHTv83nMNxJb3hdmo/ibZmGoUBkw3K7E8fzaWzUDDNSlzBk3UrGayaaLxzOw1LhO5XUfesKNWCg4HzdzjjOklNpJ61iQP4u8JRqXJaOV5RPogHYFDlGXPOaBuDxvOZZanEgaKsfFkwEvpU0km5001XVf8spM7o8f2iEalG9CMF1UVk38/BKBngxSLRyYdP/K0ZdRBSq1syKs8/KPrDWQ6eyqG2cW6Zrb8wb2IDg7Na+PfnUlQn9S+jmF9 hello@koding.com`

	privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA2BUBavwxDdAdmG4mam946/okKYS681urL7wY0C1F4igND9DX
4HnbzW6qq9lkHEAyqnWUEhyT3HH2GL48stMBMA+4avAQ5B07/N5zDcSW94XZqP4m
2ZhqFAZMNyuxPH82ls1AwzUpcwZN1Kxmsmmi8czsNS4TuV1H3rCjVgoOB83c44zp
JTaSetYkD+LvCUalyWjleUT6IB2BQ5Rlzzmgbg8bzmWWpxIGirHxZMBL6VNJJudN
NV1X/LKTO6PH9ohGpRvQjBdVFZN/PwSgZ4MUi0cmHT/ytGXUQUqtbMirPPyj6w1k
OnsqhtnFuma2/MG9iA4OzWvj351JUJ/Uvo5hfQIDAQABAoIBAEUz7QFThW5UEbZo
yJrb2pFQylYVsT4RRLED/7pkVXZJt20ySIEW5eNUJwrkk6BMmy1mfKaUSnKWeOeR
vaTT2RSIZFqzSeWUy0p//78QD/1z+7KYut6DZq3FGpGsx5WFdHk1gSFDnGO/SBvt
nvJSzKG+LIQXnq+GBWd3kMgUbkvm9mG2Xi/NTcJYB6foLavA6tbo32CZOQtqvqWt
ysFkxZM6RJuERUqz4KE8auzGcySosNBBW9EKrT6SlVMOW1hcnLMiH46VUrYrmuQe
7dQyvULYmhJ5DWptO820qhgCP7D5+FlIMCrL6Ka5vonbb3sMCKHrYy8SroW2Vt8/
+b3Kre0CgYEA/R4+DKynYSxBNiG53CIP5zSnWGmM2kjV1kyMsZhVj6wmURm1VTB3
0Z5M74dJUcR1HEZCkqiJyknZnRS+wja5xpCnCY6bDJGnbufDXyfKlHS2q/iBUJWj
qvO1tO0H1/LDSvx2JlW9iRtwnBPSH3reJu7MpVTO9M5Cb+uCrJrYzOsCgYEA2orQ
n/jwBESrFDRw2Ghodr5gEfcDZUUXRRdxqj9P/vM2Ax5y8cTw1kY8cETN4XsKbpUq
y7fpcYI1YOD/jtVPa6D9l+Oz6NlrFMvJM6XZrrKKH6EzMBkXV2hvB1i7NM0Wm1Qi
pX5uuTt459deQ0tZvkNwKUw+ElL5x4gVi8psUTcCgYEAnXxbjvc9jTBDwrJpOZXX
3zrbhB7oDEiVA6jNQRJO6f1qObuNH1vwsPOVWtMJw15Anz733NgQI1SfmSR3K89w
9yK5SzD3N3LIgjChVmsjmAqmPYl4q/LuykaoH9H6t1nMOLOrr15Zdx0ji3ipm/yO
jBq4KYYC3j4XCBu4Sjxt99ECgYAv4BsmAG3sMXLdAUP7VJwV70yygs9+nu0jGKiw
6B/JNVSOyFvb9Q3hkw9odNo5XDAD26+9YzPDxweUipXRn48/f8wuOxbNmtuneaFB
LEMMz1YM6c6B6e1AG8O+80RHo+og6wSQBquQ7qNk/rxKt7YYRUPL+ETc1MqdWIWH
0zd5bwKBgGz18tYk9Zv92c0uFUY4wSeXSdaXXU5RxZzS4bHUT5OyDMxn3krvjOLE
bC3K2UUHqakRw36hwLLadePZFUqC9AY8ITeZ43vIKhBKBQj3lhfSuc+vyHwCSuc7
6+smUFdsRVz6i3SWFOjeBYMNuDR/xgBEh9teQfn4/9aiEi2QAriN
-----END RSA PRIVATE KEY-----`

	// default ubuntu /etc/hosts file which is used to replace with our custom
	// hostname later
	hosts = `127.0.0.1       localhost
127.0.1.1       {{.}} {{.}}

# The following lines are desirable for IPv6 capable hosts
::1     ip6-localhost ip6-loopback
fe00::0 ip6-localnet
ff00::0 ip6-mcastprefix
ff02::1 ip6-allnodes
ff02::2 ip6-allrouters`
)

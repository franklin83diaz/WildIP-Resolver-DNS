# WildIP-Resolver-DNS

a Simple server DNS for resolver ip from a subdomain with formats:  0-0-0-0.mydomain.example



## Example

10.5.3.10.mydomain.example => 10.5.3.10

```
# DEBUG=true /go run ./cmd

$ nslookup  10.0.0.10.ip.mydns.com 127.0.0.5         
Server:		127.0.0.5
Address:	127.0.0.5#53

Name:	10.0.0.10.ip.mydns.com
Address: 10.0.0.10



```


## Authors

- [@Fdiaz](https://github.com/franklin83diaz/)


## License

[MIT](https://choosealicense.com/licenses/mit/)
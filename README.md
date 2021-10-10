# Http-proxy-server
Http proxy server and a simple vulnerability scanner based on it
- http proxy 
- requests repeater
- xss checker
#### Ports:
Proxy is working on :8080  
Web-api is on :8000

### Build and run
```
sudo docker build -t proxy .
sudo docker run -p 8080:8080 -p 8000:8000 -t proxy
```

# Default Backend of Ingress in Kubernetes
[![Build](https://github.com/newbiediver/default-backend/actions/workflows/build.yml/badge.svg)](https://github.com/newbiediver/default-backend/actions/workflows/build.yml)
## 1. Summary
* 초 심플한 404 page 를 위한 웹 서비스
* Ingress 호출에 대해 404 page 를 대체하기 위한 Default Backend
* 특별한 옵션 및 환경설정 없이 있는 그대로 사용 가능

## 2. Screenshot
![](screenshot/screenshot.png)

## 3. Usage
* Deployment
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: default-backend
  namespace: default
spec:
  replicas: 1
  selector:
    matchLabels:
      app: default-backend
  template:
    metadata:
      labels:
        app: default-backend
    spec:
      containers:
      - name: default-backend
        imagePullPolicy: Always
        image: ghcr.io/newbiediver/default-backend:latest
        ports:
        - containerPort: 8000
```
* Service
```yaml
apiVersion: v1
kind: Service
metadata:
  name: default-backend
  namespace: default
spec:
  selector:
    app: default-backend
  ports:
  - port: 8000
```
* Ingress (Ex> Nginx Ingress Controller)
```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: ingress-example
  namespace: default
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /
spec:
  ingressClassName: nginx
  # 여기에 정의
  defaultBackend:
    service:
      name: default-backend
      port:
        number: 8000
  rules:
  - host: example.com
    http:
      paths:
      - pathType: Prefix
        path: /
        backend:
          service:
            name: example-service
            port: 
              number: 80
```
## 4. External Dependencies
* [gin-gonic/gin](https://github.com/gin-gonic/gin)
* [HTML Template by Colorlib](https://colorlib.com/)
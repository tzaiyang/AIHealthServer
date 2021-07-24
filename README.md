# AIHealth
[![Build Status](https://www.travis-ci.com/tzaiyang/AIHealthServer.svg?branch=master)](https://www.travis-ci.com/tzaiyang/AIHealthServer)

AIHealth is an app recording your health data, including base information(height,weight,ABO,etc.), sports data, medical record, etc. The app will analyse these data to help you improve your body health.

# CI&CD
0. [travis CI](./.travis.yml)  
   Trigger by `git push origin master`, Docker Image generated will pushe to aliyun
1. Docker deployment
    ```bash
    docker pull tzaiyang/aihealth:0.1
    docker run --rm --name aihealth -d -p 8080:8080 tzaiyang/aihealth:0.1
    ```

2. Requirements
    - [Gin](https://gin-gonic.com/)  
    - [MongoDB Atlas](https://cloud.mongodb.com/)
    - [mongo-driver](https://docs.mongodb.com/drivers/go/)
    
    
3. Framework maybe to be used
    - elastics

4. API Documents
   - Preview [openapi.yaml](./docs/openapi.yaml) with [OpenAPI Online Editor](https://editor.swagger.io/) or `OpenAPI VS Code pludgin`
## Refernces
https://gin-gonic.com/docs/examples/bind-query-or-post/

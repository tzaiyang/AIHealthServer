# AIHealth
[![Build Status](https://www.travis-ci.com/tzaiyang/AIHealthServer.svg?branch=master)](https://www.travis-ci.com/tzaiyang/AIHealthServer)

AIHealth is an app recording your health data, including base information(height,weight,ABO,etc.), sports data, medical record, etc. The app will analyse these data to help you improve your body health.

# Local Test
export CONFIG_URL="config/config.yaml" && go run cmd/server/main.go

# CI&CD
0. [travis CI](./.travis.yml)  
   Trigger by `git push origin master`, Docker image generated will be pushed to dockerhub.
1. Docker deployment
    ```bash
    docker pull tzaiyang/aihealth:$tag
    docker run --rm --name aihealth -d -p 10086:10086 -v $config_file_paths:/app/config.yaml tzaiyang/aihealth:$tag
    ```

2. Requirements
    - [Gin](https://gin-gonic.com/)  
    - [MongoDB Atlas](https://cloud.mongodb.com/)
    - [mongo-driver](https://docs.mongodb.com/drivers/go/)
    
    
3. Framework maybe to be used
    - elastics

4. API Documents
   - Preview [openapi.yaml](./docs/openapi.yaml) with [OpenAPI Online Editor](https://editor.swagger.io/) or `OpenAPI VS Code pludgin`

## FAQ
1. mongo white ip list:(connection refused)
`sudo mongod --bind_ip_all`  
or edit param in `/etc/mongod.conf` then start mongod with `systemctl start mongod` 
2. mongod start with error "child process failed, exited with error number 1"  
   `sudo chmod 777 /var/log/mongodb/`

3. mongo data export
   ```bash
   # export from mongodb
   $ mongoexport -d AIHealth -c medicals -o AIHealth_medicals.dat
   # import to remote host.
   $ mongoimport -h aiwac.net:27017 -d AIHealth -c medicals AIHealth_medicals.dat
   ```
 
```
docker run -d --hostname my-rabbit --name some-rabbit -p 0.0.0.0:8087:15672 -p 0.0.0.0:5672:5672 rabbitmq:3-management 
```
```
docker run -d --name redis-container -e REDIS_ALLOW_REMOTE_CONNECTIONS=yes -e TZ=Europe/Moscow -p 0.0.0.0:30073:6379 -e REDIS_PASSWORD=cyberpass ubuntu/redis
```
```
docker run -it --rm -e RABBIT_HOST=10.22.22.32 -e RABBIT_PORT=5672 -e RABBIT_USERNAME=guest -e RABBIT_PASSWORD=guest -p 0.0.0.0:8011:8011 publisher:develop
```
```
docker run -it --rm -e RABBIT_HOST=10.22.22.32 -e RABBIT_PORT=5672 -e RABBIT_USERNAME=guest -e RABBIT_PASSWORD=guest -e REDIS_HOST=10.22.22.33 -e REDIS_PORT=30073 -e REDIS_PASSWORD=cyberpass -e REDIS_DB=0 redisoper:develop
```
```
docker run -it --rm -e REDIS_HOST=10.22.22.33 -e REDIS_PORT=30073 -e REDIS_PASSWORD=cyberpass -e REDIS_DB=0 consumer:develop
```
***
***
```
10.22.22.32:15672/#/queues/%2F/publisher
```
```
POST 10.22.22.32:8011/publish/thanks!
```
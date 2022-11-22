DOCKER_IMG1="publisher:develop"

DOCKER_IMG2="consumer:develop"

DOCKER_IMG3="redisoper:develop"

build-img1:
	docker build \
		-t $(DOCKER_IMG1) \
		-f publisher/Dockerfile .


build-img2:
	docker build \
		-t $(DOCKER_IMG2) \
		-f consumer/Dockerfile .

build-img3:
	docker build \
		-t $(DOCKER_IMG3) \
		-f redisoper/Dockerfile .
	

 
.PHONY: build-img1 build-img2 build-img3
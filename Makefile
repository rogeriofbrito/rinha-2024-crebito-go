REGISTRY := rogeriofbrito
IMAGE_TAG := $(shell openssl rand -hex 4)
IMAGE_NAME := rinha-2024-crebito-go

docker:
	docker build -t $(REGISTRY)/$(IMAGE_NAME):$(IMAGE_TAG) .
	docker tag $(REGISTRY)/$(IMAGE_NAME):$(IMAGE_TAG) $(REGISTRY)/$(IMAGE_NAME):$(IMAGE_TAG)
	docker push $(REGISTRY)/$(IMAGE_NAME):$(IMAGE_TAG)

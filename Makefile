build:
	buildctl build \
		--frontend dockerfile.v0 \
		--local context=. \
		--local dockerfile=. \
		--output type=image,name=gcr.io/moonrhythm-containers/cfztna-tls-beacon:latest,push=true

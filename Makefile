build:
	buildctl build \
		--frontend dockerfile.v0 \
		--local context=. \
		--local dockerfile=. \
		--output type=image,name=registry/moonrhythm.io/cfztna-tls-beacon:latest,push=true

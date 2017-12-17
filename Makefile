build:
	docker build -f Dockerfile.build .
	docker run -v ${HOME}/.glide/:/root/.glide/ -v $$(pwd)/release:/mnt/release --rm $$(docker build  -qf Dockerfile.build .)
	docker build -t vxlabs/vault-envexport .

publish::
	./scripts/publish

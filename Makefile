protogen:
	docker run -v ${PWD}:/defs namely/protoc-all -f article/article_service.proto -o article/api -l go
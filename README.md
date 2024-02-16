# Product Search

![Screenshot from 2024-02-12 22-54-34](https://github.com/sog01/repogen/assets/40876952/575f4290-c6ac-43fb-8fbc-e44357e9d582)

The Product Search is a web application designed to demonstrate typical search features for products. This feature is still undergoing development and will be continuously updated. So far, this project has accomplished the following features:

- [x] Search Autocompletes API
- [x] Search Results API
- [x] Search Order By and Filter by Catalog API
- [x] Product Search UI
- [x] Product and Catalog CRU API
- [x] Catalog UI
- [x] URL Shortener for Share Catalog

And for the on-going development would be:

- [ ] Product Tagging API
- [ ] User Management
- [ ] Admin UI

## Programming Stacks

- ⭐ Golang
- ⭐ Htmx
- ⭐ Bootstrap
- ⭐ Elasticsearch
- ⭐ Docker
- ⭐ Javascript
- ⭐ Html

## How to Run

Let's look at Makefile to run a command that we need.

### Run Elasticsearch

Execute `make upES` to turn on our Elasticsearch locally.

### Running CRUD Server

Execute `make run` to run our CRUD Server.

### Stop Elasticsearch

Execute `make downES` to teardown running Elasticsearch.

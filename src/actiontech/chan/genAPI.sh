which yq > /dev/null || (echo "you need download yq from https://github.com/mikefarah/yq" && exit 1)
swagger version | grep "0.19.0" || (echo "you should update go-swagger to version 0.19.0" && exit 1)
mv umc/api/restapi/configure_umc.go umc/api/restapi/configure_umc.go.bak
yq merge umc/api/definitions/*.yml > umc/api/api_defs.yml
rm -rf umc/api/models
rm -rf umc/api/restapi/operations
GOPATH=${GOPATH} swagger generate server -A umc -f umc/api/api_defs.yml -t umc/api --exclude-main || (mv umc/api/restapi/configure_umc.go.bak umc/api/restapi/configure_umc.go; rm -f umc/api/api_defs.yml; exit 1)
./api_validation_generator/gen umc/api/api_defs.yml umc/controllers/authority/apiauthority/api_authority_map.go
rm -f umc/api/api_defs.yml
mv umc/api/restapi/configure_umc.go umc/api/restapi/configure_umc.go.new
mv umc/api/restapi/configure_umc.go.bak umc/api/restapi/configure_umc.go
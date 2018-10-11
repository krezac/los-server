#!/bin/sh
SWAGGER_UI_VERSION='3.19.3'
SWAGGER_UI_ARCHIVE='v'${SWAGGER_UI_VERSION}'.zip'
DOWNLOADED_ARCHIVE='swagger-ui.zip'
UNPACK_FOLDER='swagger-ui'

wget https://github.com/swagger-api/swagger-ui/archive/${SWAGGER_UI_ARCHIVE} -O ${DOWNLOADED_ARCHIVE}
unzip -jo ${DOWNLOADED_ARCHIVE} 'swagger-ui-'${SWAGGER_UI_VERSION}'/dist/*' -d ${UNPACK_FOLDER}
sed -i 's/url: \"https:\/\/petstore.swagger.io\/v2\/swagger.json\"/url:\".\/los-server.yml\"/g' ${UNPACK_FOLDER}/index.html
cp -f swagger/los-server.yml ${UNPACK_FOLDER}/
statik -src swagger-ui -f -p swaggeruistatik
rm -rf ${UNPACK_FOLDER} ${DOWNLOADED_ARCHIVE}
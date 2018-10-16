#!/bin/sh

INPUT_FILE="data.json"
OUTPUT_FILE="scripts/sql/ranges.sql"

wget http://zbranekvalitne.cz/strelecka-mapa/${INPUT_FILE}


sed -i "s/var mapObjects = //" ${INPUT_FILE}
sed -i "s/,];/]/" ${INPUT_FILE}
sed -i "s/'/\"/g" ${INPUT_FILE}
sed -i "s/lat\\:/\"lat\"\\:/g" ${INPUT_FILE}
sed -i "s/lng\\:/\"lng\"\\:/g" ${INPUT_FILE}
sed -i "s/name\\:/\"name\"\\:/g" ${INPUT_FILE}
sed -i "s/cat\\:/\"cat\"\\:/g" ${INPUT_FILE}
sed -i "s/lnk\\:/\"lnk\"\\:/g" ${INPUT_FILE}

sed -i "s/\\\\x20/ /g" ${INPUT_FILE}
sed -i "s/\\\\x2D/-/g" ${INPUT_FILE}
sed -i "s/\\\\x2F/\\//g" ${INPUT_FILE}

sed -i "s/\\\\x26/\\&/g" ${INPUT_FILE}
sed -i "s/\\\\x28/\\(/g" ${INPUT_FILE}
sed -i "s/\\\\x29/\\)/g" ${INPUT_FILE}
sed -i "s/\\\\x2B/\\+/g" ${INPUT_FILE}

sed -i "s/\\\\u2013/-/g" ${INPUT_FILE}

sed -i "s/\\\\u011B/ě/g" ${INPUT_FILE}
sed -i "s/\\\\u0161/š/g" ${INPUT_FILE}
sed -i "s/\\\\u010D/č/g" ${INPUT_FILE}

sed -i "s/\\\\u0159/ř/g" ${INPUT_FILE}
sed -i "s/\\\\u017E/ž/g" ${INPUT_FILE}
sed -i "s/\\\\u00FD/ý/g" ${INPUT_FILE}

sed -i "s/\\\\u00E1/á/g" ${INPUT_FILE}
sed -i "s/\\\\u00ED/í/g" ${INPUT_FILE}
sed -i "s/\\\\u00E9/é/g" ${INPUT_FILE}
sed -i "s/\\\\u0148/ň/g" ${INPUT_FILE}
sed -i "s/\\\\u010F/ď/g" ${INPUT_FILE}
sed -i "s/\\\\u0165/ť/g" ${INPUT_FILE}

sed -i "s/\\\\u0158/Ř/g" ${INPUT_FILE}
sed -i "s/\\\\u0160/Š/g" ${INPUT_FILE}
sed -i "s/\\\\u017D/Ž/g" ${INPUT_FILE}
sed -i "s/\\\\u010C/Č/g" ${INPUT_FILE}
sed -i "s/\\\\u00DA/Ú/g" ${INPUT_FILE}
sed -i "s/\\\\u016F/ů/g" ${INPUT_FILE}

# filter ranges (get rid of shops)
#jq -r -s ". | select(.cat==\"r\") | map(\"prdel\")" ranges.json
#jq  '.[] | select(.cat=="r") | ["INSERT INTO ranges (name, latitude, longitude) VALUES (", .name, ", ", .lat, ", ", .lng, ", ", .lnk, ")"]' ranges.json
jq -r --arg q "'" '.[] | select(.cat=="r") | ["INSERT INTO ranges (NAME, LATITUDE, LONGITUDE, URL) VALUES (", ($q), .name, ($q), ", ", (.lat|tostring), ", ", (.lng|tostring), ", ", ($q), "https://zbranekvalitne.cz", .lnk, ($q), ");"] | add' ${INPUT_FILE} > ${OUTPUT_FILE}

rm -f ${INPUT_FILE}
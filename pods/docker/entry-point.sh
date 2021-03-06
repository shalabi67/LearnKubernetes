STAGE=$1
sed -e "s/{ENV}/${STAGE}/" MANIFEST-temp.MF > MANIFEST.MF
sed -e "s/{ENV}/${STAGE}/" pom-temp.properties > pom.properties

cat MANIFEST.MF

cp MANIFEST.MF /data/zalando/processes/p9999/ROOT/META-INF/MANIFEST.MF
cp pom.properties /META-INF/maven/de.zalando.zalos/frontend/pom.properties

./catalina.sh run

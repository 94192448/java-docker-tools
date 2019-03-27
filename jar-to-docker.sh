echo 'Starting build docker..'
pwd

IMAGES='192.168.251.157:1180/apps/'$1:$2
echo $IMAGES
docker build -t=$IMAGES .
echo 'Starting push to harbor..'
sleep 2s
docker login -u ci -p Harbor12345 192.168.251.157:1180
docker push $IMAGES
echo 'Success push to harbor...'
sleep 2s
echo 'Rest env...'
docker rmi $IMAGES
#rm -rf *.jar
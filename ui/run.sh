docker build -t atlas-admin-ui  --build-arg DATABASE_URL="mysql://root:{password}@host.docker.internal/test"  .
docker tag atlas-admin-ui 192.168.2.11:5000/atlas-admin-ui:latest
docker image push  192.168.2.11:5000/atlas-admin-ui:latest

docker run -p 3000:3000 atlas-admin-ui

# meli-challenge

To run the golang api

    go run main.go

golang sql drivers: https://go.dev/wiki/SQLDrivers

golang sqlite3
go get -u -v github.com/mattn/go-sqlite3

    go get github.com/mattn/go-sqlite3

Run the create_db_data.py script to initialize the database

python -m venv meli_venv
meli_venv\Scripts\activate
pip freeze > requirements.txt

CGO_ENABLED=1 go build
go build
.\meli-go.exe

## Deploy

sls deploy --config serverless-service.yml --stage dev --region us-east-1 --app meli-challenge --verbose

...

golang + sqlite

    https://medium.com/@orlmonteverde/api-rest-con-go-golang-y-sqlite3-e378af30719c
    https://www.youtube.com/watch?v=YpDVQC8hfik

cgo_enabled para aws linux

    https://www.reddit.com/r/golang/comments/ka5ipb/gosqlite3_requires_cgo_to_work/

    sudo apt-get update
    sudo apt-get install build-essential libsqlite3-dev




            /home/meli-challenge/golang/meli-go


            yum update -y aws-cfn-bootstrap

correr el script de crear base de datos al inicio de la ec2 instance

sudo python3 /home/meli-challenge/flask/run.py

sudo /home/meli-challenge/golang/meli-go

...

            cd /home/meli-challenge/golang && sudo /home/meli-challenge/golang/meli-go &

nginx -g "daemon on;" && gunicorn --bind 0.0.0.0:8000 run:app --log-level=debug --workers=3

gunicorn --bind 0.0.0.0:8000 run:app --log-level=debug --workers=3

cd /home/meli-challenge/golang && sudo /home/meli-challenge/golang/meli-go &

sudo /home/meli-challenge/golang/meli-go &

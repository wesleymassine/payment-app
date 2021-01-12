# Payment APP #
Payment APP is an application written in Go for carrying out transactions through a digital account.

## How to Run the Application
To use the API it is necessary to have the [Docker](https://docs.docker.com/get-docker/) installed.

After the installation, just use the command below [seguinte comando](https://docs.docker.com/compose/reference/up/)

```bash
docker-compose up --build
```
This way the API will be available to be used on the port [5000](http://localhost:5000)

Other alternatives to run the application are:

Use the executable in the root folder: ./payment_app
Or the command: go run payment_app
If you have any problems, please talk to me.

## Usage
- Make sure you have Go 1.15+ installed:

```bash
go version  # go version go1.15.4 linux/amd64
```
## Package
I sought to make the most of the native resources of the lineage. I used only the gorilla/mux package to route HTTP requests and URL matching to build web servers with Go (github.com/gorilla/mux) v1.8.0.

## Api Interface
* 1. Account creation 
    * URL: http://localhost:5000/accounts
    * METHOD: POST 
    * REPONSE: CREATED (201)
    * PAYLOAD: (JSON)

    * GET Transactions
    * URL: http://localhost:5000/transactions
    * METHOD: GET 
    * REPONSE: CREATED (200)
    * PAYLOAD: (JSON)

* 2. Transaction authorization 
    * URL: http://localhost:5000/transactions
    * METHOD: POST 
    * REPONSE: CREATED (201)
    * PAYLOAD: (JSON)

    * GET Transactions
    * URL: http://localhost:5000/transactions
    * METHOD: GET 
    * REPONSE: CREATED (200)
    * PAYLOAD: (JSON)

### Executing requests with Postman

[![Run in Postman](https://run.pstmn.io/button.svg)](https://god.postman.co/run-collection/bfa3671453b5f86f9692)

## Design decisions
The MVC standard was adopted, the model represents the application data and the business rules that govern data access and modification. The model maintains the persistent state of the business and provides the controller with the ability to access the application's functionalities encapsulated by the model itself.

This application does not have a database, and its data persistence is through temporary files.

## Files and packages structures
    .
    ├── src                         # Src main file with all source code
        ├── controllers             # Controllers temporary file storage in json format
        ├── file                    # File temporary file storage in json format
        ├── models                  # modela the data and behavior behind the business process
        ├── repositories            # Repositories persistence of data
        ├── router                  # Router manages HTTP requests
        ├── utils                   # Utils generic and auxiliary functions
    ├── .env                        # File for environment variable settings
    ├── payment_app              # Payment_app application executable file
    ├── docker-compose.yaml
    ├── Dockfile
    ├── go.mod
    ├── go.sum
    ├── go.main
    └── readme.md

## Testing
For the tests, Golang's native library was used.

    The tests cover only the model and utils where the business rules are found.
    If you miss something or a greater coverage, please do not hesitate, I will be happy to receive your requests.

Execute the tests with the command below:

    * In the root folder run: go test -v. / ...

    * Enter the models or utils folder to view the result in the terminal: go tool cover --func = cover.txt

    * To view the result of the package functions in the browser (beautiful heh): go tool cover --html = cover.txt

## Licença
[MIT](https://choosealicense.com/licenses/mit/)

## Referências ##
* https://golang.org/

# Golang Project

Just for fun!

## Installation

Clone the repository: 
```bash
git clone git@github.com:guther/hostgator-challenge.git
```
Enter the directory and do the command:

```bash
docker-compose up
```
  

## Usage

This will create the necessary containers for the application to work.

At the end, the application will be active at the address [http://localhost:6060](http://localhost:6060), however it is necessary to authenticate beforehand. Otherwise, you will receive error 401 (Not authenticated).

Authenticate sending a POST request to [http://localhost:6060/login](http://localhost:6060/login) along with the user data:

`body: '{" password ":"@#$RF@!718"," username ":" admin "}`

You will receive the token (JWT).
Use this token to perform requests.

## Items

- JWT
- Cache in MySQL
- RESTful
- Golang (using Gin framework)
- Unit Tests
- Database Schema using migrations
- CI Pipeline using Jenkins

Now we will see each item above to understand better how the application was developed. Okay?

# JWT

As I already reported, the application uses JWT to perform authentication.
![JWT](https://raw.githubusercontent.com/guther/hostgator-challenge/dev/images/postmantoken.jpg?token=AEWA2JO67DZSSGUL4TUSFPDABX67E)

The token is set to expire in 48 hours and care has been taken to use the correct HTTP codes, such as the return of the 401 code instead of 403, as it is an authentication, whose authorization is granted upon login.

![JWT2](https://raw.githubusercontent.com/guther/webnotify/master/images/postmanres.jpg) 

# Cache 

I used MySQL. The database schema was made using migration, and although it chose not to use any ORM framework, the database is consistent and fault tolerant. 

![Schema](https://raw.githubusercontent.com/guther/webnotify/master/images/migrations.jpg)

# RESTful

With good practices, I did not use verb names in the routes and used the HTTP codes as desired. I did not have to deal with the PUT and DELETE methods, as the application only has GET and POST, as stated.

# Golang
I used the Gin framework in this application.

# Unit Tests
This one was one of the best parts! Golang helps a lot to create unit tests with the command ``` go test ```.
Jenkins was responsible for running the tests and we can see the Pipeline running.

![pipeline](https://raw.githubusercontent.com/guther/webnotify/master/images/pipeline.jpg)

This job is also not created manually, it is automatic.
I used the Job DSL plugin to do this. See the SDL format, it is another format.

![jsdl](https://raw.githubusercontent.com/guther/webnotify/master/images/jdsl.jpg)

This sdl code creates the **job-generator**, which is just a seed to create the real Jenkins jobs, like the CI job, for example.

# Jenkins

To skip all **Jenkins** configuration steps, I used the **JCasC - Jenkins Configuration as Code** practice and we don't have to go through user and password creation, job creation, plugins, etc. When you access the [http://localhost:8080] (http://localhost:8080) page, you are already dropped into the login screen.
 
![cockerfile](https://raw.githubusercontent.com/guther/webnotify/master/images/dockerfile.jpg)

To achieve this, I used jenkins-plugin-cli directly in Dockerfile in conjunction with the shared casc.yaml configuration for Jenkins' image at the time of running the ```docker-composer``` up command.

![casc](https://raw.githubusercontent.com/guther/webnotify/master/images/casc.yaml.jpg)


Enter the credentials: **admin** and **password**.

You will then see two jobs created at startup.

![Jobs](https://raw.githubusercontent.com/guther/webnotify/master/images/jobs.jpg)

CI is the job that performs the Continuous Integration pipeline with the test steps. It clones the repository, runs the tests and stores the log.

# Flow Diagram for /breeds

![Jobs](https://raw.githubusercontent.com/guther/webnotify/master/images/diagrama_de_fluxo_breeds.jpg)


# Flow Diagram for /login

![Jobs](https://raw.githubusercontent.com/guther/webnotify/master/images/diagrama_de_fluxo_login.jpg)

## License
[MIT](https://choosealicense.com/licenses/mit/)

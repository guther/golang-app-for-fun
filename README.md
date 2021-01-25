# HostGator Challenge

HostGator Latin America

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

## Requisitos de Implementação

- JWT
- Cache com MySQL
- RESTful
- Golang (bonus :-) .. using Gin framework!!! :P
- Private repository on GitHub
- Several commits during the development
- Unit Tests
- Database Schema using migrations
- CI Pipeline using Jenkins

Now we will see each item above to understand better how the application was developed. Okay?

Upon becoming aware of the application's items, I researched how it could facilitate the life of my appraiser, in such a way that when starting the application it was already possible to have everything configured.
And believe me, this application is like that.


# JWT

As I already reported, the application uses JWT to perform authentication.
![JWT](https://raw.githubusercontent.com/guther/hostgator-challenge/dev/images/postmantoken.jpg?token=AEWA2JO67DZSSGUL4TUSFPDABX67E)

The token is set to expire in 48 hours and care has been taken to use the correct HTTP codes, such as the return of the 401 code instead of 403, as it is an authentication, whose authorization is granted upon login.

![JWT2](https://raw.githubusercontent.com/guther/webnotify/master/images/postmanres.jpg) 

# Cache 

I used MySQL as requested. The database schema was made using migration, and although it chose not to use any ORM framework, the database is consistent and fault tolerant. That is: (Now here's a test tip for you to do with my competitors) When the database schema already exists and you change the database password, there is an authentication error, which is the old acquaintance of the developers, huh?! But, what about when you change the database password even before the migration takes place? Hummmm ... it's another mistake. And the application has to be tolerant of this to verify the moment that it can create the database and the cache will start to function.

![Schema](https://raw.githubusercontent.com/guther/webnotify/master/images/migrations.jpg)

# RESTful

With good practices, I did not use verb names in the routes and used the HTTP codes as desired. I did not have to deal with the PUT and DELETE methods, as the application only has GET and POST, as stated.

# Golang
I had never used Golang in my life! Can you believe it? And I'm not afraid to say that, it's super serious! But, as we all know that anyone who knows programming paradigms well learns fast (and I was an Algorithm subject monitor in college), that's exactly what happened. I confess that I loved the operation of the panic, defer and recover function and added to the use of the middleware gin.Recovery () was very good.
I used the Gin framework in this application.
I learned a lot in those 5 days, since I received this test on Tuesday.

# Private repository on GitHub
Here is another tip for evaluating my competitors, because with the private repository the level of difficulty increases due to the need to exchange information between the containers and the host. There is SSL complication, key problem with .ssh/id_rsa, ssh-agent, ssh-add, etc. It's just not worth removing the repository privately or exposing credentials for everyone to see. (Personal tokens ?! No!)

# Several commits during the development
I think I committed a lot for 5 days, huh?
The commits are there, all with messages and each with its own feature.

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

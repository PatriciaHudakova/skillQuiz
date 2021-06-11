# skillQuiz

- [High Level Overview](#overview)
- [Project Structure](#structure)
- [Running the Program](#execution)
- [Future Improvements](#future improvements)

### High Level Overview <a name="overview"></a>

skillQuiz is a simple questionnaire runner developed with Go and SQL Lite and it works as follows:
- running the program will prompt you with a series of questions
- please answer with a "yes" or "no" as appropriate
- when the questionnaire is over, your current run and average rating of all subsequent runs is calculated and displayed

The Scores are calculated as described below:
- **Current Run:** a counter is initialised and incremented each time a "yes" is recorded for a question, then the average of the current
    run becomes 100 * (counter/n.o. of questions)
- **Overall Average:** firstly, the database is checked if it contains an average entry. If so, the value is retrieved and new average is 
    calculated with the current run average value in mind. This new value becomes the new average and replaces the old value in the database. 
    If not, the current run average becomes the overall average and is inserted into the table
  
SQL Table has the following columns:

- **uuid:** it's a good practice to have a unique identifier for each entry, INTEGER, PRIMARY KEY
- **overallAverage:** holds the cumulative average, INTEGER, NOT NULL

### Project Structure <a name="structure"></a>

Root directory contains the following files:
- **Dockerfile:** docker config file
- **go.mod**: configuration file for go dependencies
- **go.sum**: companion file of _go.mod_; it stores the checksums of each external dependency
- **main.go:** the entry point of the program  
- **main.go:** the entry point of the program
- **README.md:** this documentation file

Pkg level directory contains:

- **db:** package where db related code is located
- **rating:** package where rating calculation related code is located with unit tests
- **questions.go:** contains a helper function that prints questions and record responses

### Running the Program <a name="execution"></a>

To run the program:
```
1) install sql lite if already not installed (skip if on iOS: as they have this preinstalled by default)
2) in pkg > db > db.go, check the driver name and data source match your local setup. For simplicity, I've used the default values
3) start the databse and execute the script in init.sql located in the root directory to create the required table
4) from the root directory, execute go run main.go
```

To execute all unit tests:
```
go test ./...
```

### Future Improvements <a name="future improvements"></a>

Possible ways to build up on existing implementation are:
1) Dynamic creation of questions by the user for a more tailored experience
2) Multi-user support where the overall average is only displayed for skills the user has responded "yes" to. For instance,
the program would display something like "You have a skill only %x of people have".
3) Optionally one could display overall average per skill for all registered skills
4) Expanded relational database with a user table related to the skill table
5) In line with multi-user support, authentication would also be required, possibly using oAuth0
6) Build on resilience by considering corner case scenarios: more than one average in table? how so? how can we prevent it?
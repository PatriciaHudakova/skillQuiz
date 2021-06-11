# questionnaire-runner

- [High Level Overview](#overview)
- [Project Structure](#structure)
- [Running the Program](#execution)
- [Future Improvements](#futureImprovements)

### High Level Overview <a name="overview"></a>

This is a simple questionnaire runner developed with Go and SQL Lite and it works as follows:
- running the program will prompt you with a series of questions
- please answer with a "yes" or "no" as appropriate. Acceptable responses also are "y", "n", "YES", "No"...etc.
- when the questionnaire is over, your current run and average rating of all subsequent runs is calculated and displayed

The Scores are calculated as described below:
- **Current Run:** a counter is initialised and incremented each time a "yes" (or equivalent) is recorded for a question, then the average of the current
    run becomes 100 * (counter/n.o. of questions)
- **Overall Average:** firstly, the database is checked if it contains an average entry. If so, the value is retrieved and new average is 
    calculated with the current run average value in mind. This new value becomes the new average and replaces the old value in the database. 
    If absent, the current run average becomes the overall average and is inserted into the table
  
SQL Table has the following columns:

- **uuid:** it's a good practice to have a unique identifier for each entry, INTEGER, PRIMARY KEY
- **overallAverage:** holds the cumulative average, INTEGER, NOT NULL

### Project Structure <a name="structure"></a>

Root directory contains the following files:
- **go.mod**: configuration file for go dependencies
- **go.sum**: companion file of _go.mod_; it stores the checksums of each external dependency
- **init.sh:** a script file to create database resources as needed 
- **main.go:** the entry point of the program
- **README.md:** this documentation file

Pkg level directory contains:

- **db:** package where db related code is located
- **rating:** package where rating calculation related code is located with unit tests
- **questions.go:** contains a helper function that prints questions and record responses

### Running the Program <a name="execution"></a>

Required GO version:
```
GO 1.16
```

To run the program:
```
1) install sql lite if already not installed (skip if on iOS: as they have this preinstalled by default)
2) run init.sh located in the root directory
3) from the root directory, execute go run main.go
```

To execute all unit tests:
```
go test ./...
```

### Future Improvements <a name="futureImprovements"></a>

Possible ways to build up on existing implementation are:
1) Add a Dockerfile for a more seamless setup
2) Be able to load questions from the DB
3) Be able to add dependent questions e.g. a parent question can have many child questions

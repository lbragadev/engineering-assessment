# Engineering Challenge

We strive to be a practical and pragmatic team. That extends to the way that we work with you to understand if this team is a great fit for you. We want you to come away with a great understanding of the kind of things that we actually do day to day and what it is like to work in our teams.

We don't believe that whiteboard coding with someone watching over your shoulder accurately reflects our day to day. Instead we'd like to be able to discuss code that you have already written when we meet.

This can be a project of your own or a substantial pull request on an open source project, but we recognize that most people have done private or proprietary work and this engineering challenge is for you.

We realize that taking on this assignment represents a time commitment for you, and we do not take that lightly. Throughout the recruitment process we will be respectful of your time and commit to working quickly and efficiently. This will be the only technical assessment you'll be asked to do. The brief following conversations will be based on this assessment and your prior experiences.

## Challenge Guidelines

* This is meant to be an assignment that you spend approximately two to three hours of focused coding. Do not feel that you need to spend extra time to make a good impression. Smaller amounts of high quality code will let us have a much better conversation than large amounts of low quality code.

* Think of this like an open source project. Create a repo on Github, use git for source control, and use a Readme file to document what you built for the newcomer to your project.

* We build systems engineered to run in production. Given this, please organize, design, test, deploy, and document your solution as if you were going to put it into production. We completely understand this might mean you can't do much in the time budget. Prioritize production-readiness over features.

* Think out loud in your documentation. Write our tradeoffs, the thoughts behind your choices, or things you would do or do differently if you were able to spend more time on the project or do it a second time.

* We have a variety of languages and frameworks that we use, but we don't expect you to know them ahead of time. For this assignment you can make whatever choices that let you express the best solution to the problem given your knowledge and favorite tools without any restriction. Please make sure to document how to get started with your solution in terms of setup so that we'd be able to run it.

* Once this is functioning and documented to your liking, either send us a link to your public repo or compress the project directory, give the file a pithy name which includes your own name, and send the file to us.

## The Challenge

As the song says, "you've got to play the hand you're dealt", and in this case your hand is to implement something to help us manage our food truck habit.

Our team loves to eat. They are also a team that loves variety, so they also like to discover new places to eat.

In fact, we have a particular affection for food trucks. One of the great things about Food Trucks in San Francisco is that the city releases a list of them as open data.

Your assignment is to make it possible for our teams to do something interesting with this food trucks data.

This is a freeform assignment. You can write a web API that returns a set of food trucks. You can write a web frontend that visualizes the nearby food trucks for a given place. You can create a CLI that lets us get the names of all the taco trucks in the city. You can create system that spits out a container with a placeholder webpage featuring the name of each food truck to help their marketing efforts. You're not limited by these ideas at all, but hopefully those are enough help spark your own creativity.
San Francisco's food truck open dataset is [located here](https://data.sfgov.org/Economy-and-Community/Mobile-Food-Facility-Permit/rqzj-sfat/data) and there is an endpoint with a [CSV dump of the latest data here](https://data.sfgov.org/api/views/rqzj-sfat/rows.csv). We've also included a copy of the data in this repo as well.


# Setup
### 1. SQL Migrate (Required)
SQL Migrate is used for running migrations inside ./db/migrations folder

To install the library and command line program, use the following:

```bash
go get -v github.com/rubenv/sql-migrate/...
```
### 2. gorilla/mux (Required)

```bash
go get github.com/gorilla/mux
```
### 3. godotenv (Required)

```bash
go get github.com/joho/godotenv
```

# Dev notes

Plan of action:
I will be building a python script that will integrate with the sf food_trucks api. The python script will process data and insert the data into a postgres database. I will then build a golang rest api server that will serve food trucks data. This api will add extra functionality like filtering.

Architecture Decision #1 
I will be using Postgres for the database because it is something i am familiar with, it is popular relational database solution that is tried and tested.

Architecture Decision #2
I will be using Docker for deploying my database locally and online. Docker will provide many benefits easy deployment, isolation, scalability, flexibility and more.

Architecture Decision #3
I wil be using python to create a script of ingesting the data from the sfgov SODA API and inserting it to the database. I'm choosing python because its very easy to quickly write a script that will handle this use case.

Architecture Decision #4 
I will be using golang to create the rest api server. I'm choosing golang because it's something that i'm comfortable with and used recently to build a rest api service. Golang is also a highly performant compiled language. It has nice features like native support of concurrency.

### Nice to haves

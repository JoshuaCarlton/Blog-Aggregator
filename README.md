# Gator
Gator is a CLI designed to aggregate RSS feed posts and browse the most recent additions.

## Installation
For this program to work, you will need to have Go and PostgreSQL installed. Once those are installed, follow the instructions below. I am running this on a wsl/linux environment, so you may need to do things a bit differently on another os.

Create SQL Database
-------
Before we can use the gator CLI, we need a SQL database for it to store things in.
The following instructions should walk you through getting the database up and running on a wsl/linux machine with a fresh install of PostgreSQL.
First, we will set the password for postgres
```sudo passwd postgres```
Then enter your password. I simply used "postgres" for mine, but whatever you choose, just make sure to remember it.
Now we need to start up the Postgres server 
```sudo service postgresql start```
Then connect to the server.
```sudo -u postgres psql```
You should see your prompt change to something like this.
postgres=#
Now we need to create a new database
```CREATE DATABASE gator;```
The prompt should now look like this.
gator=#
Now we need to set the password for this specific database. Again, I used "postgres" for simplicity.
```ALTER USER postgres PASSWORD 'postgres';```
Now that we have our database, you can type exit to leave the psql shell.

Add Database Tables
-------
At this point, we need to run some migrations to get the database up and running with all the necessary tables. I used Goose for my SQL migrations, so you will need to install Goose.
```go install github.com/pressly/goose/v3/cmd/goose@latest```
You will also need my SQL migration files. The easiest way to do that is to clone this repository.
```git clone https://github.com/JoshuaCarlton/gator.git```
Use cd to navigate into the sql/schema folder of this repository. 
Now we need the connection string that we will use to connect to the database we just created. The format for this connection string looks like this.
```protocol://username:password@host:port/database```
If you are running this CLI on a local host database, as this section assumes, you may want to add ?sslmode=disable to ensure an unencrypted connection. If you followed all the defaults from this section so far, your connection string should look like this.
```postgres://postgres:postgres@localhost:5432/gator?sslmode=disable```
Test your connection string to make sure it's working.
```psql "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"```
You should see the gator=# prompt again if it's working.
Now, from the sql/schema folder, we need to have goose migrate the table up so that it can store the correct data
```goose postgres "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable" up```

Configure and Install
-------
Now you will need to create a ~/.gatorconfig.json file in your home directory. This will store our connection string and the currently logged-in user, so put the following content into that JSON file.
```{"db_url":"postgres://postgres:postgres@localhost:5432/gator?sslmode=disable","current_user_name":"init"}```
Now we can finally install gator
```go install github.com/JoshuaCarlton/gator```
Now you should be ready to go, and you can use terminal commands to use gator from anywhere.

## Gator Commands Usage
All commands for the gator CLI will start with gator followed by the name of the command you want to use, then depending on the command you may need to add additional arguments after that.

The first command you will want to run is register. This will create a user with the name you provide and log you in as that user.
```gator register <username>```

You can create multiple users by using register multiple times. If you want to list all the users, you can run users.
```gator users```
To switch users, use login.
```gator login <username>```

Once you are logged in, you can add an RSS feed to pull posts from with addfeed. You will need to provide the name you want the new feed to be called and the link to get to that feed. The currently logged-in user will be recorded as the creator of that feed and will automatically follow that feed.
```gator addfeed "<name of feed>" "<link to feed>"```

You can't make a duplicate feed if someone else already made one for that link, but you can follow the same feed with follow.
```gator follow <link to feed>```

You can also unfollow a feed with unfollow.
```gator unfollow <link to feed>```

You can list all the currently added feeds with feeds.
```gator feeds```

You can list all the feeds followed by the currently logged-in user with following.
```gator following```

Once you have some feeds, you can run agg to aggregate posts from those RSS feeds. This will pull posts from each feed one by one in a loop as often as you tell it to. So you will need to provide the amount of time between each request. This can be something like 10s for 10 seconds or 1m for 1 minute. This command will keep going until you stop it by pressing Ctrl + c. This will store the posts from the RSS feeds in the database.
```gator agg <interval>```

Once you have some posts, you can run browse. This will show the links to some of the most recently pulled posts. You can put a number after browse to control how many posts to show. If you do not provide a number, it will default to 2.
```gator browse <number of posts>```

If you ever want to reset the database, you can do that with reset.
```gator reset```

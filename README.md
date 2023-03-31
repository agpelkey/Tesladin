# Tesladin

## Service where a user sends a file via the curl command, and gets a response with a URL. The user can then re-download the file via the URL.

## If the user loses the URL..... YIKES.

## This project is a challenge. Only the official Go docs found [here](https://go.dev/#) can be used as reference material to build this project. No Youtube, ~~no Google~~ no ChatGPT, just me and the docs baby.
  - So some shit happened. I ended up having to do some googling. 
  - Up until the start of this challenge my experience with databases had pretty much been with relational databases (Postgres). This became a major hurdle. I learned quickly that in this use case, a relational database was probably not the move. I spent about a week trying to figure out how to serialize the data in a way that Postgres would store it as a BYTEA data type. While I do still think this would be possible, I am also a big believer in using the right tool for the job.
  - I made the decision to switch to MondoDB, and this was the correct tool for the job. 
  - Ulitamely I did have to do some googling to determine if a non-relational database was the correct route, but from there only the Mongo docs found [here](https://www.mongodb.com/docs/) were used.

#### To run the service execute "make run".

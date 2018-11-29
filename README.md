# api-wrapper-kata

In this kata we propose a simple user API, and ask the user to augment and enhance it to offer more functionality. 

The success of the kata is achieved if `./check-solution.sh {{ solution folder }}` passes with zero errors. Inside the solution folder should be a script `run.sh` which starts the new API server.

## Old API

The old API is presented in port 5000 and exposes the following methods

    /users returns a list with all users in the system

## New API definition

The new API has to be presented in port 5001 and is required to have the following methods (each with preference over the next):

    /users               returns a list with all users in the system
    /users?name=<string> returns the user with that particular name, error 404 if none found
    /users?type=<string> returns a list with all users with that particular type, empty list if none found
  

# api-wrapper-kata

In this kata we propose a simple user API, and ask the user to augment and enhance it to offer more functionality. 

The success of the kata is achieved if "go test" passes with zero errors.

## Old API

The old API is presented in port 5000 and exposes the following methods

/users returns a list with all users in the system

## New API definition

The new API has to be presented in port 50001 and is required to have the following methods:

/users returns a list with all users in the system
/users?type=<string> returns a list with all users with that particular type
  

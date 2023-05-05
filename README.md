# L33K3D

## Overview:
We used the haveibeenpwned API in order to check the security of login credentials on a demo local webserver. The overall goal was to bring attention to the importance of using secure login credentials on all accounts.

## Implementation:
### CheckPassword()
In order to achieve our goals for this project, we purchased the base API key and managaed various API calls to the haveibeenpwned database. The CheckPassword() function takes in a hash of the users password and passes the first 5 characters to through the API call. The API will then return a list of all password hashes that match the first 5 characters tat were sent, along with the number of times each was found in a data breach as a large string. CHeckPassworrd() then splits the string into a slice that contains each hash and the number of breaches separate from each other. Finally, CheckPassword() loops through this slice, finds the correct hash, and returns the number of breaches. If the hash is not found in the database, CheckPassword() returns 0. The API call for passwords does not have any restrictions on rates, so no API key was needed in this function.

### CheckEmail()
Along with ChekPassword(), we also created another function, CheckEmail(). This function takes the email that a user submits to the login page and sends it to the haveibeenpwned API. CheckEmail() then takes the response from the HTTP request and decoes it into json. It will then marshal the data into a struct that is more organized. Each struct has a section for a breach description that is recieved from the API. CheckEmail() will loop through the struct and return the description of each breach. The API call for email acounts does have rate limiting restrictions. So, we had to incorperate a go rate limiting library for each HTTP request sent for emai accounts.

## Usage:
The use of this small library is fairly simple.
### CheckPassword()
Takes in a single hashed password string as an argument. In order for it to be ran correctly, each password must be passed to the function as an sha1 hash. The function simply returns the amount of times the hash has been found in a data breach as a string. This can be displayed to the user in any capacity.

### CheckEmail()
Takes in a single email string as an argument. Each email must be passed in normal email form. The function returns an array of strings. Each element in the array is a description of some sort. If the email is not found in the API, the array will be of length 0. Each description string contains certain HTML tags as links to other pages for more information. It may fall on the user to figure out how to display this information.

### Example
We have provided a basic example on how to integrate L33K3D into a webapp. When a password or email is found in the API, the user is redirected to a security summary page when they submit their credentials. This page lays out all of the details of their email accounts and passwords being breached in a nice fashion. If the email or password are not found, the user is redirected to the home page.

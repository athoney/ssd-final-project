# L33K3D

## Overview
We used the haveibeenpwned API in order to check the security of login credentials on a demo local webserver. The overall goal was to bring attention to the importance of using secure login credentials on all accounts.

## Implementation
### CheckPassword():
In order to achieve our goals for this project, we purchased the base API key and managaed various API calls to the haveibeenpwned database. The CheckPassword() function takes in a hash of the users password and passes the first 5 characters to through the API call. The API will then return a list of all password hashes that match the first 5 characters tat were sent, along with the number of times each was found in a data breach as a large string. CHeckPassworrd() then splits the string into a slice that contains each hash and the number of breaches separate from each other. Finally, CheckPassword() loops through this slice, finds the correct hash, and returns the number of breaches. If the hash is not found in the database, CheckPassword() returns 0.

### CheckEmail():

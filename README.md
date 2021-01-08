# KARHOO API Get Started Application

## Running

This application assumes the environment is MacOS/Linux, type 
```shell
cd <path-to-project>
./start
```
runs the project

## What it does

   1. Retrieve user credentials from `./cred.sandbox.yml`
   2. Get access token with the credentials
   3. Refresh access token with refresh token if needed
   4. Register a webhook with URL [http://karhoo-webhooks.piizu.com/webhook](http://karhoo-webhooks.piizu.com/webhook)
   5. Get registered webhook URL and print it out in console
   6. Get quotes with designated origin and destination but no pick up time(immediate pick up)
   7. Retrieve quote list with quote ID from step 6, and then print out the quote list
   8. Aggregate quotes from step 7 so that for all quotes with the same vehicle class, only the quote with the lowest price will be saved, and then print out the quotes with lowest price for each vehicle class
   9. Choose a quote from step 8 and make a booking, print out the response
   10. Get booking details with the booking ID from 9, and then print out the booking details
   11. Cancel the booking

## Notes

1. The signature in webhook is actuall "X-Karhoo-Request-Signature"
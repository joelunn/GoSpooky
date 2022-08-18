## GoSpooky
Dockerised Golang Implementation of a tool I previously wrote in Powershell and C# ASP.MVC called Spooky.

Spooky is a Network Inventory & Orchestration tool which in this version is containerised for portability.
The container itself responds to agents with information for executing orchestration on the Agent's connected networks as well as receiving information from Agents to add new devices to an inventory. Both the Inventory and Orchestration features can be polled by an intutive API allow for Go.Spooky to be extended and integrated by the user as needed.

## Usage

To be added once it's working. 

It's likely going to work like every other tool similar in which user/agent must:

1. Authenticate using user credentials
2. Receive a token
3. Do your stuff until the token expires.
4. Goto 1
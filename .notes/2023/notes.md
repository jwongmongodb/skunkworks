# day1

## open questions
1. How much is the compute costing me?
2. is 13 MB of compute usage considered large?
3. What are the different ways to upload datamodels into atlas for my cms app service?
4. In sql we usually define our schema models before creating the endpoints, is the philosophy of mongodb to DISCOVER our datamodels and use the schema tool to suggest a schema?

## suggestions:
1. Perhaps the functions page should link to the documentation.
2. When creating the endpoint, it asks for authentication provider, it would be nice if it linked to the tab.

### Goals
- [x] List out use cases and select one or two to execute on.
    - Create a fully decoupled backend
    - Upload datamodels into atlas
- [ ] Have some working usages of openAI with golang code. (Problem: hit usage limit, punting for now)
- [X] Setup initial project structure.

### notes
- Tried open ai API and have no more credits. Will pick this up later
- Created an app to test out https endpoint usage
- Noticed with good prompts that a secret is needed
- Needed to create a user after hitting the endpoint with the secret
- Needed to enable email and authentication providers. The prompt noted this, would be nice if there was a link in the suggestion.
    - Seems like the compute used for a simple app is a bit large. 1_000_000 bytes * ms when the log indicated that the function took 13ms. Is this taking up 13 MB of compute. I will attempt to check how much this is costing me.
- Created new atlas cluster and three documents to simulate CMS operations.
- Created 2 https endpoints. It was useful to see errors linking to the logs from the response.
- Added rules for each collection
- Able to hit both endpoint calls. Perhaps the functions page should link to the documentation.
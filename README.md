## Faraway test task monorepo

Deviations from the test task/peculiarities:  
 - NFT Contract's `mint` function uses autoincremented token ids instead of passing it as an argument  
 - NFT minting is done via NFT collection contract, not NFT Factory, which makes it harder to catch events on the backend, as we need to monitor all collection addresses. Looks like a better practice imo
 - Backend scanner only listens to factory's `CollectionCreated` event, but setting up subscriptions to `TokenMinted` event is also possible

